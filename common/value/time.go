package values

import "time"

const (
	TimestampOneDay               = 86400     // 一天时间秒
	TimestampOneWeek              = 604800    // 一周时间秒
	TimestampThreeDayMilliseconds = 259200000 // 3天的时间毫秒
	TimestampMilliOneDay          = 86400000  // 一天时间毫秒
)

var (
	OneDayTime   = time.Hour * 24
	TwoDayTime   = time.Hour * 24 * 2
	TenDayTime   = time.Hour * 24 * 10
	OneWeekTime  = time.Hour * 24 * 7
	TwoWeekTime  = time.Hour * 24 * 14
	OneMonthTime = time.Hour * 24 * 31
	TwoMonthTime = time.Hour * 24 * 62
)
