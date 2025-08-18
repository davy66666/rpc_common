package types

import (
	"database/sql"
	"time"
)

type HelloReq struct {
	Msg string `form:"msg"`
}

type HelloResp struct {
	Msg string `json:"msg"`
}

type PayLevel struct {
	ID        int64  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	EnName    string `gorm:"column:en_name;type:varchar(100);comment:层级名称英文名" json:"en_name" db:"en_name"`                                                       // 层级名称英文名
	CnName    string `gorm:"column:cn_name;type:varchar(100);not null;default:'';comment:层级名称中文名" json:"cn_name" db:"cn_name"`                                   // 层级名称中文名
	PcBindPay string `gorm:"column:pc_bind_pay;type:text;comment:PC绑定的支付以json格式存储{支付类en_name:[具体支付id,以逗号分隔]" json:"pc_bind_pay" db:"pc_bind_pay"` // PC绑定的支付以json格式存储{支付类en_name:[具体支付id,以逗号分隔]
	BindPay   string `gorm:"column:bind_pay;type:text;comment:绑定的支付以json格式存储{支付类en_name:[具体支付id,以逗号分隔]" json:"bind_pay" db:"bind_pay"`            // 绑定的支付以json格式存储{支付类en_name:[具体支付id,以逗号分隔]
	//MaxFanshui               float64      `gorm:"column:max_fanshui;type:decimal(8,2);comment:实时返水最高额" json:"max_fanshui" db:"max_fanshui"`                                           // 实时返水最高额
	Sort      int64        `gorm:"column:sort;type:int;comment:排序" json:"sort" db:"sort"`                                                           // 排序
	Remark    string       `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark" db:"remark"`                                            // 备注
	Editor    string       `gorm:"column:editor;type:varchar(255);comment:编辑者" json:"editor" db:"editor"`                                          // 编辑者
	IsDefault int64        `gorm:"column:is_default;type:tinyint;not null;default:0;comment:是否默认层级0:否;1:是" json:"is_default" db:"is_default"` // 是否默认层级0:否;1:是
	IsOpen    int64        `gorm:"column:is_open;type:tinyint;not null;default:1;comment:是否开启0:否;1:是" json:"is_open" db:"is_open"`              // 是否开启0:否;1:是
	CreatedAt sql.NullTime `gorm:"column:created_at;type:bigint;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	//DingshiMaxFanshui        float64      `gorm:"column:dingshi_max_fanshui;type:decimal(8,2);comment:定时反水最高额" json:"dingshi_max_fanshui" db:"dingshi_max_fanshui"`                                                                                   // 定时反水最高额
	CanRedPacket             int64   `gorm:"column:can_red_packet;type:tinyint;comment:是否容许抢红包0:不可以;1:可以" json:"can_red_packet" db:"can_red_packet"`                                                                                        // 是否容许抢红包0:不可以;1:可以
	CanFanshui               int64   `gorm:"column:can_fanshui;type:tinyint;comment:是否容许返水" json:"can_fanshui" db:"can_fanshui"`                                                                                                                  // 是否容许返水
	CanRecommendFanshui      int64   `gorm:"column:can_recommend_fanshui;type:tinyint;comment:是否容许推荐返水：0:不可以;1:可以" json:"can_recommend_fanshui" db:"can_recommend_fanshui"`                                                                // 是否容许推荐返水：0:不可以;1:可以
	CanPromotionMoney        int64   `gorm:"column:can_promotion_money;type:tinyint;comment:是否可领VIP晋级彩金 0:不可以;1:可以" json:"can_promotion_money" db:"can_promotion_money"`                                                                   // 是否可领VIP晋级彩金 0:不可以;1:可以
	CanWeekMoney             int64   `gorm:"column:can_week_money;type:tinyint;comment:是否可领VIP周俸禄 0:不可以;1:可以" json:"can_week_money" db:"can_week_money"`                                                                                    // 是否可领VIP周俸禄 0:不可以;1:可以
	CanMonthMoney            int64   `gorm:"column:can_month_money;type:tinyint;comment:是否可领VIP月俸禄 0:不可以;1:可以" json:"can_month_money" db:"can_month_money"`                                                                                 // 是否可领VIP月俸禄 0:不可以;1:可以
	CanBirthdayMoney         int64   `gorm:"column:can_birthday_money;type:tinyint;comment:是否可领生日礼金 0:不可以;1:可以" json:"can_birthday_money" db:"can_birthday_money"`                                                                         // 是否可领生日礼金 0:不可以;1:可以
	PcBindPayType            string  `gorm:"column:pc_bind_pay_type;type:varchar(2048);comment:pc绑定通道，以逗号分隔通道en_name" json:"pc_bind_pay_type" db:"pc_bind_pay_type"`                                                                         // pc绑定通道，以逗号分隔通道en_name
	MobileBindPayType        string  `gorm:"column:mobile_bind_pay_type;type:varchar(2048);comment:mobile绑定通道，以逗号分隔通道en_name" json:"mobile_bind_pay_type" db:"mobile_bind_pay_type"`                                                         // mobile绑定通道，以逗号分隔通道en_name
	IncomeMaxLimit           int64   `gorm:"column:income_max_limit;type:int unsigned;not null;default:500000;comment:单笔入款限制" json:"income_max_limit" db:"income_max_limit"`                                                                      // 单笔入款限制
	RechargeTimes            int64   `gorm:"column:recharge_times;type:int;not null;default:0;comment:充值次数" json:"recharge_times" db:"recharge_times"`                                                                                              // 充值次数
	SingleRechargeAmount     int64   `gorm:"column:single_recharge_amount;type:int;not null;default:0;comment:单次充值金额" json:"single_recharge_amount" db:"single_recharge_amount"`                                                                  // 单次充值金额
	IsLock                   int64   `gorm:"column:is_lock;type:tinyint;not null;default:1;comment:是否锁定自动升级" json:"is_lock" db:"is_lock"`                                                                                                       // 是否锁定自动升级
	WithdrawTimes            int64   `gorm:"column:withdraw_times;type:int unsigned;not null;default:0;comment:取款次数" json:"withdraw_times" db:"withdraw_times"`                                                                                     // 取款次数
	QuickRechargeUploadPic   int64   `gorm:"column:quick_recharge_upload_pic;type:tinyint(1);not null;default:1;comment:0:极速充值不需要上传图片凭证,1:极速充值需要上传图片凭证" json:"quick_recharge_upload_pic" db:"quick_recharge_upload_pic"`       // 0:极速充值不需要上传图片凭证,1:极速充值需要上传图片凭证
	QuickRechargeUploadVideo int64   `gorm:"column:quick_recharge_upload_video;type:tinyint(1);not null;default:0;comment:0:极速充值不需要上传视频凭证,1:极速充值需要上传视频凭证" json:"quick_recharge_upload_video" db:"quick_recharge_upload_video"` // 0:极速充值不需要上传视频凭证,1:极速充值需要上传视频凭证
	BankStatementUploadPic   int64   `gorm:"column:bank_statement_upload_pic;type:tinyint;not null;default:0;comment:是否需要上传资金流水图 0:不需要 1:需要" json:"bank_statement_upload_pic" db:"bank_statement_upload_pic"`                           // 是否需要上传资金流水图 0:不需要 1:需要
	UserCount                int64   `gorm:"column:user_count;type:int;not null;default:0;comment:会员数" json:"user_count" db:"user_count"`                                                                                                            // 会员数
	YhzcDiscountRate         float64 `gorm:"column:yhzc_discount_rate;type:decimal(5,2);not null;default:0.00;comment:易汇直充优惠比例" json:"yhzc_discount_rate" db:"yhzc_discount_rate"`                                                              // 易汇直充优惠比例
}
type User struct {
	ID                  int64          `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	Username            string         `gorm:"column:username;type:varchar(255);not null;uniqueIndex:users_name_unique,priority:1;default:''" json:"username" db:"username"`
	APIToken            string         `gorm:"column:api_token;type:varchar(500);not null;default:0" json:"api_token" db:"api_token"`
	ParentID            int64          `gorm:"column:parent_id;type:int unsigned;default:1" json:"parent_id" db:"parent_id"`
	ForefatherIds       string         `gorm:"column:forefather_ids;type:varchar(100);index:idx_forefather_ids,priority:1;default:1;comment:父亲树" json:"forefather_ids" db:"forefather_ids"` // 父亲树
	ParentName          string         `gorm:"column:parent_name;type:varchar(16)" json:"parent_name" db:"parent_name"`
	SignScore           float64        `gorm:"column:sign_score;type:decimal(12,2);not null;default:0.00;comment:签到积分" json:"sign_score" db:"sign_score"` // 签到积分
	IsOpen              int64          `gorm:"column:is_open;type:tinyint;not null;index:idx_pay_level_is_open,priority:2;index:idx_user_level_is_open,priority:2;index:is_open,priority:1;default:1" json:"is_open" db:"is_open"`
	IsAgent             int64          `gorm:"column:is_agent;type:tinyint unsigned;comment:0: 普通用户, 1: 代理;2:总代理;3:股东" json:"is_agent" db:"is_agent"`                                       // 0: 普通用户, 1: 代理;2:总代理;3:股东
	IsTester            int64          `gorm:"column:is_tester;type:tinyint unsigned;comment:0: 普通用户, 1: 测试账号" json:"is_tester" db:"is_tester"`                                                // 0: 普通用户, 1: 测试账号
	UserLevel           string         `gorm:"column:user_level;type:varchar(30);index:idx_user_level_is_open,priority:1;comment:用户层级" json:"user_level" db:"user_level"`                          // 用户层级
	PayLevel            int64          `gorm:"column:pay_level;type:varchar(100);index:idx_pay_level_is_open,priority:1;default:1;comment:支付层级 对应pay_level表" json:"pay_level" db:"pay_level"`   // 支付层级 对应pay_level表
	BankNum             int64          `gorm:"column:bank_num;type:tinyint unsigned;default:5;comment:用户能绑定的卡数" json:"bank_num" db:"bank_num"`                                                 // 用户能绑定的卡数
	IsGuest             int64          `gorm:"column:is_guest;type:tinyint unsigned;comment:是否是游客" json:"is_guest" db:"is_guest"`                                                                 // 是否是游客
	RegCode             string         `gorm:"column:reg_code;type:varchar(60);comment:用户推广码" json:"reg_code" db:"reg_code"`                                                                      // 用户推广码
	IsAutoTransfer      int64          `gorm:"column:is_auto_transfer;type:tinyint(1);not null;default:1;comment:是否自动转入转出 1 是 0 否" json:"is_auto_transfer" db:"is_auto_transfer"`            // 是否自动转入转出 1 是 0 否
	ThirdGameBalance    string         `gorm:"column:third_game_balance;type:varchar(256);not null;default:'';comment:用于存放哪几个第三方有余额" json:"third_game_balance" db:"third_game_balance"`   // 用于存放哪几个第三方有余额
	FromUserCode        string         `gorm:"column:from_user_code;type:varchar(100);index:idx_from_user_code,priority:1;comment:来自那个用户邀请码注册的" json:"from_user_code" db:"from_user_code"` // 来自那个用户邀请码注册的
	RecommendedCount    int64          `gorm:"column:recommended_count;type:int;comment:推荐人总数" json:"recommended_count" db:"recommended_count"`                                                   // 推荐人总数
	RecommendedName     string         `gorm:"column:recommended_name;type:varchar(50);index:index_recommended_name,priority:1;comment:推荐人姓名" json:"recommended_name" db:"recommended_name"`      // 推荐人姓名
	IP                  string         `gorm:"column:ip;type:varchar(255);index:idx_ip,priority:1" json:"ip" db:"ip"`
	RegisterIP          string         `gorm:"column:register_ip;type:varchar(255);not null;default:'';comment:注册ip地址" json:"register_ip" db:"register_ip"`                                     // 注册ip地址
	JudgeBankCard       string         `gorm:"column:judge_bank_card;type:varchar(255);not null;default:'';comment:注册彩金用户判定的银行卡" json:"judge_bank_card" db:"judge_bank_card"`           // 注册彩金用户判定的银行卡
	JudgeBankCardResult int64          `gorm:"column:judge_bank_card_result;type:tinyint;comment:注册彩金银行卡判定结果0:未通过, 1:通过" json:"judge_bank_card_result" db:"judge_bank_card_result"` // 注册彩金银行卡判定结果0:未通过, 1:通过
	Province            sql.NullString `gorm:"column:province;type:varchar(255);comment:注册所在省份" json:"province" db:"province"`                                                                // 注册所在省份
	IPArea              string         `gorm:"column:ip_area;type:varchar(100);not null;default:'';comment:注册ip所在地址" json:"ip_area" db:"ip_area"`                                             // 注册ip所在地址
	LastLoginAt         sql.NullTime   `gorm:"column:last_login_at;type:timestamp;comment:登录时间" json:"last_login_at" db:"last_login_at"`                                                        // 登录时间
	TotalOnlineDuration int64          `gorm:"column:total_online_duration;type:int unsigned;comment:累计在线时长" json:"total_online_duration" db:"total_online_duration"`                         // 累计在线时长
	OnlineDuration      int64          `gorm:"column:online_duration;type:int unsigned;comment:当前登录在线时长" json:"online_duration" db:"online_duration"`                                       // 当前登录在线时长
	Remark              string         `gorm:"column:remark;type:varchar(255);not null;default:''" json:"remark" db:"remark"`
	AutoFanshui         int64          `gorm:"column:auto_fanshui;type:tinyint;comment:是否自动领取返水 1 是 0 否" json:"auto_fanshui" db:"auto_fanshui"`                                           // 是否自动领取返水 1 是 0 否
	FirstRechargeAt     sql.NullTime   `gorm:"column:first_recharge_at;type:timestamp;index:first_recharge_at,priority:1;comment:会员首次充值时间" json:"first_recharge_at" db:"first_recharge_at"` // 会员首次充值时间
	FirstRechargeAmount float64        `gorm:"column:first_recharge_amount;type:decimal(12,4);default:0.0000;comment:会员首次充值金额" json:"first_recharge_amount" db:"first_recharge_amount"`     // 会员首次充值金额
	LastRechargeAt      sql.NullTime   `gorm:"column:last_recharge_at;type:timestamp;index:last_recharge_at,priority:1;comment:会员最后充值时间" json:"last_recharge_at" db:"last_recharge_at"`     // 会员最后充值时间
	LastRechargeAmount  float64        `gorm:"column:last_recharge_amount;type:decimal(12,4);default:0.0000;comment:会员最后充值金额" json:"last_recharge_amount" db:"last_recharge_amount"`        // 会员最后充值金额
	IconPath            sql.NullString `gorm:"column:icon_path;type:varchar(255);comment:用户头像" json:"icon_path" db:"icon_path"`                                                                 // 用户头像
	CreatedAt           sql.NullTime   `gorm:"column:created_at;type:bigint;index:created_at,priority:1;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt           sql.NullTime   `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	UserLevelUpdatedAt  sql.NullTime   `gorm:"column:user_level_updated_at;type:timestamp;index:idx_user_level_updated_at,priority:1;comment:vip晋级时间" json:"user_level_updated_at" db:"user_level_updated_at"` // vip晋级时间
	IsRisk              int64          `gorm:"column:is_risk;type:tinyint(1);comment:是否风险名单用户：1是，0否" json:"is_risk" db:"is_risk"`                                                                        // 是否风险名单用户：1是，0否
	Plat                int64          `gorm:"column:plat;type:tinyint(1);default:1;comment:终端;1:h5;2安卓;3ios4:pc" json:"plat" db:"plat"`                                                                       // 终端;1:h5;2安卓;3ios4:pc
	Mute                int64          `gorm:"column:mute;type:tinyint;comment:是否禁言0:否;1:是" json:"mute" db:"mute"`                                                                                           // 是否禁言0:否;1:是
	MaxRechargeAmount   float64        `gorm:"column:max_recharge_amount;type:decimal(16,4);not null;default:0.0000;comment:最大充值金额" json:"max_recharge_amount" db:"max_recharge_amount"`                     // 最大充值金额
	DeviceID            string         `gorm:"column:device_id;type:varchar(255);not null;default:'';comment:最后一次登录设备id" json:"device_id" db:"device_id"`                                                  // 最后一次登录设备id
	DeviceName          string         `gorm:"column:device_name;type:varchar(255);not null;default:'';comment:最后一次登录设备名称" json:"device_name" db:"device_name"`                                          // 最后一次登录设备名称
	DeviceSystemVersion string         `gorm:"column:device_system_version;type:varchar(255);not null;default:'';comment:最后一次登录设备系统版本" json:"device_system_version" db:"device_system_version"`        // 最后一次登录设备系统版本
	IsLiveStreamer      int64          `gorm:"column:is_live_streamer;type:tinyint;not null;default:0;comment:是否主播：1是，0否" json:"is_live_streamer" db:"is_live_streamer"`                                     // 是否主播：1是，0否
}

