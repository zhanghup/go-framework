package app

import (
	"github.com/zhanghup/go-framework/ctx"
)

// Info [INFO] 级日志输出，不带错误栈
func LogInfo(format string, args ...interface{}) {
	ctx.LogInfo(format, args...)
}

// Error [ERROR] 级日志输出，附带错误栈
func LogError(format string, args ...interface{}) {
	ctx.LogError(format, args...)
}
