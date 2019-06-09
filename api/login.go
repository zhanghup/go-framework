package api

import (
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
)

func Login(g *gin.RouterGroup, e *xorm.Engine) {

	g.POST("/register", func(c *gin.Context) {

	})
}

