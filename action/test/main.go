package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/action"
	"github.com/zhanghup/go-framework/app"
)

func main() {
	cfg := rice.MustFindBox("../../conf")
	ctx := new(app.Context)
	app.InitApp(ctx, cfg)
	action.InitGin()

}
