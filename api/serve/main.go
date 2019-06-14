package main

import (
	rice "github.com/GeertJohan/go.rice"
	_ "github.com/go-sql-driver/mysql"
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/api"
	"github.com/zhanghup/go-framework/ctx"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
)

func main() {
	TestContext := &ctx.Cfg{}
	cfg := rice.MustFindBox("../../conf")
	app.StartCommon(TestContext, cfg)

	g := gin.Default()
	e, _ := xorm.NewCfgEngine(TestContext)
	app.Sync(e)

	Public := g.Group("/")
	//Auth := g.Group("/")

	api.Login(Public, e)

	g.CfgRun()
}
