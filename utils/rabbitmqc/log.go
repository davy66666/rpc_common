package rabbitmqc

import (
	"github.com/zeromicro/go-zero/core/logx"
)

type mqLogger struct {
}

func (l *mqLogger) Info(msg string, fields map[string]interface{}) {
	//if msg == "" && len(fields) == 0 {
	//	return
	//}
	//var ls []logx.LogField
	//for k, v := range fields {
	//	ls = append(ls, logx.Field(k, v))
	//}
	//logx.Infow(msg, ls...)
}

func (l *mqLogger) Warning(msg string, fields map[string]interface{}) {
	return
}

func (l *mqLogger) Error(msg string, fields map[string]interface{}) {
	if msg == "" && len(fields) == 0 {
		return
	}
	var ls []logx.LogField
	for k, v := range fields {
		ls = append(ls, logx.Field(k, v))
	}
	logx.Errorw(msg, ls...)
}

func (l *mqLogger) Fatal(msg string, fields map[string]interface{}) {
	return
}

func (l *mqLogger) Debug(msg string, fields map[string]interface{}) {
	//if msg == "" && len(fields) == 0 {
	//	return
	//}
	//var ls []logx.LogField
	//for k, v := range fields {
	//	ls = append(ls, logx.Field(k, v))
	//}
	//logx.Debugw(msg, ls...)
}

func (l *mqLogger) Level(level string) {
	return
}

func (l *mqLogger) OutputPath(path string) (err error) {
	return nil
}
