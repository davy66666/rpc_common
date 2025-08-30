package errorx

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	errorx "github.com/zeromicro/x/errors"
)

func NewErr(code int, messages ...string) *errorx.CodeMsg {
	var msg string
	if len(messages) == 0 {
		msg = MapErrMsg(code)
	} else {
		msg = messages[0]
	}

	return &errorx.CodeMsg{
		Code: code,
		Msg:  msg,
	}
}

// ExtractStackTrace 获取原始堆栈信息并返回结构化的 JSON 格式
func ExtractStackTrace(err error) []map[string]string {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	var stackErr stackTracer
	if errors.As(err, &stackErr) {
		var filteredStack []map[string]string
		for _, f := range stackErr.StackTrace() {
			frame := fmt.Sprintf("%+v", f)
			parts := strings.Split(frame, "\n\t") // 分割函数名和调用位置
			if len(parts) == 2 {
				funcName := strings.TrimSpace(parts[0])
				caller := strings.TrimSpace(parts[1])

				// 过滤掉不需要的堆栈
				if !strings.Contains(caller, "github") &&
					!strings.Contains(caller, "net/http") &&
					!strings.Contains(caller, "runtime") &&
					!strings.Contains(caller, "middleware") &&
					!strings.Contains(caller, "handler") {
					filteredStack = append(filteredStack, map[string]string{
						"func":   funcName,
						"caller": caller,
					})
				}
			}
		}
		return filteredStack
	}
	return nil
}
