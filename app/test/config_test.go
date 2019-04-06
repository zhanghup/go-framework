package main

import (
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/app"
	"github.com/zhanghup/go-framework/tools"
)

type testContext struct {
	*app.Context
	//System struct {
	//HttpPort string `json:"http-port"`
	//} `json:"system"`
}

func (this *testContext) GetContext() *app.Context {
	return this.Context
}

var TestContext = new(testContext)

func TestIniConfig(t *testing.T) {
	cfg := rice.MustFindBox("../../conf")
	app.InitApp(TestContext, cfg)
	tools.PrintStruct(TestContext)
}
