package api

import (
	app "github.com/zhanghup/go-framework"
	cfg2 "github.com/zhanghup/go-framework/ctx/cfg"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/wechat/wx"
	"github.com/zhanghup/go-framework/pkg/wechat/wx/message"
	"github.com/zhanghup/go-framework/pkg/xorm"
)

func WcCallback(g *gin.RouterGroup, e *xorm.Engine) {

	cfg := cfg2.GetCfg()
	//配置微信参数
	config := &wx.Config{
		AppID:          cfg.Wc.AppID,
		AppSecret:      cfg.Wc.AppSecret,
		Token:          cfg.Wc.Token,
		EncodingAESKey: cfg.Wc.EncodingAESKey,
		PayKey:         cfg.Wc.PayKey,
		PayMchID:       cfg.Wc.PayMchID,
		PayNotifyURL:   "",
	}
	wc := wx.NewWechat(config)

	g.GET("/wc/callback", func(c *gin.Context) {
		// 传入request和responseWriter
		server := wc.GetServer(c.Request, c.Writer)

		//处理消息接收以及回复
		err := server.Serve()
		if err != nil {
			app.LogError("微信公众好处理消息接收以及回复 - " + err.Error())
			return
		}
		//发送回复的消息
		server.Send()
	})

	g.POST("/wc/callback", func(c *gin.Context) {
		// 传入request和responseWriter
		server := wc.GetServer(c.Request, c.Writer)

		//设置接收消息的处理方法
		server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

			//回复消息：演示回复用户发送的消息
			text := message.NewText(msg.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		})

		//处理消息接收以及回复
		err := server.Serve()
		if err != nil {
			app.LogError("微信公众好处理消息接收以及回复 - " + err.Error())
			return
		}
		//发送回复的消息
		server.Send()
	})
}
