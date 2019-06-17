package api

import (
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/api/dt"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
	"github.com/zhanghup/go-framework/tools"
)

func Login(g *gin.RouterGroup, e *xorm.Engine) {
	user := new(app.User)
	ok, err := e.SF(`select * from {{ table "user" }} where id = 'root'`).Get(user)
	if err == nil && !ok {
		user.Id = tools.PtrOfString("root")
		user.Account = tools.PtrOfString("root")
		user.Password = tools.PtrOfString("Aa123456")
		user.Status = tools.PtrOfInt(1)
		user.Admin = tools.PtrOfInt(1)
		_, _ = e.Insert(user)
		return
	}

	g.Action("/register/account", func(c *gin.Context, param interface{}) (obj interface{}, err error) {
		p := param.(dt.UserRegisterParam)
		return nil, dt.NewUserService(e).Register(p)
	}, dt.UserRegisterParam{})

	g.Action("/login", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
		param := p.(dt.UserAccountLogin)
		return dt.NewUserService(e).Login(c, param)
	}, new(dt.UserAccountLogin))

	g.Action("/logout", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
		return nil, dt.NewBaseService(e).TokenDelete(c)
	})

	g.Action("/login/wx", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
		return nil, nil
	})
	g.Action("/login/wc", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
		return nil, nil
	})
}
