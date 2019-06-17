package main

import (
	"github.com/zhanghup/go-framework/ctx/cfg"
	"testing"
)

type testContext struct {
	*cfg.Cfg
	//System struct {
	//HttpPort string `json:"http-port"`
	//} `json:"system"`
}

func (this *testContext) GetCfg() *cfg.Cfg {
	return this.Cfg
}

var TestContext = new(testContext)

func TestIniConfig(t *testing.T) {

}
