package main

import (
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/ctx"
)

type testLogContext struct {
	*ctx.Cfg
	System struct {
		HttpPort string `json:"http-port"`
	} `json:"system"`
}

func (this *testLogContext) GetCfg() *ctx.Cfg {
	return this.Cfg
}
func TestLogError(t *testing.T) {
	TestContext := new(testLogContext)
	cfg := rice.MustFindBox("../../conf")
	ctx.InitCfg(TestContext, cfg)
	for i := 0; i < 100000; i++ {
		//go func(i int) {
		ctx.LogError("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", i)
		ctx.LogInfo("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", i)
		//}(i)
	}

	ctx.LogError("111")
	//for {
	//time.Sleep(time.Second)
	//}
}