type UserES struct {
	ID                  int64   `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	Username            string  `gorm:"column:username;type:varchar(255);not null;uniqueIndex:users_name_unique,priority:1;default:''" json:"username" db:"username"`
	APIToken            string  `gorm:"column:api_token;type:varchar(500);not null;default:0" json:"api_token" db:"api_token"`
	ParentID            int64   `gorm:"column:parent_id;type:int unsigned;default:1" json:"parent_id" db:"parent_id"`
	ForefatherIds       string  `gorm:"column:forefather_ids;type:varchar(100);index:idx_forefather_ids,priority:1;default:1;comment:父亲树" json:"forefather_ids" db:"forefather_ids"` // 父亲树
	ParentName          string  `gorm:"column:parent_name;type:varchar(16)" json:"parent_name" db:"parent_name"`
	SignScore           float64 `gorm:"column:sign_score;type:decimal(12,2);not null;default:0.00;comment:签到积分" json:"sign_score" db:"sign_score"` // 签到积分
	IsOpen              int64   `gorm:"column:is_open;type:tinyint;not null;index:idx_pay_level_is_open,priority:2;index:idx_user_level_is_open,priority:2;index:is_open,priority:1;default:1" json:"is_open" db:"is_open"`
	IsAgent             int64   `gorm:"column:is_agent;type:tinyint unsigned;comment:0: 普通用户, 1: 代理;2:总代理;3:股东" json:"is_agent" db:"is_agent"`                                       // 0: 普通用户, 1: 代理;2:总代理;3:股东
	IsTester            int64   `gorm:"column:is_tester;type:tinyint unsigned;comment:0: 普通用户, 1: 测试账号" json:"is_tester" db:"is_tester"`                                                // 0: 普通用户, 1: 测试账号
	UserLevel           string  `gorm:"column:user_level;type:varchar(30);index:idx_user_level_is_open,priority:1;comment:用户层级" json:"user_level" db:"user_level"`                          // 用户层级
	PayLevel            int64   `gorm:"column:pay_level;type:varchar(100);index:idx_pay_level_is_open,priority:1;default:1;comment:支付层级 对应pay_level表" json:"pay_level" db:"pay_level"`   // 支付层级 对应pay_level表
	BankNum             int64   `gorm:"column:bank_num;type:tinyint unsigned;default:5;comment:用户能绑定的卡数" json:"bank_num" db:"bank_num"`                                                 // 用户能绑定的卡数
	IsGuest             int64   `gorm:"column:is_guest;type:tinyint unsigned;comment:是否是游客" json:"is_guest" db:"is_guest"`                                                                 // 是否是游客
	RegCode             string  `gorm:"column:reg_code;type:varchar(60);comment:用户推广码" json:"reg_code" db:"reg_code"`                                                                      // 用户推广码
	IsAutoTransfer      int64   `gorm:"column:is_auto_transfer;type:tinyint(1);not null;default:1;comment:是否自动转入转出 1 是 0 否" json:"is_auto_transfer" db:"is_auto_transfer"`            // 是否自动转入转出 1 是 0 否
	ThirdGameBalance    string  `gorm:"column:third_game_balance;type:varchar(256);not null;default:'';comment:用于存放哪几个第三方有余额" json:"third_game_balance" db:"third_game_balance"`   // 用于存放哪几个第三方有余额
	FromUserCode        string  `gorm:"column:from_user_code;type:varchar(100);index:idx_from_user_code,priority:1;comment:来自那个用户邀请码注册的" json:"from_user_code" db:"from_user_code"` // 来自那个用户邀请码注册的
	RecommendedCount    int64   `gorm:"column:recommended_count;type:int;comment:推荐人总数" json:"recommended_count" db:"recommended_count"`                                                   // 推荐人总数
	RecommendedName     string  `gorm:"column:recommended_name;type:varchar(50);index:index_recommended_name,priority:1;comment:推荐人姓名" json:"recommended_name" db:"recommended_name"`      // 推荐人姓名
	IP                  string  `gorm:"column:ip;type:varchar(255);index:idx_ip,priority:1" json:"ip" db:"ip"`
	RegisterIP          string  `gorm:"column:register_ip;type:varchar(255);not null;default:'';comment:注册ip地址" json:"register_ip" db:"register_ip"`                                     // 注册ip地址
	JudgeBankCard       string  `gorm:"column:judge_bank_card;type:varchar(255);not null;default:'';comment:注册彩金用户判定的银行卡" json:"judge_bank_card" db:"judge_bank_card"`           // 注册彩金用户判定的银行卡
	JudgeBankCardResult int64   `gorm:"column:judge_bank_card_result;type:tinyint;comment:注册彩金银行卡判定结果0:未通过, 1:通过" json:"judge_bank_card_result" db:"judge_bank_card_result"` // 注册彩金银行卡判定结果0:未通过, 1:通过
	Province            string  `gorm:"column:province;type:varchar(255);comment:注册所在省份" json:"province" db:"province"`                                                                // 注册所在省份
	IPArea              string  `gorm:"column:ip_area;type:varchar(100);not null;default:'';comment:注册ip所在地址" json:"ip_area" db:"ip_area"`                                             // 注册ip所在地址
	LastLoginAt         string  `gorm:"column:last_login_at;type:timestamp;comment:登录时间" json:"last_login_at" db:"last_login_at"`                                                        // 登录时间
	TotalOnlineDuration int64   `gorm:"column:total_online_duration;type:int unsigned;comment:累计在线时长" json:"total_online_duration" db:"total_online_duration"`                         // 累计在线时长
	OnlineDuration      int64   `gorm:"column:online_duration;type:int unsigned;comment:当前登录在线时长" json:"online_duration" db:"online_duration"`                                       // 当前登录在线时长
	Remark              string  `gorm:"column:remark;type:varchar(255);not null;default:''" json:"remark" db:"remark"`
	AutoFanshui         int64   `gorm:"column:auto_fanshui;type:tinyint;comment:是否自动领取返水 1 是 0 否" json:"auto_fanshui" db:"auto_fanshui"`                                           // 是否自动领取返水 1 是 0 否
	FirstRechargeAt     string  `gorm:"column:first_recharge_at;type:timestamp;index:first_recharge_at,priority:1;comment:会员首次充值时间" json:"first_recharge_at" db:"first_recharge_at"` // 会员首次充值时间
	FirstRechargeAmount float64 `gorm:"column:first_recharge_amount;type:decimal(12,4);default:0.0000;comment:会员首次充值金额" json:"first_recharge_amount" db:"first_recharge_amount"`     // 会员首次充值金额
	LastRechargeAt      string  `gorm:"column:last_recharge_at;type:timestamp;index:last_recharge_at,priority:1;comment:会员最后充值时间" json:"last_recharge_at" db:"last_recharge_at"`     // 会员最后充值时间
	LastRechargeAmount  float64 `gorm:"column:last_recharge_amount;type:decimal(12,4);default:0.0000;comment:会员最后充值金额" json:"last_recharge_amount" db:"last_recharge_amount"`        // 会员最后充值金额
	IconPath            string  `gorm:"column:icon_path;type:varchar(255);comment:用户头像" json:"icon_path" db:"icon_path"`                                                                 // 用户头像
	CreatedAt           string  `gorm:"column:created_at;type:bigint;index:created_at,priority:1;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt           string  `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	UserLevelUpdatedAt  string  `gorm:"column:user_level_updated_at;type:timestamp;index:idx_user_level_updated_at,priority:1;comment:vip晋级时间" json:"user_level_updated_at" db:"user_level_updated_at"` // vip晋级时间
	IsRisk              int64   `gorm:"column:is_risk;type:tinyint(1);comment:是否风险名单用户：1是，0否" json:"is_risk" db:"is_risk"`                                                                        // 是否风险名单用户：1是，0否
	Plat                int64   `gorm:"column:plat;type:tinyint(1);default:1;comment:终端;1:h5;2安卓;3ios4:pc" json:"plat" db:"plat"`                                                                       // 终端;1:h5;2安卓;3ios4:pc
	Mute                int64   `gorm:"column:mute;type:tinyint;comment:是否禁言0:否;1:是" json:"mute" db:"mute"`                                                                                           // 是否禁言0:否;1:是
	MaxRechargeAmount   float64 `gorm:"column:max_recharge_amount;type:decimal(16,4);not null;default:0.0000;comment:最大充值金额" json:"max_recharge_amount" db:"max_recharge_amount"`                     // 最大充值金额
	DeviceID            string  `gorm:"column:device_id;type:varchar(255);not null;default:'';comment:最后一次登录设备id" json:"device_id" db:"device_id"`                                                  // 最后一次登录设备id
	DeviceName          string  `gorm:"column:device_name;type:varchar(255);not null;default:'';comment:最后一次登录设备名称" json:"device_name" db:"device_name"`                                          // 最后一次登录设备名称
	DeviceSystemVersion string  `gorm:"column:device_system_version;type:varchar(255);not null;default:'';comment:最后一次登录设备系统版本" json:"device_system_version" db:"device_system_version"`        // 最后一次登录设备系统版本
	IsLiveStreamer      int64   `gorm:"column:is_live_streamer;type:tinyint;not null;default:0;comment:是否主播：1是，0否" json:"is_live_streamer" db:"is_live_streamer"`                                     // 是否主播：1是，0否
}
type UserSuccessRechargeData struct {
	ID               int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:自增id" json:"id"`                                  // 自增id
	UserID           int64  `gorm:"column:user_id;type:int unsigned;not null;index:idx_date_at_user_id,priority:2;default:0;comment:用户id" json:"user_id"` // 用户id
	UserPayLevel     int64  `gorm:"column:user_pay_level;type:int;comment:用户支付层级ID 对应pay_levels表id" json:"user_pay_level"`                         // 用户支付层级ID 对应pay_levels表id
	UserPayLevelName string `gorm:"column:user_pay_level_name;type:varchar(100);comment:用户支付层级名称" json:"user_pay_level_name"`                       // 用户支付层级名称
}

