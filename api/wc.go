package api

import (
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
)

func Wc(g *gin.RouterGroup, e *xorm.Engine) {

	g.Action("/wc/h5auth", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
		return nil, nil
	})

}
