package app

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/ctx"
	"github.com/zhanghup/go-framework/tools"
	"testing"
)

func TestStartCommon(t *testing.T) {
	TestContext := &ctx.Context{}
	cfg := rice.MustFindBox("conf")
	StartCommon(TestContext, cfg)
	tools.PrintStruct(TestContext)
}

func TestStartCmd(t *testing.T) {
	TestContext := &ctx.Context{}
	cfg := rice.MustFindBox("conf")
	StartCmd(TestContext, cfg, nil)
}
