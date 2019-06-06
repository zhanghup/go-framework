package main

import (
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/config"
)

type testLogContext struct {
	*config.Context
	System struct {
		HttpPort string `json:"http-port"`
	} `json:"system"`
}

func (this *testLogContext) GetContext() *config.Context {
	return this.Context
}
func TestLogError(t *testing.T) {
	TestContext := new(testLogContext)
	cfg := rice.MustFindBox("../../conf")
	config.InitApp(TestContext, cfg)
	for i := 0; i < 100000; i++ {
		//go func(i int) {
		config.LogError("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", i)
		config.LogInfo("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", i)
		//}(i)
	}

	config.LogError("111")
	//for {
	//time.Sleep(time.Second)
	//}
}
