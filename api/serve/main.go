package main

import (
	rice "github.com/GeertJohan/go.rice"
	_ "github.com/go-sql-driver/mysql"
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/api"
	"github.com/zhanghup/go-framework/api/deploy"
	cfg2 "github.com/zhanghup/go-framework/ctx/cfg"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
)

func main() {
	deploy.InitConfigFolder()

	TestContext := &cfg2.Cfg{}
	cfg := rice.MustFindBox("conf")
	app.StartCommon(TestContext, cfg)

	g := gin.Default()
	e, _ := xorm.NewCfgEngine(TestContext)
	app.Sync(e)

	Public := g.Group(TestContext.Gin.Prefix + "/")
	//Auth := g.Group("/")

	api.Login(Public, e)
	api.Resource(Public, Public, e)
	api.WcCallback(Public, e)
	api.Wc(Public,e)

	g.CfgRun()
}
