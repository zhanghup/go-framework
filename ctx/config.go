package ctx

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/mgo"
	"github.com/zhanghup/go-framework/pkg/xorm"
	"github.com/zhanghup/go-framework/tools"
	"gopkg.in/ini.v1"
)

type ICfg interface {
	GetCfg() *Cfg
}

// 框架基本配置对象
type Cfg struct {
	// 包含一些系统级的通用配置，system 中配置的优先级最高
	System struct {
		Name     string `json:"name" cfg:"zander框架服务"`
		Brief    string `json:"brief" cfg:"zander框架服务"`
		Version  string `json:"version" cfg:"0.0.1"`
		HTTPPort string `json:"http-port" cfg:"40018"`
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

	Engine struct {
		Gin  *gin.Engine
		Xorm *xorm.Engine
		Mgo  *mgo.Database
	}
}

func (this *Cfg) GetCfg() *Cfg {
	return this
}

var appconfig *Cfg

// 初始化框架的配置文件
func InitCfg(afg ICfg, box *rice.Box) {

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
	appconfig = afg.GetCfg()

	// 初始化日志系统
	SetLogConfig(afg)
}

func GetCfg() *Cfg {
	return appconfig
}
