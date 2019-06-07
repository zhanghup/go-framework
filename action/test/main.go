package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/action"
	"github.com/zhanghup/go-framework/context"
	"github.com/zhanghup/go-framework/pkg/gin"
	"net/http"
)

var t *gin.RouterGroup

func main() {
	cfg := rice.MustFindBox("../../conf")
	ctx := new(context.Context)
	context.InitApp(ctx, cfg)
	//app.LogError("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", 1)
	action.InitGin()
	t = action.RegisterGroup("/test")

	action.Run()
}

func init() {
	action.RegisterRouters(func() {
		t.GET("/aaa", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world")
		})
	})
}