type MessageJob struct {
	ID        int64        `db:"id"`
	MsgID     int64        `db:"msg_id"`   // 消息类型
	MsgName   string       `db:"msg_name"` // 消息名称
	Data      string       `db:"data"`     // 队列数据
	Status    int64        `db:"status"`   // 队列执行状态：0未执行，1失败，2成功
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type Data struct {
	AdId                      int64  `json:"ad_id"`
	AutoCreateFrontCondition  int64  `json:"auto_create_front_condition"`
	BatchType                 int64  `json:"batch_type"`
	BetAmountRate             string `json:"bet_amount_rate"`
	BonusAmount               string `json:"bonus_amount"`
	Content                   string `json:"content"`
	SendPayLevel              string `json:"send_pay_level"`
	SendUsernames             string `json:"send_usernames"`
	SendVipLevel              string `json:"send_vip_level"`
	Title                     string `json:"title"`
	AdminName                 string `json:"admin_name"`
	PrivateMessageType        int64  `json:"private_message_type"`
	TotalUser                 int64  `json:"total_user"`
	Type                      int64  `json:"type"`
	RedEnvelopeExpirationTime int64  `json:"red_envelope_expiration_time"`
	IsOpenPush                int64  `json:"is_open_push"`
	SendPlat                  string `json:"send_plat"`
}
type MessageInsertData struct {
	Title              string `gorm:"column:title;type:varchar(255);not null;default:'';comment:消息标题" json:"title"`                                                                                        // 消息标题
	Content            string `gorm:"column:content;type:text;comment:消息内容" json:"content"`                                                                                                                // 消息内容
	MsgID              int64  `gorm:"column:msg_id;type:tinyint unsigned;not null;index:idx_messages_main_index,priority:1;index:messages_msg_id_index,priority:1;default:0;comment:消息类型id" json:"msg_id"` // 消息类型id
	MsgName            string `gorm:"column:msg_name;type:varchar(255);not null;default:'';comment:消息类型标题" json:"msg_name"`                                                                              // 消息类型标题
	FromUserid         int64  `gorm:"column:from_userid;type:int unsigned;comment:发送者用户id" json:"from_userid"`                                                                                            // 发送者用户id
	FromUsername       string `gorm:"column:from_username;type:varchar(255);index:messages_from_username_index,priority:1;comment:发件人" json:"from_username"`                                                // 发件人
	AdminName          string `gorm:"column:admin_name;type:varchar(255);comment:添加消息的后台管理员" json:"admin_name"`                                                                                      // 添加消息的后台管理员
	CreatedAt          string `gorm:"column:created_at;type:bigint;index:idx_messages_main_index,priority:2;default:CURRENT_TIMESTAMP;autoCreateTime:milli" json:"created_at"`
	UpdatedAt          string `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at"`
	Sort               int64  `gorm:"column:sort;type:int;not null;default:0;comment:权重" json:"sort"` // 权重
	PrivateMessageType int64  `json:"private_message_type"`
}
type ReplyInsert struct {
	Type                      int64  `json:"type"`     // 1:消息的回复
	Content                   string `json:"content"`  // 回复内容
	FromUID                   int64  `json:"from_uid"` // 评论用户id
	Status                    int64  `json:"status"`   // 0:未启用, 1:启用
	RedEnvelopeExpirationTime int64  `json:"red_envelope_expiration_time"`
	BonusAmount               string `json:"bonus_amount"`
	PrivateMessageType        int64  `json:"private_message_type"`
	BetAmountRate             string `json:"bet_amount_rate"`
}
type MessageJobData struct {
	Data              Data              `json:"data"`
	MessageInsertData MessageInsertData `json:"message_insert_data"`
	//ReplyInsert       ReplyInsert       `json:"reply_insert"`
}

type UserLevel struct {
	ID        int    `db:"id" json:"id"`
	UserLevel string `db:"user_level" json:"user_level"`
}

type EsUser struct {
	ID                  int64     `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Username            string    `gorm:"column:username;type:varchar(255);not null;uniqueIndex:users_name_unique,priority:1;default:''" json:"username"`
	Password            string    `gorm:"column:password;type:varchar(255);not null;default:''" json:"password"`
	APIToken            string    `gorm:"column:api_token;type:varchar(500);not null;default:0" json:"api_token"`
	PayPassword         string    `gorm:"column:pay_password;type:varchar(255)" json:"pay_password"`
	PaySecret           string    `gorm:"column:pay_secret;type:varchar(20);comment:支付明文密码" json:"pay_secret"` // 支付明文密码
	ParentID            int64     `gorm:"column:parent_id;type:int unsigned;default:1" json:"parent_id"`
	ForefatherIds       string    `gorm:"column:forefather_ids;type:varchar(100);index:idx_forefather_ids,priority:1;default:1;comment:父亲树" json:"forefather_ids"` // 父亲树
	ParentName          string    `gorm:"column:parent_name;type:varchar(16)" json:"parent_name"`
	SignScore           float64   `gorm:"column:sign_score;type:decimal(12,2);not null;default:0.00;comment:签到积分" json:"sign_score"` // 签到积分
	IsOpen              int64     `gorm:"column:is_open;type:tinyint;not null;index:idx_pay_level_is_open,priority:2;index:idx_user_level_is_open,priority:2;index:is_open,priority:1;default:1" json:"is_open"`
	IsAgent             int64     `gorm:"column:is_agent;type:tinyint unsigned;comment:0: 普通用户, 1: 代理;2:总代理;3:股东" json:"is_agent"`                                    // 0: 普通用户, 1: 代理;2:总代理;3:股东
	IsTester            int64     `gorm:"column:is_tester;type:tinyint unsigned;comment:0: 普通用户, 1: 测试账号" json:"is_tester"`                                              // 0: 普通用户, 1: 测试账号
	UserLevel           string    `gorm:"column:user_level;type:varchar(30);index:idx_user_level_is_open,priority:1;comment:用户层级" json:"user_level"`                         // 用户层级
	PayLevel            string    `gorm:"column:pay_level;type:varchar(100);index:idx_pay_level_is_open,priority:1;default:1;comment:支付层级 对应pay_level表" json:"pay_level"` // 支付层级 对应pay_level表
	BankNum             int64     `gorm:"column:bank_num;type:tinyint unsigned;default:5;comment:用户能绑定的卡数" json:"bank_num"`                                              // 用户能绑定的卡数
	IsGuest             int64     `gorm:"column:is_guest;type:tinyint unsigned;comment:是否是游客" json:"is_guest"`                                                              // 是否是游客
	RegCode             string    `gorm:"column:reg_code;type:varchar(60);comment:用户推广码" json:"reg_code"`                                                                   // 用户推广码
	IsAutoTransfer      int64     `gorm:"column:is_auto_transfer;type:tinyint(1);not null;default:1;comment:是否自动转入转出 1 是 0 否" json:"is_auto_transfer"`                 // 是否自动转入转出 1 是 0 否
	ThirdGameBalance    string    `gorm:"column:third_game_balance;type:varchar(256);not null;default:'';comment:用于存放哪几个第三方有余额" json:"third_game_balance"`          // 用于存放哪几个第三方有余额
	FromUserCode        string    `gorm:"column:from_user_code;type:varchar(100);index:idx_from_user_code,priority:1;comment:来自那个用户邀请码注册的" json:"from_user_code"`    // 来自那个用户邀请码注册的
	RecommendedCount    int64     `gorm:"column:recommended_count;type:int;comment:推荐人总数" json:"recommended_count"`                                                         // 推荐人总数
	RecommendedName     string    `gorm:"column:recommended_name;type:varchar(50);index:index_recommended_name,priority:1;comment:推荐人姓名" json:"recommended_name"`           // 推荐人姓名
	IP                  string    `gorm:"column:ip;type:varchar(255);index:idx_ip,priority:1" json:"ip"`
	RegisterIP          string    `gorm:"column:register_ip;type:varchar(255);not null;default:'';comment:注册ip地址" json:"register_ip"`                          // 注册ip地址
	JudgeBankCard       string    `gorm:"column:judge_bank_card;type:varchar(255);not null;default:'';comment:注册彩金用户判定的银行卡" json:"judge_bank_card"`    // 注册彩金用户判定的银行卡
	JudgeBankCardResult int64     `gorm:"column:judge_bank_card_result;type:tinyint;comment:注册彩金银行卡判定结果0:未通过, 1:通过" json:"judge_bank_card_result"` // 注册彩金银行卡判定结果0:未通过, 1:通过
	Province            string    `gorm:"column:province;type:varchar(255);comment:注册所在省份" json:"province"`                                                  // 注册所在省份
	IPArea              string    `gorm:"column:ip_area;type:varchar(100);not null;default:'';comment:注册ip所在地址" json:"ip_area"`                              // 注册ip所在地址
	LastLoginAt         time.Time `gorm:"column:last_login_at;type:timestamp;comment:登录时间" json:"last_login_at"`                                               // 登录时间
	TotalOnlineDuration int64     `gorm:"column:total_online_duration;type:int unsigned;comment:累计在线时长" json:"total_online_duration"`                        // 累计在线时长
	OnlineDuration      int64     `gorm:"column:online_duration;type:int unsigned;comment:当前登录在线时长" json:"online_duration"`                                // 当前登录在线时长
	Remark              string    `gorm:"column:remark;type:varchar(255);not null;default:''" json:"remark"`
	AutoFanshui         int64     `gorm:"column:auto_fanshui;type:tinyint;comment:是否自动领取返水 1 是 0 否" json:"auto_fanshui"`                                      // 是否自动领取返水 1 是 0 否
	FirstRechargeAt     time.Time `gorm:"column:first_recharge_at;type:timestamp;index:first_recharge_at,priority:1;comment:会员首次充值时间" json:"first_recharge_at"` // 会员首次充值时间
	FirstRechargeAmount float64   `gorm:"column:first_recharge_amount;type:decimal(12,4);default:0.0000;comment:会员首次充值金额" json:"first_recharge_amount"`         // 会员首次充值金额
	LastRechargeAt      time.Time `gorm:"column:last_recharge_at;type:timestamp;index:last_recharge_at,priority:1;comment:会员最后充值时间" json:"last_recharge_at"`    // 会员最后充值时间
	LastRechargeAmount  float64   `gorm:"column:last_recharge_amount;type:decimal(12,4);default:0.0000;comment:会员最后充值金额" json:"last_recharge_amount"`           // 会员最后充值金额
	IconPath            string    `gorm:"column:icon_path;type:varchar(255);comment:用户头像" json:"icon_path"`                                                         // 用户头像
	CreatedAt           time.Time `gorm:"column:created_at;type:bigint;index:created_at,priority:1;autoCreateTime:milli" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at"`
	UserLevelUpdatedAt  time.Time `gorm:"column:user_level_updated_at;type:timestamp;index:idx_user_level_updated_at,priority:1;comment:vip晋级时间" json:"user_level_updated_at"` // vip晋级时间
	IsRisk              int64     `gorm:"column:is_risk;type:tinyint(1);comment:是否风险名单用户：1是，0否" json:"is_risk"`                                                          // 是否风险名单用户：1是，0否
	NewPassword         string    `gorm:"column:new_password;type:varchar(255);comment:hash密码" json:"new_password"`                                                              // hash密码
	Plat                int64     `gorm:"column:plat;type:tinyint(1);default:1;comment:终端;1:h5;2安卓;3ios4:pc" json:"plat"`                                                      // 终端;1:h5;2安卓;3ios4:pc
	Mute                int64     `gorm:"column:mute;type:tinyint;comment:是否禁言0:否;1:是" json:"mute"`                                                                          // 是否禁言0:否;1:是
	MaxRechargeAmount   float64   `gorm:"column:max_recharge_amount;type:decimal(16,4);not null;default:0.0000;comment:最大充值金额" json:"max_recharge_amount"`                   // 最大充值金额
	DeviceID            string    `gorm:"column:device_id;type:varchar(255);not null;default:'';comment:最后一次登录设备id" json:"device_id"`                                      // 最后一次登录设备id
	DeviceName          string    `gorm:"column:device_name;type:varchar(255);not null;default:'';comment:最后一次登录设备名称" json:"device_name"`                                // 最后一次登录设备名称
	DeviceSystemVersion string    `gorm:"column:device_system_version;type:varchar(255);not null;default:'';comment:最后一次登录设备系统版本" json:"device_system_version"`        // 最后一次登录设备系统版本
	IsLiveStreamer      int64     `gorm:"column:is_live_streamer;type:tinyint;not null;default:0;comment:是否主播：1是，0否" json:"is_live_streamer"`                                // 是否主播：1是，0否
}

type GetTransListByTransFatherIdParams struct {
	TransFatherID int64 `json:"trans_father_id"` // 账变类型父类
	StartAt       int64 `json:"start_at"`
	EndAt         int64 `json:"end_at"`
}

type TransAggResult struct {
	UniqueUserCount int64
	SumAmount       float64
	UserIDs         []int64
}

type Message struct {
	ID           int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	Title        string `gorm:"column:title;type:varchar(255);not null;default:'';comment:消息标题" json:"title" db:"title"`                                                                                         // 消息标题
	Content      string `gorm:"column:content;type:text;comment:消息内容" json:"content" db:"content"`                                                                                                               // 消息内容
	MsgID        int64  `gorm:"column:msg_id;type:tinyint unsigned;not null;index:idx_messages_main_index,priority:1;index:messages_msg_id_index,priority:1;default:0;comment:消息类型id" json:"msg_id" db:"msg_id"` // 消息类型id
	MsgName      string `gorm:"column:msg_name;type:varchar(255);not null;default:'';comment:消息类型标题" json:"msg_name" db:"msg_name"`                                                                            // 消息类型标题
	FromUserid   int64  `gorm:"column:from_userid;type:int unsigned;comment:发送者用户id" json:"from_userid" db:"from_userid"`                                                                                       // 发送者用户id
	FromUsername string `gorm:"column:from_username;type:varchar(255);index:messages_from_username_index,priority:1;comment:发件人" json:"from_username" db:"from_username"`                                         // 发件人
	UserID       int64  `gorm:"column:user_id;type:int unsigned;index:idx_messages_main_index,priority:3;comment:收件用户id" json:"user_id" db:"user_id"`                                                            // 收件用户id
	Username     string `gorm:"column:username;type:varchar(255);index:messages_to_username_index,priority:1;comment:收件人" json:"username" db:"username"`                                                          // 收件人
	PayLevelID   string `gorm:"column:pay_levels;type:varchar(256);comment:指定层级（逗号分隔)" json:"pay_level_id" db:"pay_level_id"`                                                                                // 指定层级（逗号分隔)
	IsReceived   int    `json:"is_received" db:"is_received"`
	UserLevelID  int    `json:"user_level_id" db:"user_level_id"`
	MessageJobID int64  `json:"message_job_id" db:"message_job_id"`
	AgentNames   string `gorm:"column:agent_names;type:varchar(256);comment:指定代理（逗号分隔）" json:"agent_names" db:"agent_names"` // 指定代理（逗号分隔）
	MsgType      int64  `gorm:"column:msg_type;type:tinyint;comment:消息主体类型 0文本 1图片 2视频" json:"msg_type" db:"msg_type"`   // 消息主体类型 0文本 1图片 2视频
	Path         string `gorm:"column:path;type:json;comment:图片、视频地址" json:"path" db:"path"`                                   // 图片、视频地址
}

type PGMessage struct {
	UserID       int64  `gorm:"column:user_id;type:int unsigned;index:idx_messages_main_index,priority:3;comment:收件用户id" json:"user_id" db:"user_id"`   // 收件用户id
	Username     string `gorm:"column:username;type:varchar(255);index:messages_to_username_index,priority:1;comment:收件人" json:"username" db:"username"` // 收件人
	PayLevelID   string `gorm:"column:pay_levels;type:varchar(256);comment:指定层级（逗号分隔)" json:"pay_level_id" db:"pay_level_id"`                       // 指定层级（逗号分隔)
	IsReceived   int    `json:"is_received" db:"is_received"`
	UserLevelID  int    `json:"user_level_id" db:"user_level_id"`
	MessageJobID int64  `json:"message_job_id" db:"message_job_id"`
}

type CacheUpdate struct {
	CacheKey string
	Time     int64
}

type ReplyContent struct {
	//ID       int64 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	ObjectID int64 `gorm:"column:object_id;type:bigint unsigned;index:idx_object_id,priority:1;comment:对象id" json:"object_id" db:"object_id"` // 对象id
	//Type         int64     `gorm:"column:type;type:tinyint unsigned;comment:1:消息的回复" json:"type" db:"type"`                                           // 1:消息的回复
	//Content      string    `gorm:"column:content;type:varchar(800);comment:回复内容" json:"content" db:"content"`                                         // 回复内容
	//FromUID      int64     `gorm:"column:from_uid;type:int unsigned;comment:评论用户id" json:"from_uid" db:"from_uid"`                                    // 评论用户id
	ToUID int64 `gorm:"column:to_uid;type:int unsigned;comment:评论目标用户id(对某用户的评论)" json:"to_uid" db:"to_uid"` // 评论目标用户id(对某用户的评论)
	//Status       int64     `gorm:"column:status;type:tinyint unsigned;comment:0:未启用, 1:启用" json:"status" db:"status"`                                 // 0:未启用, 1:启用
	//CreatedAt    time.Time `gorm:"column:created_at;type:bigint;index:idx_created_at,priority:1;autoCreateTime:milli" json:"created_at" db:"created_at"`
	//UpdatedAt    time.Time  `gorm:"column:updated_at;type:bigint;not null;default:CURRENT_TIMESTAMP;autoUpdateTime:milli;comment:更新时间" json:"updated_at" db:"updated_at"` // 更新时间
	//MsgType      int64     `gorm:"column:msg_type;type:tinyint;comment:消息主体类型 0文本 1图片 2视频" json:"msg_type" db:"msg_type"`                                                // 消息主体类型 0文本 1图片 2视频
	//Path         string    `gorm:"column:path;type:json;comment:图片、视频地址" json:"path" db:"path"`                                                                          // 图片、视频地址
	IsReceived   int64 `json:"is_received" db:"is_received"`
	MessageJobID int64 `json:"message_job_id" db:"message_job_id"`
}

type MessageType struct {
	ID        int64     `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" db:"id"`
	ParentID  int64     `gorm:"column:parent_id;type:bigint;comment:父分类id" db:"parent_id"`                             // 父分类id
	EnName    string    `gorm:"column:en_name;type:varchar(255);not null;default:'';comment:类型名称 英文" db:"en_name"`  // 类型名称 英文
	CnName    string    `gorm:"column:cn_name;type:varchar(255);not null;default:'';comment:类型名称 中文" db:"cn_name"`  // 类型名称 中文
	IsOpen    int64     `gorm:"column:is_open;type:tinyint(1);not null;default:1;comment:是否开启0:否;1:是" db:"is_open"` // 是否开启0:否;1:是
	Plat      int64     `gorm:"column:plat;type:tinyint;default:1;comment:1:手机端;2:pc" db:"plat"`                       // 1:手机端;2:pc
	AppPlat   int64     `gorm:"column:app_plat;type:tinyint;default:1;comment:1:综合app;2:体育app" db:"app_plat"`         // 1:综合app;2:体育app
	IsBtn     int64     `gorm:"column:is_btn;type:tinyint(1);default:1;comment: 是否tap按钮 0:否 1:是" db:"is_btn"`       //  是否tap按钮 0:否 1:是
	Sort      int64     `gorm:"column:sort;type:int;comment:排序" db:"sort"`                                              // 排序
	Remark    string    `gorm:"column:remark;type:varchar(255);comment:备注" db:"remark"`                                 // 备注
	CreatedAt time.Time `gorm:"column:created_at;type:bigint;autoCreateTime:milli" db:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" db:"updated_at"`
}

type MessagePushBatch struct {
	//ID          int64  `json:"id" db:"id"`
	ObjectID    int64  `json:"object_id" db:"object_id"`
	MsgTypeID   int64  `json:"msg_type_id" db:"msg_type_id"`
	RelatedType int64  `json:"related_type" db:"related_type"`
	RelatedID   int64  `json:"related_id" db:"related_id"`
	UserID      int64  `json:"user_id" db:"user_id"`
	Username    string `json:"username" db:"username"`
	Title       string `json:"title" db:"title"`
	Content     string `json:"content" db:"content"`
	SendPlats   string `json:"send_plats" db:"send_plats"`
}

type UserEngagelabRegisterID struct {
	ID             int64  `json:"id" db:"id"`
	UserID         int64  `json:"user_id" db:"user_id"`
	Username       string `json:"username" db:"username"`
	Plat           int64  `json:"plat" db:"plat"`                       // 终端设备
	RegistrationID string `json:"registration_id" db:"registration_id"` // 注册id
	//CreatedAt      time.Time `json:"created_at" db:"created_at"`
	//UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	SendPlats string `json:"send_plats" db:"send_plats"`
	Host      string `json:"host" db:"host"`
}

type MessagePushRecord struct {
	ID        int64     `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	ObjectID  int64     `json:"object_id"`
	BatchID   int64     `json:"batch_id"`
	MessageID int64     `json:"message_id"`
	PushAt    time.Time `json:"push_at"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Plat      int64     `json:"plat"`
}
type HttpResponse struct {
	Plat  int64 `json:"plat"`
	MsgID int64 `json:"msg_id"`
}

type AppRequest struct {
	From       string     `json:"from"`
	To         ToField    `json:"to"`
	Body       Body       `json:"body"`
	CustomArgs CustomArgs `json:"custom_args"`
}
type Body struct {
	Platform     string       `json:"platform"`
	Notification Notification `json:"notification"`
}

type Notification struct {
	Android Android `json:"android"`
	Ios     Ios     `json:"ios"`
}

type WebBody struct {
	Platform     string          `json:"platform"`
	Notification WebNotification `json:"notification"`
}
type WebNotification struct {
	Web Web `json:"web"`
}

type Web struct {
	Alert  string    `json:"alert"`
	Title  string    `json:"title"`
	Url    string    `json:"url"`
	Extras WebExtras `json:"extras"`
}
type Android struct {
	Alert     string `json:"alert"`
	Title     string `json:"title"`
	SmallIcon string `json:"small_icon"`
	Extras    Extras `json:"extras"`
	Intent    Intent `json:"intent"`
}

type Extras struct {
	JumpUrl string `json:"jump_url"`
}

type Intent struct {
	Url string `json:"url"`
}
type Ios struct {
	Alert  Alert  `json:"alert"`
	Extras Extras `json:"extras"`
}
type Alert struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
type CustomArgs struct {
	PushCode      string `json:"push_code"`
	PushBatchCode string `json:"push_batch_code"`
}

type WebExtras struct {
	MsgID   int64  `json:"msg_id"`
	MsgType string `json:"msg_type"`
	Host    string `json:"host"`
	Plat    int64  `json:"plat"`
}
type WebRequest struct {
	From       string     `json:"from"`
	To         ToField    `json:"to"`
	Body       WebBody    `json:"body"`
	CustomArgs CustomArgs `json:"custom_args"`
}

type ToField struct {
	RegistrationIDs []string `json:"registration_id"`
}

type TransactionType struct {
	ID             int64        `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	ParentID       int64        `gorm:"column:parent_id;type:int unsigned;index:parent_id,priority:1" json:"parent_id" db:"parent_id"`
	CnTitle        string       `gorm:"column:cn_title;type:varchar(30)" json:"cn_title" db:"cn_title"`
	EnTitle        string       `gorm:"column:en_title;type:varchar(30)" json:"en_title" db:"en_title"`
	Description    string       `gorm:"column:description;type:varchar(30);uniqueIndex:description,priority:1" json:"description" db:"description"`
	Amount         int64        `gorm:"column:amount;type:tinyint;not null;default:0" json:"amount" db:"amount"`
	Deposit        int64        `gorm:"column:deposit;type:tinyint;comment:充值" json:"deposit" db:"deposit"`                                           // 充值
	DepositTimes   int64        `gorm:"column:deposit_times;type:int;comment:总的充值次数" json:"deposit_times" db:"deposit_times"`                     // 总的充值次数
	Withdraw       int64        `gorm:"column:withdraw;type:tinyint;comment:提款" json:"withdraw" db:"withdraw"`                                        // 提款
	WithdrawTimes  int64        `gorm:"column:withdraw_times;type:int;comment:总的提款次数" json:"withdraw_times" db:"withdraw_times"`                  // 总的提款次数
	TodayProfit    int64        `gorm:"column:today_profit;type:tinyint;comment:0:不统计今日盈亏;1:统计今日盈亏" json:"today_profit" db:"today_profit"` // 0:不统计今日盈亏;1:统计今日盈亏
	BetAmount      int64        `gorm:"column:bet_amount;type:tinyint;comment:0:不统计打码;1:统计打码" json:"bet_amount" db:"bet_amount"`               // 0:不统计打码;1:统计打码
	Tuijian        int64        `gorm:"column:tuijian;type:tinyint;comment:推荐奖励" json:"tuijian" db:"tuijian"`                                       // 推荐奖励
	AgentYongjin   int64        `gorm:"column:agent_yongjin;type:tinyint;comment:是否统计代理佣金0:否；1：统计" json:"agent_yongjin" db:"agent_yongjin"`  // 是否统计代理佣金0:否；1：统计
	IncomeMaxLimit int64        `gorm:"column:income_max_limit;type:tinyint;comment:充值时单笔限额" json:"income_max_limit" db:"income_max_limit"`      // 充值时单笔限额
	PayLevel       int64        `gorm:"column:pay_level;type:tinyint;comment:是否需要判断支付层级升级" json:"pay_level" db:"pay_level"`                 // 是否需要判断支付层级升级
	IsOpen         int64        `gorm:"column:is_open;type:tinyint unsigned;default:1;comment:是否开启0:否1：是" json:"is_open" db:"is_open"`            // 是否开启0:否1：是
	CreatedAt      sql.NullTime `gorm:"column:created_at;type:bigint;index:created_at,priority:1;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt      sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
}

type TransactionTypeEs struct {
	ID             int64  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	ParentID       int64  `gorm:"column:parent_id;type:int unsigned;index:parent_id,priority:1" json:"parent_id" db:"parent_id"`
	CnTitle        string `gorm:"column:cn_title;type:varchar(30)" json:"cn_title" db:"cn_title"`
	EnTitle        string `gorm:"column:en_title;type:varchar(30)" json:"en_title" db:"en_title"`
	Description    string `gorm:"column:description;type:varchar(30);uniqueIndex:description,priority:1" json:"description" db:"description"`
	Amount         int64  `gorm:"column:amount;type:tinyint;not null;default:0" json:"amount" db:"amount"`
	Deposit        int64  `gorm:"column:deposit;type:tinyint;comment:充值" json:"deposit" db:"deposit"`                                           // 充值
	DepositTimes   int64  `gorm:"column:deposit_times;type:int;comment:总的充值次数" json:"deposit_times" db:"deposit_times"`                     // 总的充值次数
	Withdraw       int64  `gorm:"column:withdraw;type:tinyint;comment:提款" json:"withdraw" db:"withdraw"`                                        // 提款
	WithdrawTimes  int64  `gorm:"column:withdraw_times;type:int;comment:总的提款次数" json:"withdraw_times" db:"withdraw_times"`                  // 总的提款次数
	TodayProfit    int64  `gorm:"column:today_profit;type:tinyint;comment:0:不统计今日盈亏;1:统计今日盈亏" json:"today_profit" db:"today_profit"` // 0:不统计今日盈亏;1:统计今日盈亏
	BetAmount      int64  `gorm:"column:bet_amount;type:tinyint;comment:0:不统计打码;1:统计打码" json:"bet_amount" db:"bet_amount"`               // 0:不统计打码;1:统计打码
	Tuijian        int64  `gorm:"column:tuijian;type:tinyint;comment:推荐奖励" json:"tuijian" db:"tuijian"`                                       // 推荐奖励
	AgentYongjin   int64  `gorm:"column:agent_yongjin;type:tinyint;comment:是否统计代理佣金0:否；1：统计" json:"agent_yongjin" db:"agent_yongjin"`  // 是否统计代理佣金0:否；1：统计
	IncomeMaxLimit int64  `gorm:"column:income_max_limit;type:tinyint;comment:充值时单笔限额" json:"income_max_limit" db:"income_max_limit"`      // 充值时单笔限额
	PayLevel       int64  `gorm:"column:pay_level;type:tinyint;comment:是否需要判断支付层级升级" json:"pay_level" db:"pay_level"`                 // 是否需要判断支付层级升级
	IsOpen         int64  `gorm:"column:is_open;type:tinyint unsigned;default:1;comment:是否开启0:否1：是" json:"is_open" db:"is_open"`            // 是否开启0:否1：是
	CreatedAt      string `gorm:"column:created_at;type:bigint;index:created_at,priority:1;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt      string `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
}

type Transaction struct {
	ID                int64        `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID            int64        `gorm:"column:user_id;type:int unsigned;not null;default:0" json:"user_id" db:"user_id"`
	Username          string       `gorm:"column:username;type:varchar(255);not null;index:idx_uname,priority:1;default:''" json:"username" db:"username"`
	IsTester          int64        `gorm:"column:is_tester;type:tinyint(1)" json:"is_tester" db:"is_tester"`
	ParentID          int64        `gorm:"column:parent_id;type:int;not null;default:0;comment:上级id" json:"parent_id" db:"parent_id"` // 上级id
	ForefatherIds     string       `gorm:"column:forefather_ids;type:varchar(1024)" json:"forefather_ids" db:"forefather_ids"`
	Amount            float64      `gorm:"column:amount;type:decimal(14,2);not null;index:idx_amount,priority:1;default:0.00;comment:用户当前交易金额" json:"amount" db:"amount"`          // 用户当前交易金额
	BankType          int64        `gorm:"column:bank_type;type:tinyint;comment:1:银行卡, 2:数字钱包, 3:易汇钱包, 4:支付宝, 5:微信, 6:数字人民币" json:"bank_type" db:"bank_type"`         // 1:银行卡, 2:数字钱包, 3:易汇钱包, 4:支付宝, 5:微信, 6:数字人民币
	TransFatherID     int64        `gorm:"column:trans_father_id;type:int unsigned;index:idx_trans_father_id,priority:1;comment:账变类型父类" json:"trans_father_id" db:"trans_father_id"` // 账变类型父类
	TransTypesID      int64        `gorm:"column:trans_types_id;type:mediumint unsigned;not null;index:act_id,priority:1;default:0" json:"trans_types_id" db:"trans_types_id"`
	TransTypesCnTitle string       `gorm:"column:trans_types_cn_title;type:varchar(30)" json:"trans_types_cn_title" db:"trans_types_cn_title"`
	TransTypesEnTitle string       `gorm:"column:trans_types_en_title;type:varchar(50);not null;default:''" json:"trans_types_en_title" db:"trans_types_en_title"`
	IsIncome          int64        `gorm:"column:is_income;type:tinyint(1);comment:是否入款" json:"is_income" db:"is_income"` // 是否入款
	BeforeMoney       float64      `gorm:"column:before_money;type:decimal(12,2);not null;default:0.00" json:"before_money" db:"before_money"`
	Money             float64      `gorm:"column:money;type:decimal(12,2);not null;default:0.00;comment:交易后金额" json:"money" db:"money"`               // 交易后金额
	GameCode          string       `gorm:"column:game_code;type:varchar(30);index:lottery_id,priority:1;comment:游戏game" json:"game_code" db:"game_code"` // 游戏game
	Issue             string       `gorm:"column:issue;type:varchar(20);not null;default:''" json:"issue" db:"issue"`
	GameName          string       `gorm:"column:game_name;type:varchar(90);not null;default:'';comment:具体玩法" json:"game_name" db:"game_name"` // 具体玩法
	BillID            string       `gorm:"column:bill_id;type:varchar(50);index:idx_bill_id,priority:1" json:"bill_id" db:"bill_id"`
	AdminID           int64        `gorm:"column:admin_id;type:int unsigned" json:"admin_id" db:"admin_id"`
	Adminname         string       `gorm:"column:adminname;type:varchar(16);index:adminname,priority:1" json:"adminname" db:"adminname"`
	IP                string       `gorm:"column:ip;type:varchar(100);comment:ip" json:"ip" db:"ip"`                                                                          // ip
	Status            int64        `gorm:"column:status;type:tinyint;default:1;comment:状态：0已删除,1:成功" json:"status" db:"status"`                                        // 状态：0已删除,1:成功
	Remark            string       `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark" db:"remark"`                                                            // 备注
	PayType           string       `gorm:"column:pay_type;type:varchar(255);not null;default:'';comment:支付类型" json:"pay_type" db:"pay_type"`                              // 支付类型
	CreatedAt         sql.NullTime `gorm:"column:created_at;type:bigint;index:created_at,priority:1;autoCreateTime:milli;comment:交易时间" json:"created_at" db:"created_at"` // 交易时间
	UpdatedAt         sql.NullTime `gorm:"column:updated_at;type:bigint;index:idx_updated_at,priority:1;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	Description       string       `gorm:"column:description;type:varchar(255);comment:转账说明" json:"description" db:"description"`                                                 // 转账说明
	IsFreeze          int64        `gorm:"column:is_freeze;type:tinyint;not null;default:0;comment:是否冻结 1:是 0:否" json:"is_freeze" db:"is_freeze"`                               // 是否冻结 1:是 0:否
	ThirdMerchantName string       `gorm:"column:third_merchant_name;type:varchar(200);not null;default:'';comment:三方商户名称" json:"third_merchant_name" db:"third_merchant_name"` // 三方商户名称
	MerchantNum       string       `gorm:"column:merchant_num;type:varchar(200);not null;default:'';comment:商户编号" json:"merchant_num" db:"merchant_num"`                          // 商户编号
	ThirdTrackNum     string       `gorm:"column:third_track_num;type:varchar(200);not null;default:'';comment:三方单号" json:"third_track_num" db:"third_track_num"`                 // 三方单号
	ThirdBillNo       string       `gorm:"column:third_bill_no;type:varchar(255);comment:第三方平台订单号" json:"third_bill_no" db:"third_bill_no"`                                   // 第三方平台订单号
}

type GiftMoneyTransaction struct {
	ID                int64        `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID            int64        `gorm:"column:user_id;type:int unsigned;not null;index:user_id,priority:1;default:0" json:"user_id" db:"user_id"`
	Username          string       `gorm:"column:username;type:varchar(255);not null;index:idx_uname,priority:1;default:''" json:"username" db:"username"`
	IsTester          int64        `gorm:"column:is_tester;type:tinyint(1)" json:"is_tester" db:"is_tester"`
	ForefatherIds     string       `gorm:"column:forefather_ids;type:varchar(1024)" json:"forefather_ids" db:"forefather_ids"`
	Amount            float64      `gorm:"column:amount;type:decimal(12,2);not null;index:idx_amount,priority:1;default:0.00;comment:用户当前交易金额" json:"amount" db:"amount"`            // 用户当前交易金额
	BankType          int64        `gorm:"column:bank_type;type:tinyint;comment:1:银行卡, 2:数字钱包, 3:易汇钱包, 4:支付宝, 5:微信, 6:数字人民币" json:"bank_type" db:"bank_type"`           // 1:银行卡, 2:数字钱包, 3:易汇钱包, 4:支付宝, 5:微信, 6:数字人民币
	TransFatherID     int64        `gorm:"column:trans_father_id;type:int unsigned;index:trans_father_id_index,priority:1;comment:账变类型父类" json:"trans_father_id" db:"trans_father_id"` // 账变类型父类
	TransTypesID      int64        `gorm:"column:trans_types_id;type:mediumint unsigned;not null;index:act_id,priority:1;default:0" json:"trans_types_id" db:"trans_types_id"`
	TransTypesCnTitle string       `gorm:"column:trans_types_cn_title;type:varchar(30)" json:"trans_types_cn_title" db:"trans_types_cn_title"`
	TransTypesEnTitle string       `gorm:"column:trans_types_en_title;type:varchar(50);not null;default:''" json:"trans_types_en_title" db:"trans_types_en_title"`
	IsIncome          int64        `gorm:"column:is_income;type:tinyint(1);comment:是否入款" json:"is_income" db:"is_income"` // 是否入款
	BeforeMoney       float64      `gorm:"column:before_money;type:decimal(12,2);not null;default:0.00" json:"before_money" db:"before_money"`
	Money             float64      `gorm:"column:money;type:decimal(12,2);not null;default:0.00;comment:交易后金额" json:"money" db:"money"`               // 交易后金额
	GameCode          string       `gorm:"column:game_code;type:varchar(30);index:lottery_id,priority:1;comment:游戏game" json:"game_code" db:"game_code"` // 游戏game
	Issue             string       `gorm:"column:issue;type:varchar(20);not null;default:''" json:"issue" db:"issue"`
	GameName          string       `gorm:"column:game_name;type:varchar(90);not null;default:'';comment:具体玩法" json:"game_name" db:"game_name"` // 具体玩法
	BillID            string       `gorm:"column:bill_id;type:varchar(50);index:idx_bill_id,priority:1" json:"bill_id" db:"bill_id"`
	AdminID           int64        `gorm:"column:admin_id;type:int unsigned" json:"admin_id" db:"admin_id"`
	Adminname         string       `gorm:"column:adminname;type:varchar(16);index:adminname,priority:1" json:"adminname" db:"adminname"`
	IP                string       `gorm:"column:ip;type:varchar(100);comment:ip" json:"ip" db:"ip"`                                                                              // ip
	Status            int64        `gorm:"column:status;type:tinyint;default:1;comment:状态：0未处理（默认）,1:审核中,2充值成功,3:失败,4:已撤销,5:已拒绝" json:"status" db:"status"` // 状态：0未处理（默认）,1:审核中,2充值成功,3:失败,4:已撤销,5:已拒绝
	Remark            string       `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark" db:"remark"`                                                                // 备注
	CreatedAt         sql.NullTime `gorm:"column:created_at;type:bigint;index:created_at,priority:1;autoCreateTime:milli;comment:交易时间" json:"created_at" db:"created_at"`     // 交易时间
	UpdatedAt         sql.NullTime `gorm:"column:updated_at;type:bigint;index:idx_updated_at,priority:1;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	Description       string       `gorm:"column:description;type:varchar(255);comment:转账说明" json:"description" db:"description"`                                                 // 转账说明
	PayType           string       `gorm:"column:pay_type;type:varchar(255);not null;default:'';comment:支付类型" json:"pay_type" db:"pay_type"`                                      // 支付类型
	ParentID          int64        `gorm:"column:parent_id;type:int;not null;index:transactions_idx_parent_id,priority:1;default:0;comment:上级id" json:"parent_id" db:"parent_id"`   // 上级id
	ThirdBillNo       string       `gorm:"column:third_bill_no;type:varchar(255);not null;default:'';comment:第三方平台订单号" json:"third_bill_no" db:"third_bill_no"`               // 第三方平台订单号
	ThirdMerchantName string       `gorm:"column:third_merchant_name;type:varchar(200);not null;default:'';comment:三方商户名称" json:"third_merchant_name" db:"third_merchant_name"` // 三方商户名称
	MerchantNum       string       `gorm:"column:merchant_num;type:varchar(200);not null;default:'';comment:商户编号" json:"merchant_num" db:"merchant_num"`                          // 商户编号
	ThirdTrackNum     string       `gorm:"column:third_track_num;type:varchar(200);not null;default:'';comment:三方单号" json:"third_track_num" db:"third_track_num"`                 // 三方单号
}

type BetAmount struct {
	ID          int64          `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID      int64          `gorm:"column:user_id;type:int;not null;index:user_id_index,priority:1;default:0;comment:用户id" json:"user_id" db:"user_id"`                // 用户id
	Username    string         `gorm:"column:username;type:varchar(100);not null;index:user_name_index,priority:1;default:'';comment:用户名" json:"username" db:"username"` // 用户名
	TotalAmount float64        `gorm:"column:total_amount;type:decimal(12,2);not null;default:0.00;comment:要求打码量" json:"total_amount" db:"total_amount"`               // 要求打码量
	Comment     sql.NullString `gorm:"column:comment;type:varchar(255);comment:备注" json:"comment" db:"comment"`                                                           // 备注
	IsOpen      int64          `gorm:"column:is_open;type:tinyint;not null;default:1;comment:是否开启0:否;1:是" json:"is_open" db:"is_open"`                                // 是否开启0:否;1:是
	DeleteAt    sql.NullTime   `gorm:"column:delete_at;type:timestamp;comment:重置时间 就是失效时间" json:"delete_at" db:"delete_at"`                                       // 重置时间 就是失效时间
	CreatedAt   sql.NullTime   `gorm:"column:created_at;type:bigint;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt   sql.NullTime   `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
}

type UserGameBetAmount struct {
	ID             int64        `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	DateAt         string       `gorm:"column:date_at;type:varchar(50);index:date_at,priority:1;comment:当天时间" json:"date_at" db:"date_at"`                                                                             // 当天时间
	SeriesID       int64        `gorm:"column:series_id;type:int;not null;default:0;comment:游戏系列id" json:"series_id" db:"series_id"`                                                                                   // 游戏系列id
	GameCode       string       `gorm:"column:game_code;type:varchar(50);not null;index:game_code_index,priority:1;default:'';comment:游戏类型" json:"game_code" db:"game_code"`                                           // 游戏类型
	UserID         int64        `gorm:"column:user_id;type:int;not null;index:user_game_bet_amounts_user_id_index,priority:1;default:0;comment:用户id" json:"user_id" db:"user_id"`                                        // 用户id
	Username       string       `gorm:"column:username;type:varchar(100);not null;default:'';comment:用户名" json:"username" db:"username"`                                                                                // 用户名
	ParentID       int64        `gorm:"column:parent_id;type:int;not null;default:0;comment:用户id" json:"parent_id" db:"parent_id"`                                                                                       // 用户id
	ParentName     string       `gorm:"column:parent_name;type:varchar(100);not null;default:'';comment:用户名" json:"parent_name" db:"parent_name"`                                                                       // 用户名
	ForefatherIds  string       `gorm:"column:forefather_ids;type:varchar(255);comment:上级树,多个以逗号隔开" json:"forefather_ids" db:"forefather_ids"`                                                                   // 上级树,多个以逗号隔开
	BetNums        int64        `gorm:"column:bet_nums;type:int;not null;default:0;comment:注单量" json:"bet_nums" db:"bet_nums"`                                                                                          // 注单量
	BetAmount      float64      `gorm:"column:bet_amount;type:decimal(10,2);default:0.00;comment:用户投注金额 包括无效金额" json:"bet_amount" db:"bet_amount"`                                                             // 用户投注金额 包括无效金额
	ValidBetAmount float64      `gorm:"column:valid_bet_amount;type:decimal(12,2);not null;default:0.00;comment:有效投注金额" json:"valid_bet_amount" db:"valid_bet_amount"`                                               // 有效投注金额
	NetAmount      float64      `gorm:"column:net_amount;type:decimal(12,2);not null;default:0.00;comment:玩家的所赢金额" json:"net_amount" db:"net_amount"`                                                               // 玩家的所赢金额
	TotalFanshui   float64      `gorm:"column:total_fanshui;type:decimal(12,2);comment:返水金额" json:"total_fanshui" db:"total_fanshui"`                                                                                  // 返水金额
	IsFanshui      int64        `gorm:"column:is_fanshui;type:tinyint;not null;index:is_fanshui_index,priority:1;default:0;comment:0:未返水;1:返水进行中;2:已返水;" json:"is_fanshui" db:"is_fanshui"`                     // 0:未返水;1:返水进行中;2:已返水;
	FanshuiAt      sql.NullTime `gorm:"column:fanshui_at;type:timestamp;index:user_game_bet_amounts_fanshui_at_index,priority:1;comment:返水时间" json:"fanshui_at" db:"fanshui_at"`                                       // 返水时间
	IsManual       int64        `gorm:"column:is_manual;type:tinyint;comment:是否手动返水0:否;1:是手动" json:"is_manual" db:"is_manual"`                                                                                   // 是否手动返水0:否;1:是手动
	IsOpen         int64        `gorm:"column:is_open;type:tinyint;not null;index:user_game_bet_amounts_is_open_index,priority:1;default:1;comment:是否启用0:否;1:是;如果为0就重新统计打码量" json:"is_open" db:"is_open"` // 是否启用0:否;1:是;如果为0就重新统计打码量
	IsYongjin      int64        `gorm:"column:is_yongjin;type:tinyint;index:user_game_bet_amounts_is_yongjin_index,priority:1;comment:是否统计了代理佣金0否，1进行中；2:已经返佣" json:"is_yongjin" db:"is_yongjin"`         // 是否统计了代理佣金0否，1进行中；2:已经返佣
	IsTotalBet     int64        `gorm:"column:is_total_bet;type:tinyint;index:user_game_bet_amounts_is_total_bet_index,priority:1;comment:是否统计到总注单表：0否，1进行中；2:已统计" json:"is_total_bet" db:"is_total_bet"`  // 是否统计到总注单表：0否，1进行中；2:已统计
	IsTodayProfit  int64        `gorm:"column:is_today_profit;type:tinyint;comment:是否统计到今日盈亏表：0否，1进行中；2:已统计；today_profits" json:"is_today_profit" db:"is_today_profit"`                                   // 是否统计到今日盈亏表：0否，1进行中；2:已统计；today_profits
	IsDeleted      int64        `gorm:"column:is_deleted;type:tinyint;comment:是否删除0:否;1:已删除" json:"is_deleted" db:"is_deleted"`                                                                                    // 是否删除0:否;1:已删除
	CreatedAt      sql.NullTime `gorm:"column:created_at;type:bigint;index:user_game_bet_amounts_created_at_index,priority:1;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt      sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
}

type ClearBetAmountLog struct {
	ID                   int64        `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID               int64        `gorm:"column:user_id;type:int;not null;index:idx_clear_bet_amount_logs_user_id,priority:1;default:0;comment:用户id" json:"user_id" db:"user_id"`               // 用户id
	Username             string       `gorm:"column:username;type:varchar(100);not null;default:'';comment:用户名" json:"username" db:"username"`                                                     // 用户名
	BetAmount            float64      `gorm:"column:bet_amount;type:decimal(12,2);not null;default:0.00;comment:清空总打码金额" json:"bet_amount" db:"bet_amount"`                                    // 清空总打码金额
	UserGameBetAmountIds string       `gorm:"column:user_game_bet_amount_ids;type:longtext;comment:清空打码记录的ids" json:"user_game_bet_amount_ids" db:"user_game_bet_amount_ids"`                  // 清空打码记录的ids
	Content              string       `gorm:"column:content;type:text;comment:打码详情内容，jsonArr格式" json:"content" db:"content"`                                                                  // 打码详情内容，jsonArr格式
	Status               int64        `gorm:"column:status;type:smallint;not null;index:idx_clear_bet_amount_logs_status,priority:1;default:0;comment:是否完成[0:否;1:是]" json:"status" db:"status"` // 是否完成[0:否;1:是]
	FinishedAt           sql.NullTime `gorm:"column:finished_at;type:datetime;comment:完成时间" json:"finished_at" db:"finished_at"`                                                                  // 完成时间
	CreatedAt            sql.NullTime `gorm:"column:created_at;type:bigint;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt            sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	AdminName            string       `gorm:"column:admin_name;type:varchar(255);not null;default:'';comment:操作员" json:"admin_name" db:"admin_name"` // 操作员
}
type BetAmountLog struct {
	ID            int64        `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID        int64        `gorm:"column:user_id;type:int;not null;index:user_id,priority:1;default:0;comment:用户id" json:"user_id" db:"user_id"`        // 用户id
	Username      string       `gorm:"column:username;type:varchar(100);not null;default:'';comment:用户名" json:"username" db:"username"`                    // 用户名
	TotalAmount   float64      `gorm:"column:total_amount;type:decimal(12,2);not null;default:0.00;comment:要求打码量" json:"total_amount" db:"total_amount"` // 要求打码量
	Remark        string       `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark" db:"remark"`                                                // 备注
	TransactionID int64        `gorm:"column:transaction_id;type:int" json:"transaction_id" db:"transaction_id"`
	State         int64        `gorm:"column:state;type:tinyint;not null;index:bet_amount_state_index,priority:1;default:1;comment:是否有效0:否;1:是" json:"state" db:"state"` // 是否有效0:否;1:是
	DeleteAt      sql.NullTime `gorm:"column:delete_at;type:timestamp;comment:重置时间 就是失效时间" json:"delete_at" db:"delete_at"`                                          // 重置时间 就是失效时间
	CreatedAt     sql.NullTime `gorm:"column:created_at;type:bigint;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt     sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
}

type LevelUpgradeLog struct {
	ID            int64        `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID        int64        `gorm:"column:user_id;type:int;comment:用户uid" json:"user_id" db:"user_id"`                                                               // 用户uid
	Username      string       `gorm:"column:username;type:varchar(50);comment:用户账号" json:"username" db:"username"`                                                   // 用户账号
	OldPayLevelID int64        `gorm:"column:old_pay_level_id;type:int unsigned;not null;default:0;comment:原始支付分层ID" json:"old_pay_level_id" db:"old_pay_level_id"` // 原始支付分层ID
	NewPayLevelID int64        `gorm:"column:new_pay_level_id;type:int unsigned;not null;default:0;comment:新支付分层ID" json:"new_pay_level_id" db:"new_pay_level_id"`   // 新支付分层ID
	AdminUser     string       `gorm:"column:admin_user;type:varchar(50);comment:操作人" json:"admin_user" db:"admin_user"`                                               // 操作人
	Remark        string       `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark" db:"remark"`                                                            // 备注
	CreatedAt     sql.NullTime `gorm:"column:created_at;type:bigint;autoCreateTime:milli" json:"created_at" db:"created_at"`
	Reason        string       `gorm:"column:reason;type:varchar(255);not null;default:'';comment:变动原因" json:"reason" db:"reason"` // 变动原因
	BillID        string       `gorm:"column:bill_id;type:varchar(50);comment:订单号" json:"bill_id" db:"bill_id"`                     // 订单号
}

type FissionSetting struct {
	ID         int64        `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	FissionKey string       `gorm:"column:fission_key;type:varchar(255);not null;index:index_key,priority:1;default:'';comment:配置名称下标" json:"fission_key" db:"fission_key"` // 配置名称下标
	FissionVal string       `gorm:"column:fission_val;type:text;comment:值" json:"fission_val" db:"fission_val"`                                                                  // 值
	IsOpen     int64        `gorm:"column:is_open;type:tinyint(1);not null;default:1;comment:是否启用" json:"is_open" db:"is_open"`                                               // 是否启用
	CreatedAt  sql.NullTime `gorm:"column:created_at;type:bigint;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt  sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	AdminUser  string       `gorm:"column:admin_user;type:varchar(255);not null;default:''" json:"admin_user" db:"admin_user"`
}

type FissionInviterReward struct {
	StartNumber string `json:"start_number"`
	EndNumber   string `json:"end_number"`
	WinMoney    string `json:"win_money"`
	MoreMoney   string `json:"more_money"`
	Ge          bool   `json:"ge"`
}

type FissionReward struct {
	ID                     int64        `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:主键" json:"id" db:"id"`                                                                          // 主键
	UserID                 int64        `gorm:"column:user_id;type:int unsigned;not null;default:0;comment:用戶id" json:"user_id" db:"user_id"`                                                                    // 用戶id
	Username               string       `gorm:"column:username;type:varchar(100);not null;default:'';comment:用户账号" json:"username" db:"username"`                                                              // 用户账号
	PayLevelsID            int64        `gorm:"column:pay_levels_id;type:int unsigned;not null;default:0;comment:支付层级id" json:"pay_levels_id" db:"pay_levels_id"`                                              // 支付层级id
	PayLevelsName          string       `gorm:"column:pay_levels_name;type:varchar(100);not null;default:'';comment:支付层级name" json:"pay_levels_name" db:"pay_levels_name"`                                     // 支付层级name
	FromUserID             int64        `gorm:"column:from_user_id;type:int unsigned;not null;default:0;comment:来自用戶id" json:"from_user_id" db:"from_user_id"`                                                 // 来自用戶id
	FromUsername           string       `gorm:"column:from_username;type:varchar(100);not null;default:'';comment:来自用户账号" json:"from_username" db:"from_username"`                                           // 来自用户账号
	RewardType             int64        `gorm:"column:reward_type;type:tinyint unsigned;comment:0好友绑定账户、1好友首充、2好友打码返利" json:"reward_type" db:"reward_type"`                                        // 0好友绑定账户、1好友首充、2好友打码返利
	Money                  float64      `gorm:"column:money;type:decimal(10,2);default:0.00;comment:金额" json:"money" db:"money"`                                                                                 // 金额
	RewardStatus           int64        `gorm:"column:reward_status;type:tinyint unsigned;index:index_fission_reward_status,priority:1;comment:0待派发,1已派发,2拒绝派发" json:"reward_status" db:"reward_status"` // 0待派发,1已派发,2拒绝派发
	CreatedAt              sql.NullTime `gorm:"column:created_at;type:bigint;index:index_fission_created_at,priority:1;autoCreateTime:milli;comment:生成时间" json:"created_at" db:"created_at"`                   // 生成时间
	UpdatedAt              sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli;comment:更新时间" json:"updated_at" db:"updated_at"`                                                             // 更新时间
	AuditRemark            string       `gorm:"column:audit_remark;type:varchar(255);comment:审核备注" json:"audit_remark" db:"audit_remark"`                                                                      // 审核备注
	AuditTime              sql.NullTime `gorm:"column:audit_time;type:timestamp;comment:审核时间" json:"audit_time" db:"audit_time"`                                                                               // 审核时间
	AuditUsername          string       `gorm:"column:audit_username;type:varchar(50);comment:审核人" json:"audit_username" db:"audit_username"`                                                                   // 审核人
	BackflowExt            string       `gorm:"column:backflow_ext;type:json;comment:返水扩展字段" json:"backflow_ext" db:"backflow_ext"`                                                                          // 返水扩展字段
	UsernameParentName     string       `gorm:"column:username_parent_name;type:varchar(50);not null;default:'';comment:邀请人代理商" json:"username_parent_name" db:"username_parent_name"`                       // 邀请人代理商
	FromUsernameParentName string       `gorm:"column:from_username_parent_name;type:varchar(50);not null;default:'';comment:好友代理商" json:"from_username_parent_name" db:"from_username_parent_name"`          // 好友代理商
}

type FissionExclusiveReward struct {
	ID           int64        `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	ExtID        string       `gorm:"column:ext_id;type:varchar(255);comment:vip等级id，支付分层id" json:"ext_id" db:"ext_id"`                                                           // vip等级id，支付分层id
	Name         string       `gorm:"column:name;type:varchar(40);comment:奖励名称" json:"name" db:"name"`                                                                              // 奖励名称
	RewardType   int64        `gorm:"column:reward_type;type:tinyint unsigned;comment:奖励类型:0好友负盈利比例, 1好友充值比例,2线下陪玩,3首存彩金" json:"reward_type" db:"reward_type"` // 奖励类型:0好友负盈利比例, 1好友充值比例,2线下陪玩,3首存彩金
	RewardValue  string       `gorm:"column:reward_value;type:varchar(255);comment:奖励值" json:"reward_value" db:"reward_value"`                                                       // 奖励值
	Remark       string       `gorm:"column:remark;type:varchar(600);comment:描述" json:"remark" db:"remark"`                                                                           // 描述
	Type         int64        `gorm:"column:type;type:tinyint unsigned;default:1;comment:1按照vip等级设置奖励,2按支付分层设置奖励" json:"type" db:"type"`                               // 1按照vip等级设置奖励,2按支付分层设置奖励
	CreatedAt    sql.NullTime `gorm:"column:created_at;type:bigint;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt    sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	OpAdminUsers string       `gorm:"column:op_admin_users;type:varchar(255);comment:操作人" json:"op_admin_users" db:"op_admin_users"` // 操作人
}

type Activity struct {
	ID                    int64        `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id" db:"id"`
	Title                 string       `gorm:"column:title;type:varchar(50);comment:活动名称" json:"title" db:"title"`                                                                                      // 活动名称
	CaptainInvite         float64      `gorm:"column:captain_invite;type:decimal(12,2);default:0.00;comment:队伍每邀请1个有效会员奖励" json:"captain_invite" db:"captain_invite"`                           // 队伍每邀请1个有效会员奖励
	CaptainValidBetAmount float64      `gorm:"column:captain_valid_bet_amount;type:decimal(12,2);default:0.00;comment:队伍有效会员7日存款额" json:"captain_valid_bet_amount" db:"captain_valid_bet_amount"` // 队伍有效会员7日存款额
	CaptainHighestReward  float64      `gorm:"column:captain_highest_reward;type:decimal(12,2);default:0.00;comment:队长最高奖励" json:"captain_highest_reward" db:"captain_highest_reward"`                // 队长最高奖励
	CaptainCycle          int64        `gorm:"column:captain_cycle;type:tinyint;comment:队长奖励周期" json:"captain_cycle" db:"captain_cycle"`                                                              // 队长奖励周期
	TeamCycle             int64        `gorm:"column:team_cycle;type:tinyint(1);comment:团队奖励周期" json:"team_cycle" db:"team_cycle"`                                                                    // 团队奖励周期
	LuckyDrawDeposit      float64      `gorm:"column:lucky_draw_deposit;type:decimal(12,2);default:0.00;comment:抽奖资格存款金额" json:"lucky_draw_deposit" db:"lucky_draw_deposit"`                        // 抽奖资格存款金额
	Status                int64        `gorm:"column:status;type:tinyint;not null;default:0;comment:活动开关 0关 1开" json:"status" db:"status"`                                                            // 活动开关 0关 1开
	StartTime             sql.NullTime `gorm:"column:start_time;type:timestamp;comment:开始时间" json:"start_time" db:"start_time"`                                                                         // 开始时间
	EndTime               sql.NullTime `gorm:"column:end_time;type:timestamp;comment:结束时间" json:"end_time" db:"end_time"`                                                                               // 结束时间
	CreatedAt             sql.NullTime `gorm:"column:created_at;type:bigint;autoCreateTime:milli;comment:创建时间" json:"created_at" db:"created_at"`                                                       // 创建时间
	UpdatedAt             sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli;comment:更新时间" json:"updated_at" db:"updated_at"`                                                       // 更新时间
	IsPop                 int64        `gorm:"column:is_pop;type:tinyint;not null;default:0;comment:是否弹窗 0:否 1:是" json:"is_pop" db:"is_pop"`                                                          // 是否弹窗 0:否 1:是
	PopPcImgPath          string       `gorm:"column:pop_pc_img_path;type:varchar(255);not null;default:'';comment:pc弹窗图片" json:"pop_pc_img_path" db:"pop_pc_img_path"`                                 // pc弹窗图片
	PopMImgPath           string       `gorm:"column:pop_m_img_path;type:varchar(255);not null;default:'';comment:mobile弹窗图片" json:"pop_m_img_path" db:"pop_m_img_path"`                                // mobile弹窗图片
	IsFloat               int64        `gorm:"column:is_float;type:tinyint;not null;default:0;comment:是否悬浮" json:"is_float" db:"is_float"`                                                              // 是否悬浮
	FloatPcImgPath        string       `gorm:"column:float_pc_img_path;type:varchar(255);not null;default:'';comment:pc悬浮图片" json:"float_pc_img_path" db:"float_pc_img_path"`                           // pc悬浮图片
	FloatMImgPath         string       `gorm:"column:float_m_img_path;type:varchar(255);not null;default:'';comment:mobile悬浮图片" json:"float_m_img_path" db:"float_m_img_path"`                          // mobile悬浮图片
	GuessRule             string       `gorm:"column:guess_rule;type:text;comment:团队竞猜规则" json:"guess_rule" db:"guess_rule"`                                                                          // 团队竞猜规则
	BetRule               string       `gorm:"column:bet_rule;type:text;comment:投注规则" json:"bet_rule" db:"bet_rule"`                                                                                    // 投注规则
	LastBetAt             sql.NullTime `gorm:"column:last_bet_at;type:timestamp;comment:冠军赛最后投注时间" json:"last_bet_at" db:"last_bet_at"`                                                            // 冠军赛最后投注时间
}

type TeamUser struct {
	ID              int64        `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID          int64        `gorm:"column:user_id;type:int unsigned;index:idx_user_id_created_at,priority:1;comment:用户id" json:"user_id" db:"user_id"`                                                       // 用户id
	Username        string       `gorm:"column:username;type:varchar(255);not null;uniqueIndex:users_name_unique,priority:1;index:idx_username,priority:1;default:'';comment:用户名" json:"username" db:"username"` // 用户名
	ForefatherIds   string       `gorm:"column:forefather_ids;type:varchar(150);not null;default:'';comment:父亲树" json:"forefather_ids" db:"forefather_ids"`                                                      // 父亲树
	ParentName      string       `gorm:"column:parent_name;type:varchar(16);not null;default:'';comment:父级用户名" json:"parent_name" db:"parent_name"`                                                            // 父级用户名
	LeaderID        int64        `gorm:"column:leader_id;type:int unsigned;index:idx_leader_id_is_leader,priority:1;comment:队长id" json:"leader_id" db:"leader_id"`                                                // 队长id
	LeaderName      string       `gorm:"column:leader_name;type:varchar(255);not null;default:'';comment:队长名称" json:"leader_name" db:"leader_name"`                                                             // 队长名称
	RegisterAt      sql.NullTime `gorm:"column:register_at;type:timestamp;comment:注册时间" json:"register_at" db:"register_at"`                                                                                    // 注册时间
	TeamInviteCode  string       `gorm:"column:team_invite_code;type:varchar(255);not null;default:'';comment:队伍邀请码" json:"team_invite_code" db:"team_invite_code"`                                            // 队伍邀请码
	Remark          string       `gorm:"column:remark;type:varchar(255);not null;default:'';comment:备注" json:"remark" db:"remark"`                                                                                // 备注
	IsCalReward     int64        `gorm:"column:is_cal_reward;type:tinyint;not null;default:0" json:"is_cal_reward" db:"is_cal_reward"`
	IsLeader        int64        `gorm:"column:is_leader;type:tinyint unsigned;index:idx_leader_id_is_leader,priority:2;comment:0:不是, 1:是" json:"is_leader" db:"is_leader"`   // 0:不是, 1:是
	TotalDeposit    float64      `gorm:"column:total_deposit;type:decimal(12,2);not null;default:0.00;comment:累计存款" json:"total_deposit" db:"total_deposit"`                 // 累计存款
	TotalTeamReward float64      `gorm:"column:total_team_reward;type:decimal(12,2);not null;default:0.00;comment:累计团队奖励" json:"total_team_reward" db:"total_team_reward"` // 累计团队奖励
	WinBetNum       int64        `gorm:"column:win_bet_num;type:int unsigned;comment:竞猜赢的总数" json:"win_bet_num" db:"win_bet_num"`                                          // 竞猜赢的总数
	TotalBetNum     int64        `gorm:"column:total_bet_num;type:int unsigned;comment:竞猜总数" json:"total_bet_num" db:"total_bet_num"`                                        // 竞猜总数
	CreatedAt       sql.NullTime `gorm:"column:created_at;type:bigint;index:idx_user_id_created_at,priority:2;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt       sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
}

type ActTransaction struct {
	ID                int64        `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID            int64        `gorm:"column:user_id;type:int unsigned;not null;default:0" json:"user_id" db:"user_id"`
	Username          string       `gorm:"column:username;type:varchar(255);not null;index:idx_uname,priority:1;default:''" json:"username" db:"username"`
	IsTester          int64        `gorm:"column:is_tester;type:tinyint(1)" json:"is_tester" db:"is_tester"`
	ParentID          int64        `gorm:"column:parent_id;type:int;not null;default:0;comment:上级id" json:"parent_id" db:"parent_id"` // 上级id
	ForefatherIds     string       `gorm:"column:forefather_ids;type:varchar(1024)" json:"forefather_ids" db:"forefather_ids"`
	Amount            float64      `gorm:"column:amount;type:decimal(12,2);not null;index:uaaa,priority:1;default:0.00;comment:用户当前交易金额" json:"amount" db:"amount"` // 用户当前交易金额
	TransFatherID     int64        `gorm:"column:trans_father_id;type:int unsigned;comment:账变类型父类" json:"trans_father_id" db:"trans_father_id"`                       // 账变类型父类
	TransTypesID      int64        `gorm:"column:trans_types_id;type:mediumint unsigned;not null;index:act_id,priority:1;default:0" json:"trans_types_id" db:"trans_types_id"`
	TransTypesCnTitle string       `gorm:"column:trans_types_cn_title;type:varchar(30)" json:"trans_types_cn_title" db:"trans_types_cn_title"`
	TransTypesEnTitle string       `gorm:"column:trans_types_en_title;type:varchar(50);not null;default:''" json:"trans_types_en_title" db:"trans_types_en_title"`
	IsIncome          int64        `gorm:"column:is_income;type:tinyint(1);comment:是否入款" json:"is_income" db:"is_income"` // 是否入款
	BeforeMoney       float64      `gorm:"column:before_money;type:decimal(12,2);not null;default:0.00" json:"before_money" db:"before_money"`
	Money             float64      `gorm:"column:money;type:decimal(12,2);not null;default:0.00;comment:交易后金额" json:"money" db:"money"` // 交易后金额
	GameCode          string       `gorm:"column:game_code;type:varchar(30);comment:游戏game" json:"game_code" db:"game_code"`               // 游戏game
	Issue             string       `gorm:"column:issue;type:varchar(20);not null;default:''" json:"issue" db:"issue"`
	GameName          string       `gorm:"column:game_name;type:varchar(90);not null;default:'';comment:具体玩法" json:"game_name" db:"game_name"` // 具体玩法
	BillID            string       `gorm:"column:bill_id;type:varchar(50);index:bill_type,priority:1" json:"bill_id" db:"bill_id"`
	AdminID           int64        `gorm:"column:admin_id;type:int unsigned" json:"admin_id" db:"admin_id"`
	Adminname         string       `gorm:"column:adminname;type:varchar(16);index:adminname,priority:1" json:"adminname" db:"adminname"`
	IP                string       `gorm:"column:ip;type:varchar(18)" json:"ip" db:"ip"`
	Status            int64        `gorm:"column:status;type:tinyint;default:1;comment:状态：0已删除,1:成功" json:"status" db:"status"`                                        // 状态：0已删除,1:成功
	Remark            string       `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark" db:"remark"`                                                            // 备注
	PayType           string       `gorm:"column:pay_type;type:varchar(255);not null;default:'';comment:支付类型" json:"pay_type" db:"pay_type"`                              // 支付类型
	CreatedAt         sql.NullTime `gorm:"column:created_at;type:bigint;index:created_at,priority:1;autoCreateTime:milli;comment:交易时间" json:"created_at" db:"created_at"` // 交易时间
	UpdatedAt         sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
	Description       string       `gorm:"column:description;type:varchar(255);comment:转账说明" json:"description" db:"description"`                                                                          // 转账说明
	IsCal             int64        `gorm:"column:is_cal;type:tinyint;not null;index:is_cal,priority:1;default:0;comment:is_cal 0:未结算 1:结算中 2：已结算" json:"is_cal" db:"is_cal"`                          // is_cal 0:未结算 1:结算中 2：已结算
	LeaderID          int64        `gorm:"column:leader_id;type:int;not null;index:leader_id,priority:1;default:0;comment:队长id" json:"leader_id" db:"leader_id"`                                             // 队长id
	BankType          int64        `gorm:"column:bank_type;type:tinyint unsigned;not null;default:0;comment:1:银行卡, 2:数字钱包, 3:易汇钱包, 4:支付宝, 5:微信, 6:数字人民币" json:"bank_type" db:"bank_type"` // 1:银行卡, 2:数字钱包, 3:易汇钱包, 4:支付宝, 5:微信, 6:数字人民币
	ThirdMerchantName string       `gorm:"column:third_merchant_name;type:varchar(255);not null;default:'';comment:三方商户名称" json:"third_merchant_name" db:"third_merchant_name"`                          // 三方商户名称
	MerchantNum       string       `gorm:"column:merchant_num;type:varchar(255);not null;default:'';comment:商户编号" json:"merchant_num" db:"merchant_num"`                                                   // 商户编号
	ThirdTrackNum     string       `gorm:"column:third_track_num;type:varchar(255);not null;default:'';comment:三方单号" json:"third_track_num" db:"third_track_num"`                                          // 三方单号
}

type ActivityReport struct {
	ID               int64        `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id" db:"id"`
	UserID           int64        `gorm:"column:user_id;type:int unsigned;not null;uniqueIndex:activity_reports_user_id_unique,priority:1;default:0;comment:队长id" json:"user_id" db:"user_id"` // 队长id
	TotalWinTimes    int64        `gorm:"column:total_win_times;type:int unsigned;not null;default:0;comment:猜赢次数" json:"total_win_times" db:"total_win_times"`                              // 猜赢次数
	TotalCoefficient float64      `gorm:"column:total_coefficient;type:decimal(6,5);not null;default:0.00000;comment:竞猜系数" json:"total_coefficient" db:"total_coefficient"`                  // 竞猜系数
	TotalDeposit     float64      `gorm:"column:total_deposit;type:decimal(12,2);not null;default:0.00;comment:团队存款" json:"total_deposit" db:"total_deposit"`                                // 团队存款
	TotalReward      float64      `gorm:"column:total_reward;type:decimal(12,2);not null;default:0.00;comment:团队奖励" json:"total_reward" db:"total_reward"`                                   // 团队奖励
	MatchEventSum    string       `gorm:"column:match_event_sum;type:text;comment:分组竞猜统计" json:"match_event_sum" db:"match_event_sum"`                                                     // 分组竞猜统计
	CreatedAt        sql.NullTime `gorm:"column:created_at;type:bigint;autoCreateTime:milli" json:"created_at" db:"created_at"`
	UpdatedAt        sql.NullTime `gorm:"column:updated_at;type:bigint;autoUpdateTime:milli" json:"updated_at" db:"updated_at"`
}
