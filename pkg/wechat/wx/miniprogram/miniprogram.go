package miniprogram

import (
	"github.com/zhanghup/go-framework/pkg/wechat/wx/context"
)

// MiniProgram struct extends ctx
type MiniProgram struct {
	*context.Context
}

// NewMiniProgram 实例化小程序接口
func NewMiniProgram(context *context.Context) *MiniProgram {
	miniProgram := new(MiniProgram)
	miniProgram.Context = context
	return miniProgram
}
