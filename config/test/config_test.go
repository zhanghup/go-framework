package main

import (
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/config"
	"github.com/zhanghup/go-framework/tools"
)

type testContext struct {
	*config.Context
	//System struct {
	//HttpPort string `json:"http-port"`
	//} `json:"system"`
}

func (this *testContext) GetContext() *config.Context {
	return this.Context
}

var TestContext = new(testContext)

func TestIniConfig(t *testing.T) {
	cfg := rice.MustFindBox("../../conf")
	config.InitApp(TestContext, cfg)
	tools.PrintStruct(TestContext)
}
