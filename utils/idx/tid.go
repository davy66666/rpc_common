package idx

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cast"
)

func GenTid() int64 {
	// 获取当前时间
	now := time.Now()

	// 格式化时间到秒级别
	timeStr := now.Format("060102150405")
	// 使用 Unix 时间戳（毫秒级）创建一个新的随机源
	source := rand.NewSource(time.Now().UnixNano())

	// 创建一个新的随机数生成器
	r := rand.New(source)

	// 使用生成器生成随机数
	randomInt := 100 + r.Intn(899) // 生成 100 到 999 的随机整数
	n := fmt.Sprintf("%s%d", timeStr, randomInt)
	return cast.ToInt64(n)
}

func GenSecondId() int64 {
	// 获取当前时间
	now := time.Now()
	// 格式化时间到秒级别
	timeStr := now.Format("060102150405")
	return cast.ToInt64(timeStr)
}
