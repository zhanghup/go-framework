package app

import (
	rice "github.com/GeertJohan/go.rice"
	cfg2 "github.com/zhanghup/go-framework/ctx/cfg"
	"github.com/zhanghup/go-framework/tools"
	"testing"
)

func TestStartCommon(t *testing.T) {
	TestContext := &cfg2.Cfg{}
	cfg := rice.MustFindBox("conf")
	StartCommon(TestContext, cfg)
	tools.PrintStruct(TestContext)

}
