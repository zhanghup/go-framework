package main

import (
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/app"
)

type testLogContext struct {
	*app.Context
	System struct {
		HttpPort string `json:"http-port"`
	} `json:"system"`
}

func (this *testLogContext) GetContext() *app.Context {
	return this.Context
}
func TestLogError(t *testing.T) {
	TestContext := new(testLogContext)
	cfg := rice.MustFindBox("../../conf")
	app.InitApp(TestContext, cfg)
	for i := 0; i < 100000; i++ {
		//go func(i int) {
		app.LogError("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", i)
		app.LogInfo("滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答滴滴答答的等待滴滴答答滴滴答答滴滴答答 %v", i)
		//}(i)
	}

	app.LogError("111")
	//for {
	//time.Sleep(time.Second)
	//}
}
