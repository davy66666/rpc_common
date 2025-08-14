package values

// 不需要加前缀的
const (
	RedisKeyNodeKeyPrefix           = "snowflake:node:%d"         // 节点 ID 的 Redis 键前缀
	RedisKeyLockUniqueRequest       = "game:unique:request:%v"    // 请求幂等锁
	RedisKeyLockUniqueRoundID       = "game:unique:round:%v"      // 局数唯一ID
	RedisKeyUniqueUserId            = "user:unique:id"            // 用户唯一ID
	RedisKeyUniqueMsgId             = "msg:unique:id:%v"          // 消息唯一ID
	RedisKeyUniqueSysId             = "{sys:unique:id}"           // 系统唯一ID
	RedisKeyUniqueAd                = "sys:unique:ad:%s:%s:%s:%s" // 广告唯一事件
	SELFMSGSENDTIMECACHEKEY         = "self_msg_send_time_cache_key_%s"
	SendUserMessageJobCacheComplete = "send_user_message_job_cache_complete_%s"
	SendUserMessageJobCacheTotal    = "send_user_message_job_cache_%s"
)
