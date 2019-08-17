package main

import (
	"github.com/zhanghup/go-framework/ctx/cfg"
	"testing"

	rice "github.com/GeertJohan/go.rice"
)

type testLogContext struct {
	*cfg.Cfg
	System struct {
		HttpPort string `json:"http-port"`
	} `json:"system"`
}

func (this *testLogContext) GetCfg() *cfg.Cfg {
	return this.Cfg
}
func TestLogError(t *testing.T) {
	TestContext := new(testLogContext)
	cfg := rice.MustFindBox("../../conf")
	cfg.InitCfg(TestContext, cfg)
	for i := 0; i < 100000; i++ {
		//go func(i int) {
		cfg.LogError("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", i)
		cfg.LogInfo("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", i)
		//}(i)
	}

	cfg.LogError("111")
	//for {
	//time.Sleep(time.Second)
	//}
}
