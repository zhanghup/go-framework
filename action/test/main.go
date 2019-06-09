package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/ctx"
	"github.com/zhanghup/go-framework/pkg/gin"
)

var t *gin.RouterGroup

func main() {
	cfg := rice.MustFindBox("../../conf")
	ctx2 := new(ctx.Cfg)
	ctx.InitCfg(ctx2, cfg)
	//app.LogError("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", 1)
	g := gin.Default(ctx.GetCfg())
	//t = action.RegisterGroup("/test")
	g.CfgRun()

}

//func init() {
//	action.RegisterRouters(func() {
//		t.GET("/aaa", func(c *gin.Context) {
//			c.String(http.StatusOK, "hello world")
//		})
//	})
//}
