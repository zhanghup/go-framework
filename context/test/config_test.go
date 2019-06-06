package main

import (
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/context"
	"github.com/zhanghup/go-framework/tools"
)

type testContext struct {
	*context.Context
	//System struct {
	//HttpPort string `json:"http-port"`
	//} `json:"system"`
}

func (this *testContext) GetContext() *context.Context {
	return this.Context
}

var TestContext = new(testContext)

func TestIniConfig(t *testing.T) {
	cfg := rice.MustFindBox("../../conf")
	context.InitApp(TestContext, cfg)
	tools.PrintStruct(TestContext)
}
