package ctx

import (
	"github.com/zhanghup/go-framework/ctx/cfg"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/wechat/wx"
	"github.com/zhanghup/go-framework/pkg/xorm"
)

type ctx struct {
	Cfg  *cfg.Cfg
	Wc   *wx.Wechat
	Gin  *gin.Engine
	Xorm *xorm.Engine
}

func NewCtx() *ctx {
	return &ctx{}
}
