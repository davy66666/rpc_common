package loggerx

import (
	"context"
	"runtime"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
}

func NewGormLogger() *GormLogger {
	return &GormLogger{}
}
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	callerSkip := getCallerSkip()
	var fields []logx.LogField
	fields = append(fields, logx.Field("flag", "sql"))
	if len(data) > 0 {
		fields = append(fields, logx.Field("data", data))
	}
	logx.WithContext(ctx).WithCallerSkip(callerSkip).Infow(msg, fields...)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	callerSkip := getCallerSkip()
	var fields []logx.LogField
	fields = append(fields, logx.Field("flag", "sql"))
	if len(data) > 0 {
		fields = append(fields, logx.Field("data", data))
	}
	logx.WithContext(ctx).WithCallerSkip(callerSkip).Errorw(msg, fields...)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, _ := fc()
	elapsed := time.Since(begin)
	callerSkip := getCallerSkip()
	var fields []logx.LogField
	fields = append(fields, logx.Field("flag", "sql"))
	if elapsed > 300000000 { // 300毫秒
		logx.WithContext(ctx).WithCallerSkip(callerSkip).Sloww(sql, fields...)
	} else {
		logx.WithContext(ctx).WithCallerSkip(callerSkip).Debugw(sql, fields...)
	}
}

func getCallerSkip() int {
	index := 5
	_, file, _, _ := runtime.Caller(index)
	if strings.Contains(file, "runtime") {
		return index + 1
	}
	if strings.Contains(file, "query") || strings.Contains(file, "gen") {
		return index + 1
	}
	return index
}
