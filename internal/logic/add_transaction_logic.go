package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	values "ks_api_service/common/value"
	"ks_api_service/internal/model"
	"ks_api_service/internal/svc"
	"ks_api_service/internal/types"
	"ks_api_service/pb/api"

	json "github.com/bytedance/sonic"
	g "github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddTransactionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTransactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTransactionLogic {
	return &AddTransactionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 增加新的账变
func (l *AddTransactionLogic) AddTransaction(in *api.AddTransactionReq) (*api.AddTransactionResp, error) {

	var (
		resp *api.AddTransactionResp
	)

	if in.Amount == 0 {
		resp.Status = values.TRANS_CREATE_ERROR_DATA
		return resp, errors.New("")
	}

	u, err := model.UserFindOne(g.Ex{"id": in.Account.UserId})
	if err != nil {
		return resp, err
	}

	transactionType, err := model.TransactionTypeFindOne(g.Ex{"id": in.TypeId})
	if err != nil {
		return resp, err
	}

	// 判断是否用户提现 如果是处理冻结金额
	tran := CompileTransactionData(in, transactionType, u)
	if (transactionType.ID == values.TYPE_WITHDRAW || transactionType.ID == values.QUICK_WITHDRAW || transactionType.ID == values.PROCESS_FEE) && transactionType.Amount == 0 {
		if in.Account.Amount >= in.Amount {
			in.Account.Amount = in.Account.Amount - in.Amount
		} else {
			resp.Status = values.TRANS_LOCK_MONEY_ERROR
			return resp, errors.New("")
		}
	}

	id, err := model.TransactionInsert(tran)
	if err != nil {
		return resp, err
	}

	ex := g.Ex{"id": in.Account.Id}
	record := g.Record{}
	record["amount"] = in.Account.Amount
	record["lock_money"] = in.Account.LockMoney
	record["tuijian"] = in.Account.Tuijian
	record["deposit"] = in.Account.Deposit
	record["withdraw"] = in.Account.Withdraw
	record["deposit_times"] = in.Account.DepositTimes
	record["withdraw_times"] = in.Account.WithdrawTimes
	err = model.AccountUpdate(ex, record)
	if err != nil {
		return resp, err
	}

	// 赠送彩金同步赠送彩金表
	if in.ExtraData.IsSyncGiftMoneyTransaction != 0 {
		err = model.GiftMoneyTransactionsInsert(tran)
		if err != nil {
			return resp, err
		}
	}

	//统计打码量
	if in.ExtraData.BetAmount != 0 {
		in.ExtraData.BetAmount = in.ExtraData.BetAmount * transactionType.BetAmount
	} else {
		in.ExtraData.BetAmount = in.Amount * transactionType.BetAmount
	}
	in.ExtraData.Username = tran.Username
	in.ExtraData.TransactionID = id

	//重置的时候 把用户的打码量更新到打码量表
	err = l.DealBetAmount(in.ExtraData, transactionType)
	if err != nil {
		resp.Status = values.TRANS_BET_AMOUNT_ERROR
		return resp, err
	}

	//充值时 单笔限额
	level, err := model.PayLevelFindOne(g.Ex{"id": u.PayLevel, "is_open": 1})
	if err != nil {
		resp.Status = values.TRANS_INCOME_MAX_LIMIT
		return resp, err
	}

	if u.PayLevel != 0 {
		if in.Amount > level.IncomeMaxLimit {
			resp.Status = values.TRANS_INCOME_MAX_LIMIT
			return resp, nil
		}
	}

	//升级会员支付层级
	if transactionType.PayLevel != 0 {
		//获得用户的支付层级 判断当前用户支付层级是否被锁定
		//如果没锁住就判断升级
		if level.IsLock != 1 {
			plex := g.Ex{}
			plex["recharge_times"] = g.Op{"lte": in.Account.DepositTimes}
			plex["withdraw_times"] = g.Op{"lte": in.Account.WithdrawTimes}
			plex["single_recharge_amount"] = g.Op{"lte": tran.Amount}
			plex["id"] = g.Op{"gt": u.PayLevel}
			plex["is_lock"] = 0
			userPayLevel, err := model.PayLevelFindOne(plex)
			if err != nil {
				resp.Status = values.TRANS_INCOME_MAX_LIMIT
				return resp, err
			}

			if userPayLevel.ID > u.PayLevel {
				uex := g.Ex{"id": u.ID}
				urecord := g.Record{}
				urecord["pay_level"] = userPayLevel.ID
				err = model.UserUpdate(uex, urecord)
				if err != nil {
					return resp, err
				}

				ulog := types.LevelUpgradeLog{
					UserID:        u.ID,
					Username:      u.Username,
					OldPayLevelID: u.PayLevel,
					NewPayLevelID: userPayLevel.ID,
					AdminUser:     "system_auto",
					Remark:        "管理后台自动触发",
					CreatedAt:     time.Now(),
				}
				_, err = model.LevelUpgradeLogInsert(&ulog)
				if err != nil {
					return resp, err
				}
			}
		}
	}

	// 如果是存款、手动存款、极速存款类型更新会员首次和末次存款
	if transactionType.ParentID == 2 {

		tr, err := model.TransactionLastOne(g.Ex{"id": id})
		if err != nil {
			return resp, err
		}

		err = l.UpdateFirstAndLastDeposit(u, tr)
		if err != nil {
			return resp, err
		}

	}

	// 世界杯存款
	activity, err := model.ActivityFindOne(g.Ex{})
	if err != nil {
		return resp, err
	}

	now := time.Now()
	if activity.StartTime.Before(now) && activity.EndTime.After(now) && activity.Status == 1 && transactionType.ParentID == 2 {
		tu, err := model.TeamUserFindOne(g.Ex{"user_id": tran.UserID})
		if err != nil {
			return resp, err
		}

		var act types.ActTransaction
		err = copier.Copy(&act, tran)
		if err != nil {
			return resp, err
		}

		act.LeaderID = tu.LeaderID
		_, err = model.ActTransactionInsert(&act)
		if err != nil {
			return resp, err
		}

		report, err := model.ActivityReportFindOne(g.Ex{"user_id": tu.LeaderID})
		if err != nil {
			return resp, err
		}

		if report.ID != 0 {
			reportex := g.Ex{"id": report.ID}
			reportrecord := g.Record{}
			reportrecord["total_deposit"] = g.L("total_deposit + ?", tran.Amount)
			reportrecord["updated_at"] = time.Now()
			err = model.ActivityReportUpdate(reportex, reportrecord)
			if err != nil {
				return resp, err
			}
		} else {
			r := types.ActivityReport{
				TotalDeposit: tran.Amount,
				UserID:       tu.LeaderID,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			_, err = model.ActivityReportInsert(&r)
			if err != nil {
				return resp, err
			}
		}

		teamex := g.Ex{"id": report.ID}
		teamrecord := g.Record{}
		teamrecord["total_deposit"] = g.L("total_deposit + ?", tran.Amount)
		teamrecord["updated_at"] = time.Now()
		err = model.TeamUserUpdate(teamex, teamrecord)
		if err != nil {
			return resp, err
		}
	}

	model.EsTransaction(l.ctx, id)
	model.EsUser(l.ctx, u.ID)

	return resp, nil
}

func CompileTransactionData(
	in *api.AddTransactionReq,
	transType types.TransactionType,
	user types.User,
) *types.Transaction {

	attr := &types.Transaction{
		UserID:            in.Account.UserId,
		Username:          user.Username,
		ParentID:          user.ParentID,
		ForefatherIds:     user.ForefatherIds,
		TransFatherID:     transType.ParentID,
		TransTypesID:      transType.ID,
		TransTypesCnTitle: transType.CnTitle,
		TransTypesEnTitle: transType.EnTitle,
		Amount:            in.Amount,
		IsIncome:          transType.Amount,
		BeforeMoney:       in.Account.Amount,
		Description:       in.ExtraData.Description,
		PayType:           in.ExtraData.PayType,
		BankType:          in.ExtraData.BankType,
		IP:                in.ExtraData.IP,
		ThirdMerchantName: in.ExtraData.ThirdMerchantName,
		MerchantNum:       in.ExtraData.MerchantNum,
		ThirdTrackNum:     in.ExtraData.ThirdTrackNum,
		Issue:             in.ExtraData.Issue,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		IsTester:          in.ExtraData.IsTester,
	}

	// Optional fields
	if in.ExtraData.AdminID != 0 {
		attr.AdminID = in.ExtraData.AdminID
	}
	if in.ExtraData.AdminName != "" {
		attr.Adminname = in.ExtraData.AdminName
	}
	if in.ExtraData.BillID != "" {
		attr.BillID = in.ExtraData.BillID
	}
	if in.ExtraData.GameCode != "" {
		attr.GameCode = in.ExtraData.GameCode
	}
	if in.ExtraData.Remark != "" {
		attr.Remark = in.ExtraData.Remark
	}
	if in.ExtraData.ForefatherIDs != "" {
		attr.ForefatherIds = in.ExtraData.ForefatherIDs
	}

	// 更新账户金额和次数
	subFields := map[string]int64{
		"amount":   transType.Amount,
		"deposit":  transType.Deposit,
		"withdraw": transType.Withdraw,
		"tuijian":  transType.Tuijian,
	}

	for field, value := range subFields {
		if value == 0 {
			continue
		}
		switch field {
		case "amount":
			in.Account.Amount += float64(value) * in.Amount
		case "deposit":
			in.Account.Deposit += float64(value) * in.Amount
			in.Account.DepositTimes += transType.DepositTimes
		case "withdraw":
			in.Account.Withdraw += float64(value) * in.Amount
			in.Account.WithdrawTimes += transType.WithdrawTimes
		case "tuijian":
			in.Account.Tuijian += float64(value) * in.Amount
		}
	}

	if transType.ID == values.TYPE_WITHDRAW || transType.ID == values.QUICK_WITHDRAW || transType.ID == values.PROCESS_FEE {
		attr.IsIncome = -1
		attr.BeforeMoney = in.Account.Amount + in.Account.LockMoney
		attr.Money = in.Account.Amount + in.Account.LockMoney - in.Amount
	} else {
		attr.Money = in.Account.Amount
	}

	return attr
}

func (l *AddTransactionLogic) DealBetAmount(aExtraData *api.ExtraData, oTransactionType types.TransactionType) error {

	ex := g.Ex{}
	ex["user_id"] = aExtraData.UserId
	ex["is_open"] = 1
	ba, err := model.BetAmountFindOne(ex)
	if err != nil {
		return err
	}

	if oTransactionType.ID == values.TYPE_WITHDRAW || oTransactionType.ID == values.WITHDRAW_BY_ADMIN || oTransactionType.ID == values.QUICK_WITHDRAW {

		// 把用户的打码量重置为0
		bex := g.Ex{}
		bex["id"] = ba.ID
		brecord := g.Record{}
		brecord["is_open"] = 0
		brecord["delete_at"] = time.Now()
		err = model.BetAmountUpdate(bex, brecord)
		if err != nil {
			return err
		}

		//重置的时候 把用户的打码量更新到打码量表
		ugbaex := g.Ex{}
		ugbaex["user_id"] = aExtraData.UserId
		ugbaex["is_open"] = 1
		ugbas, err := model.GetUserGameBetAmount(ugbaex)
		if err != nil {
			return err
		}

		var ids []int
		for _, v := range ugbas {
			ids = append(ids, int(v.ID))
		}
		idstr := joinSortedIDs(ids)
		cbalex := g.Ex{}
		cbalex["user_id"] = aExtraData.UserId
		cbalex["status"] = 0
		log, err := model.ClearBetAmountLogFindOne(cbalex)
		if err != nil {
			return err
		}

		if idstr == log.UserGameBetAmountIds {
			return nil
		}

		if idstr != "" {
			alog := types.ClearBetAmountLog{
				UserID:               aExtraData.UserId,
				Username:             aExtraData.Username,
				UserGameBetAmountIds: idstr,
				CreatedAt:            time.Now(),
				UpdatedAt:            time.Now(),
			}
			id, err := model.ClearBetAmountLogInsert(&alog)
			if err != nil {
				return err
			}

			logMap := make(map[string]interface{})
			logMap["log_id"] = id
			logMap["user_id"] = aExtraData.UserId
			// 创建Channel
			ch, err := l.svcCtx.Rabbitmq.Channel()
			if err != nil {
				logx.Errorf("rabbitmq Channel error: %v", err)
				return err
			}
			defer ch.Close()

			logMapStr, err := json.Marshal(logMap)
			if err != nil {
				logx.Errorf("Marshal Error: %v", err)
				return err
			}

			// 发布消息
			err = ch.Publish(
				"",                 // 使用默认交换机
				"KsClearBetAmount", // 路由键（这里直接用队列名）
				false,              // 强制标志（如果队列不存在则报错）
				false,              // 立即标志（如果无消费者则报错）
				amqp.Publishing{
					DeliveryMode: amqp.Persistent, // 消息持久化（服务重启后仍存在）
					ContentType:  "text/plain",
					Body:         logMapStr,
				})
			if err != nil {
				logx.Errorf("rabbitmq push error: %v", err)
				return err
			}

			//从新插入新的数据
			//$newData['pre_total_amount'] = $oBetAmount->pre_total_amount + $oBetAmount->user_amount; // 用户之前的打码量+实际打码量
			ba := types.BetAmount{
				UserID:    aExtraData.UserId,
				Username:  aExtraData.Username,
				IsOpen:    1,
				CreatedAt: time.Now(),
			}
			_, err = model.BetAmountInsert(&ba)
			if err != nil {
				return err
			}

			baex := g.Ex{"user_id": aExtraData.UserId, "state": 1}
			barecord := g.Record{}
			barecord["state"] = 0
			barecord["delete_at"] = time.Now()
			err = model.BetAmountLogUpdate(baex, barecord)
			if err != nil {
				return err
			}
		}
	} else {
		rate := 1.0
		if aExtraData.BetAmountRate != 0 {
			rate = aExtraData.BetAmountRate
		}
		amount := aExtraData.BetAmount * rate
		if ba.ID != 0 {
			baex := g.Ex{"user_id": aExtraData.UserId, "state": 1}
			barecord := g.Record{}
			newTotal := ba.TotalAmount + amount
			comment := fmt.Sprintf("%d%d%f", time.Now().Unix(), rand.Intn(1000), newTotal)
			barecord["total_amount"] = newTotal
			barecord["comment"] = comment
			barecord["updated_at"] = time.Now()
			err = model.BetAmountUpdate(baex, barecord)
			if err != nil {
				return err
			}
		} else {
			ba := types.BetAmount{
				TotalAmount: amount,
				UserID:      aExtraData.UserId,
				Username:    aExtraData.Username,
				IsOpen:      1,
				CreatedAt:   time.Now(),
			}
			_, err = model.BetAmountInsert(&ba)
			if err != nil {
				return err
			}
		}

		bal := types.BetAmountLog{
			TotalAmount:   amount,
			UserID:        aExtraData.UserId,
			Username:      aExtraData.Username,
			State:         1,
			TransactionID: aExtraData.TransactionID,
			Remark:        aExtraData.Remark,
			CreatedAt:     time.Now(),
		}
		_, err = model.BetAmountLogInsert(&bal)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *AddTransactionLogic) UpdateFirstAndLastDeposit(user types.User, tran types.Transaction) error {

	fs, err := model.RedisGetFissionSetting(l.ctx)
	if err != nil {
		return err
	}

	fsMap := make(map[string]*types.FissionSetting)
	for _, f := range fs {
		fsMap[f.FissionKey] = f
	}

	if user.FromUserCode != "" {
		u, err := model.GetUserById(l.ctx, user.ID)
		if err != nil {
			return err
		}

		level, err := model.PayLevelFindOne(g.Ex{"id": u.PayLevel, "is_open": 1})
		if err != nil {
			return err
		}

		var (
			exclusiveRewardsType int
			//firstRechargeReward  int
			fissionInviterReward []types.FissionInviterReward
		)
		f, ok := fsMap[values.FISSION_EXCLUSIVE_REWARDS]
		if ok {
			num, err := strconv.Atoi(f.FissionVal)
			if err != nil {
				return err
			}

			exclusiveRewardsType = num
		}

		//f, ok = fsMap[values.FISSION_FIRST_RECHARGE_REWARD]
		//if ok {
		//	num, err := strconv.Atoi(f.FissionVal)
		//	if err != nil {
		//		return err
		//	}
		//
		//	firstRechargeReward = num
		//}

		f, ok = fsMap[values.FISSION_INVITE_REWARDS]
		if ok {
			if f.FissionVal != "" {
				err = json.Unmarshal([]byte(f.FissionVal), &fissionInviterReward)
				if err != nil {
					return err
				}
			}
		}

		if u.FirstRechargeAt.IsZero() || tran.CreatedAt.Before(u.FirstRechargeAt) {
			user.FirstRechargeAt = tran.CreatedAt
			user.FirstRechargeAmount = tran.Amount

			if tran.Amount > 100 {
				rewardMoney := 0.0
				c, err := model.FissionRewardCount(g.Ex{"user_id": u.ID, "reward_type": 1})
				if err != nil {
					return err
				}

				c += 1

				for _, reward := range fissionInviterReward {
					if reward.EndNumber != "" && reward.StartNumber != "" && reward.WinMoney != "" {
						startNumber, err := strconv.Atoi(reward.StartNumber)
						if err != nil {
							return err
						}

						endNumber, err := strconv.Atoi(reward.EndNumber)
						if err != nil {
							return err
						}

						winMoney, err := strconv.ParseFloat(reward.WinMoney, 64)
						if err != nil {
							return err
						}

						if c >= startNumber && c <= endNumber && !reward.Ge {
							rewardMoney = winMoney
							break
						} else if c >= startNumber && reward.Ge {
							rewardMoney = winMoney
							break
						}
					}
				}

				fr := types.FissionReward{
					UserID:                 u.ID,
					Username:               u.Username,
					PayLevelsID:            u.PayLevel,
					PayLevelsName:          level.CnName,
					FromUserID:             user.ID,
					FromUsername:           user.Username,
					RewardType:             values.REWARD_TYPE_FIRST_RECHARGE,
					Money:                  rewardMoney,
					RewardStatus:           0,
					CreatedAt:              time.Now(),
					UpdatedAt:              time.Now(),
					UsernameParentName:     u.ParentName,
					FromUsernameParentName: user.ParentName,
				}
				_, err = model.FissionRewardInsert(&fr)
				if err != nil {
					return err
				}

				// 好友首存专属奖励
				if exclusiveRewardsType > 0 {
					fer, err := model.GetFissionExclusiveRewardFirst(exclusiveRewardsType, u.UserLevel, u.PayLevel)
					if err != nil {
						return err
					}

					if fer.RewardType == 3 {
						rv, err := strconv.ParseFloat(fer.RewardValue, 64)
						if err != nil {
							return err
						}

						fr := types.FissionReward{
							UserID:                 u.ID,
							Username:               u.Username,
							PayLevelsID:            u.PayLevel,
							PayLevelsName:          level.CnName,
							FromUserID:             user.ID,
							FromUsername:           user.Username,
							Money:                  rv,
							RewardStatus:           0,
							CreatedAt:              time.Now(),
							UpdatedAt:              time.Now(),
							UsernameParentName:     u.ParentName,
							FromUsernameParentName: user.ParentName,
							RewardType:             values.REWARD_TYPE_FRIEND_FIRST_DEPOSIT_EXCLUSIVE_REWARD, // 0好友绑定账户,1好友首充,2好友打码返利,3好友负盈利比例,4好友充值比例,5线下陪玩,6好友首存专属奖励
						}
						_, err = model.FissionRewardInsert(&fr)
						if err != nil {
							return err
						}
					}
				}
			}
		}

		// 好友充值专属奖励
		if exclusiveRewardsType > 0 {
			fer, err := model.GetFissionExclusiveRewardFirst(exclusiveRewardsType, u.UserLevel, u.PayLevel)
			if err != nil {
				return err
			}

			if fer.RewardType == 1 {
				rv, err := strconv.ParseFloat(fer.RewardValue, 64)
				if err != nil {
					return err
				}

				rewardMoney := tran.Amount * rv / 100
				fr := types.FissionReward{
					UserID:                 u.ID,
					Username:               u.Username,
					PayLevelsID:            u.PayLevel,
					PayLevelsName:          level.CnName,
					FromUserID:             user.ID,
					FromUsername:           user.Username,
					Money:                  rewardMoney,
					RewardStatus:           0,
					CreatedAt:              time.Now(),
					UpdatedAt:              time.Now(),
					UsernameParentName:     u.ParentName,
					FromUsernameParentName: user.ParentName,
					RewardType:             values.REWARD_TYPE_FRIEND_RECHARGE_RATIO, // 0好友绑定账户,1好友首充,2好友打码返利,3好友负盈利比例,4好友充值比例,5线下陪玩,6好友首存专属奖励
				}
				_, err = model.FissionRewardInsert(&fr)
				if err != nil {
					return err
				}
			}
		}

		if user.LastRechargeAt.IsZero() || tran.CreatedAt.After(user.LastRechargeAt) {
			user.LastRechargeAt = tran.CreatedAt
			user.LastRechargeAmount = tran.Amount
		}
		if user.MaxRechargeAmount == 0 || tran.Amount > user.MaxRechargeAmount {
			user.MaxRechargeAmount = tran.Amount
		}

		uex := g.Ex{"id": user.ID}
		urecord := g.Record{}
		urecord["max_recharge_amount"] = user.MaxRechargeAmount
		urecord["last_recharge_at"] = user.LastRechargeAt
		urecord["last_recharge_amount"] = user.LastRechargeAmount
		urecord["first_recharge_at"] = user.FirstRechargeAt
		urecord["first_recharge_amount"] = user.FirstRechargeAmount
		err = model.UserUpdate(uex, urecord)
		if err != nil {
			return err
		}
	}

	return nil
}

// 将 ID 列表转换为排序后的逗号分隔字符串
func joinSortedIDs(ids []int) string {
	sort.Ints(ids)
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = fmt.Sprintf("%d", id)
	}
	return strings.Join(strIDs, ",")
}
