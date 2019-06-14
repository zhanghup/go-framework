package dt

import (
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
	"github.com/zhanghup/go-framework/tools"
)

const (
	TOKEN_COOKIE = "__zander_token__"
	TOKEN_AJAX   = "Z-Token"
)

type BaseService struct {
	DB *xorm.Engine
}

func (this *BaseService) Token(c *gin.Context, uid, typ string) (*app.UserToken, error) {
	token := new(app.UserToken)
	_, err := this.DB.TO(func(s *xorm.Session) (i interface{}, e error) {
		_, e = s.SF(`update {{ table "user_token" }} set status = 0 where user = :user`, map[string]interface{}{
			"user": uid,
		}).Execute()
		if e != nil {
			return
		}
		token.Id = tools.ObjectString()
		token.Status = tools.PtrOfInt(1)
		token.User = &uid
		token.Agent = tools.PtrOfString(c.Request.UserAgent())
		token.Expire = tools.PtrOfInt64(7 * 24 * 60 * 60)
		token.Type = &typ
		token.Ops = tools.PtrOfInt64(0)
		i, e = s.Insert(token)
		if e != nil {
			return
		}
		c.SetCookie(TOKEN_COOKIE, *token.Id, 2*60*60, "/", "localhost", false, true)
		return
	})
	return token, err
}
func (this *BaseService) TokenDelete(c *gin.Context) error {
	tok := c.GetHeader(TOKEN_AJAX)
	var err error
	if len(tok) == 0 {
		tok, err = c.Cookie(TOKEN_COOKIE)
		if err != nil {
			return err
		}
	}
	_, err = this.DB.SF(`update {{ table "user_token" }} set status = 0 where id = :id`, map[string]interface{}{
		"id": tok,
	}).Execute()
	if err != nil {
		return err
	}
	c.SetCookie("__zander_token__", "", 1, "/", "localhost", false, true)
	return nil
}

var baseService *BaseService

func NewBaseService(e *xorm.Engine) *BaseService {
	if baseService != nil {
		return baseService
	}
	baseService = &BaseService{DB: e}
	return baseService
}
