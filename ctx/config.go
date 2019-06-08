package ctx

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/tools"
	"gopkg.in/ini.v1"
)

type IContext interface {
	GetContext() *Context
}

// 框架基本配置对象
type Context struct {
	// 包含一些系统级的通用配置，system 中配置的优先级最高
	System struct {
		Name     string `json:"name"`
		Brief    string `json:"brief"`
		Version  string `json:"version"`
		HTTPPort string `json:"http-port"`
	} `json:"system"`
	Gin struct {
		Enable   bool   `json:"enable" cfg:"true"`
		HTTPPort string `json:"http-port" cfg:"40018"`
		Gzip     bool   `json:"gizp" cfg:"true"`
		TLS      bool   `json:"tls" cfg:"false"`
	} `json:"gin"`
	Log struct {
		Filename   string `json:"filename" cfg:"./resource/log/app.log"`
		MaxBackups int    `json:"max-backup" cfg:"0"`
		MaxSize    int    `json:"max-size" cfg:"10"`
		MaxAge     int    `json:"max-age" cfg:"0"`
		Level      int    `json:"level" cfg:"1"`
	} `json:"log"`
}

func (this *Context) GetContext() *Context {
	return this
}

var appconfig *Context

// 初始化框架的配置文件
func InitApp(afg IContext, box *rice.Box) {

	f, err := box.Open("config-default.ini")
	if err != nil {
		panic("找不到配置文件")
	}
	cfg, _ := ini.Load(f)
	tools.IniToInterface(cfg, afg)
	f, _ = box.Open("config.ini")
	if f != nil {
		cfg, _ := ini.Load(f)
		tools.IniToInterface(cfg, afg)
	}
	appconfig = afg.GetContext()

	// 初始化日志系统
	SetLogConfig(afg)
}

func GetAppConfig() *Context {
	return appconfig
}
