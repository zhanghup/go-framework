package action

import (
	"io"
	"io/ioutil"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-framework/app"
)

var Gin *gin.Engine

func InitGin() {
	gin.DefaultWriter = io.MultiWriter(app.LogWriter(), os.Stdout)

	Gin = gin.New()
	ctx := app.GetAppConfig()
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
		Gin.RunTLS(":"+port, "./conf/ssl/server.crt", "./conf/ssl/server.key")

	} else {
		Gin.Run(":" + port)

	}
	return
}
