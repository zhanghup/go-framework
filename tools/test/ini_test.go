package test

import (
	"testing"

	"github.com/zhanghup/go-framework/tools"
	"gopkg.in/ini.v1"
)

func TestIniToJson(t *testing.T) {
	tt := &struct {
		System struct {
			HTTPPort *string `json:"http-port"`
		} `json:"system"`
		Log struct {
			MaxSize string `json:"max-size"`
		} `json:"log"`
	}{}
	cfg, _ := ini.Load("../../conf/config-default.ini")
	tools.IniToInterface(cfg, tt)
	tools.PrintStruct(tt)
}
