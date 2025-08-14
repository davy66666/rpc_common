package values

const (
	TRANS_CREATE_SUCCESSFUL     = 1
	TRANS_CREATE_ERROR_DATA     = -101
	TRANS_CREATE_ERROR_SAVE     = -102
	TRANS_CREATE_ERROR_BALANCE  = -103
	TRANS_CREATE_LOW_BALANCE    = -104
	TRANS_CREATE_MONGODB_FAILED = -105
	TRANS_ERROR_TODAY_PROFIT    = -106
	TRANS_INCOME_MAX_LIMIT      = -107
	TRANS_LOCK_MONEY_ERROR      = -108
	TRANS_BET_AMOUNT_ERROR      = -109
)

const (
	TYPE_DEPOSIT              = 1  // 线上支付
	TYPE_WITHDRAW             = 2  // 提取现金
	WITHDRAWAL_FAILED         = 3  // 拒绝提现
	DEPOSIT_BY_ADMIN          = 11 // 手动存款
	WITHDRAW_BY_ADMIN         = 12 // 手动扣款
	RECOMMENT_REWARD          = 13 // 推荐奖励
	CANCEL_REWARD             = 14 // 撤销推荐奖励
	USER_FANSHUI              = 16 // 返水
	CANCEL_USER_FANSHUI       = 17 // 撤销返水
	THIRD_TRANSFER_IN         = 19 // 三方转入
	THIRD_TRANSFER_OUT        = 20 // 三方转出
	PROCESS_FEE               = 26 // 手续费
	CANCEL_PROCESS_FEE        = 27 // 取消手续费
	TYPE_DISCOUNT             = 28 // 优惠
	CANCEL_DISCOUNT           = 29 // 取消优惠
	GIFT_MONEY_IN             = 34 // 赠送彩金
	GIFT_MONEY_BACK           = 35 // 收回彩金
	OFFLINE_RECHARGE          = 36 // 公司入款 会员手动充值
	CANCEL_DEPOSIT            = 37 // 支出 撤销用户充值
	TRANS_PROMOTION_MONEY     = 38 // 晋级礼金promotion_money
	TRANS_BIRTHDAY_MONEY      = 39 // 生日礼金birthday_money
	TRANS_WEEK_MONEY          = 40 // VIP周俸禄week_money
	TRANS_MONTH_MONEY         = 41 // VIP月俸禄month_money
	TRANS_BORROW_MONEY        = 42 // 可借款金额borrow_money
	TRANS_RETURN_MONEY        = 43 // 可借款金额return_money
	TRANS_HONGBAO_MONEY       = 44 // 领取红包
	TRANS_TURNTABLE_MONEY     = 45 // 大转盘红包
	ACTIVITY_APPLY_BONUS      = 46 // 优惠活动申请奖金
	QUICK_RECHARGE            = 47 // 快速充值
	QUICK_WITHDRAW            = 48 // 快速提现
	QUICK_DISCOUNT            = 49 // 快速优惠
	WITHDRAW_DISCOUNT         = 50 // 提现优惠
	NEW_USER_REGISTER_GIFT    = 51 // 新用户注册礼金
	WORLD_CUP_REWARD          = 52 // 世界杯竞猜奖励
	WORLD_CUP_REGISTER_REWARD = 53 // 世界杯推荐奖励
	TRANS_TOURNAMENT          = 60 // 争霸赛奖金
	EUROPEAN_CUP_BET          = 61 // 欧洲杯投注
	EUROPEAN_CUP_WIN          = 62 // 欧洲杯中奖
	BILLION_SUBSIDIES         = 63 // 百亿津贴
)

const (
	FISSION_BIND_PHONE_REWARD        = "fission_bind_phone_reward"        // 被邀请人绑定手机号后，邀请人获取XX彩金
	FISSION_FIRST_RECHARGE_REWARD    = "fission_first_recharge_reward"    // 被邀请人首存≥100元，邀请人获取XX彩金
	FISSION_REPEAT_BIND_PHONE_NUMBER = "fission_repeat_bind_phone_number" // 同一手机号被多个账户绑定后，最多可获取彩金次数
	FISSION_USER_INVITE_URL_H5       = "fission_user_invite_url_h5"       // 移动端邀请链接
	FISSION_USER_INVITE_URL_PC       = "fission_user_invite_url_pc"       // 电脑端邀请链接
	FISSION_INDEX_POP_IMAGE_H5       = "fission_index_pop_image_h5"       // 移动端首页裂变弹窗图
	FISSION_INDEX_POP_IMAGE_PC       = "fission_index_pop_image_pc"       // 电脑端首页裂变弹窗图
	FISSION_INDEX_LEVITATE_IMAGE_H5  = "fission_index_levitate_image_h5"  // 移动端首页裂变悬浮图
	FISSION_INDEX_LEVITATE_IMAGE_PC  = "fission_index_levitate_image_pc"  // 电脑端首页裂变悬浮图

	FISSION_RECOMMEND_POP_IMAGE_H5 = "fission_recommend_pop_image_h5" // 移动端邀请赚钱页面弹窗图
	FISSION_RECOMMEND_POP_IMAGE_PC = "fission_recommend_pop_image_pc" // 电脑端邀请赚钱页面弹窗图
	FISSION_RULES_EXPLAIN          = "fission_rules_explain"          // 裂变规则说明

	FISSION_POSTER_IMAGE_PC = "fission_poster_image_pc" //裂变海報素材图pc
	FISSION_POSTER_IMAGE_H5 = "fission_poster_image_h5" //裂变 海報素材图h5

	FISSION_INVITE_BUTTON_IMAGE_PC = "fission_invite_button_image_pc" //裂变邀請按鈕 pc
	FISSION_INVITE_BUTTON_IMAGE_H5 = "fission_invite_button_image_h5" //裂变邀請按鈕 h5

	FISSION_EXCLUSIVE_REWARDS = "fission_exclusive_rewards" //专属奖励     //0：无专属奖励，1：按照vip等级设置奖励，2：按照支付分层设置奖励
	FISSION_INVITE_REWARDS    = "fission_inviter_reward"    //邀请奖励阶梯

)

const (
	//0好友绑定账户、1好友首充、2好友打码返利
	REWARD_TYPE_BIND_USERNAME                         = 0 //
	REWARD_TYPE_FIRST_RECHARGE                        = 1 //
	REWARD_TYPE_CODE_REBATE                           = 2 //
	REWARD_TYPE_FRIEND_NEGATIVE_PROFIT_RATIO          = 3
	REWARD_TYPE_FRIEND_RECHARGE_RATIO                 = 4
	REWARD_TYPE_OFFLINE_PLAY                          = 5
	REWARD_TYPE_FRIEND_FIRST_DEPOSIT_EXCLUSIVE_REWARD = 6
)
