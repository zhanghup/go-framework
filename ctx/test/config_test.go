package main

import (
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/ctx"
	"github.com/zhanghup/go-framework/tools"
)

type testContext struct {
	*ctx.Cfg
	//System struct {
	//HttpPort string `json:"http-port"`
	//} `json:"system"`
}

func (this *testContext) GetCfg() *ctx.Cfg {
	return this.Cfg
}

var TestContext = new(testContext)

func TestIniConfig(t *testing.T) {

}
