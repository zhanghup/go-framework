package action

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-framework/context"
	"github.com/zhanghup/go-framework/pkg/gin"
	"io"
	"io/ioutil"
	"os"
)

var Gin *gin.Engine

func Run() error {
	ginRoute()
	// 读取配置
	ctx := context.GetAppConfig()
	if ctx == nil {
		panic("系统尚未初始化！")
	}
	port := ctx.System.HTTPPort
	if len(port) == 0 {
		port = ctx.Gin.HTTPPort
	}
	if ctx.Gin.TLS {
		ssl, err := rice.FindBox("conf")
		if err != nil {
			panic(err)

		}
		if _, err := os.Stat("./conf/ssl/server.crt"); os.IsNotExist(err) {
			bs, err := ssl.Bytes("ssl/server.crt")
			if err != nil {
				panic(err)

			}
			os.MkdirAll("./conf/ssl", 0666)
			ioutil.WriteFile("./conf/ssl/server.crt", bs, 0666)

		}

		if _, err := os.Stat("./conf/ssl/server.key"); os.IsNotExist(err) {
			bs, err := ssl.Bytes("ssl/server.key")
			if err != nil {
				panic(err)

			}
			os.MkdirAll("./conf/ssl", 0666)
			ioutil.WriteFile("./conf/ssl/server.key", bs, 0666)

		}
		return Gin.RunTLS(":"+port, "./conf/ssl/server.crt", "./conf/ssl/server.key")
	}
	return Gin.Run(":" + port)
}

func InitGin() {

	// gin日志
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(context.LogBean())

	// 创建对象
	Gin = gin.Default()

	return
}
