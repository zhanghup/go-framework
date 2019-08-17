package dt

import (
	"errors"
	"fmt"
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
	"github.com/zhanghup/go-framework/tools"
	"reflect"
)

type crud_type string

const (
	TOKEN_COOKIE = "__zander_token__"
	TOKEN_AJAX   = "Z-Token"

	CRUD_CREATE crud_type = "C"
	CRUD_GET    crud_type = "R"
	CRUD_UPDATE crud_type = "U"
	CRUD_DELETE crud_type = "D"
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
		c.SetCookie(TOKEN_COOKIE, *token.Id, 2*60*60, "/", "", false, true)
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
func (this *BaseService) CRUD(g *gin.RouterGroup, prefix string, inf interface{}, ty ...crud_type) {
	if reflect.TypeOf(inf).Kind() != reflect.Ptr {
		panic("输入必须为指针类型")
	}
	tymap := map[crud_type]bool{}
	for _, str := range ty {
		tymap[str] = true
	}

	if ok1, ok2 := tymap[CRUD_GET]; ok1 && ok2 {
		g.Action(prefix+"/get", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
			param := p.(*app.Bean)
			result := reflect.New(reflect.TypeOf(inf).Elem()).Interface()
			ok, err := this.DB.Table(inf).Where("id = ?", param.Id).Get(result)
			if err != nil {
				return nil, err
			}
			if !ok {
				return nil, errors.New("对象不存在")
			}
			return result, nil

		}, new(app.Bean))
	}

	if ok1, ok2 := tymap[CRUD_CREATE]; ok1 && ok2 {
		g.Action(prefix+"/create", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
			tools.RftStructDeep(p, func(t1 reflect.Type, v1 reflect.Value, tg reflect.StructTag, fieldName string) bool {
				if fieldName == "Id" && v1.CanSet() && len(v1.String()) == 0 {
					v1.Set(reflect.ValueOf(*tools.ObjectString()))
					return false
				}
				if fieldName == "Status" && v1.CanSet() {
					v1.Set(reflect.ValueOf(1))
					return false
				}
				return true
			})
			_, err = this.DB.Table(inf).Insert(p)
			return p, err
		}, inf)
	}

	if ok1, ok2 := tymap[CRUD_UPDATE]; ok1 && ok2 {
		g.Action(prefix+"/update", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
			id := ""
			fmt.Println(reflect.TypeOf(p).String())
			tools.RftStructDeep(p, func(t1 reflect.Type, v1 reflect.Value, tg reflect.StructTag, fieldName string) bool {
				fmt.Println(t1.String(), fieldName)
				if t1.Kind() == reflect.String && fieldName == "Id" && v1.CanSet() {
					id = v1.String()
					return false
				}
				return true
			})
			if len(id) == 0 {
				return nil, errors.New("id不能为空")
			}
			_, err = this.DB.Table(inf).Where("id = ?", id).AllCols().Update(p)
			return p, err
		}, inf)
	}

	if ok1, ok2 := tymap[CRUD_DELETE]; ok1 && ok2 {
		g.Action(prefix+"/delete", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
			param := p.(*app.Bean)
			tab := this.DB.TableInfo(inf).Name
			_, err = this.DB.SF(`delete from %s where id = :id`, tab, map[string]interface{}{"id": param.Id}).Execute()
			return
		}, new(app.Bean))
	}
}

func Auth(e *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok := c.GetHeader(TOKEN_AJAX)
		var err error
		if len(tok) == 0 {
			tok, err = c.Cookie(TOKEN_COOKIE)
			if err != nil {
				c.Fail(gin.NewErr("[0] 未授权"), nil, 401)
				c.Abort()
				return
			}
			if len(tok) == 0 {
				c.Status(401)
				c.Abort()
				c.Fail(gin.NewErr("[1] 未授权"), 401)
				return
			}
		}

		token := new(app.UserToken)
		ok, err := e.SF(`select * from {{ table "user_token" }} where status = 1 and id = :id`, map[string]interface{}{
			"id": tok,
		}).Get(token)

		if !ok {
			c.Status(401)
			c.Abort()
			c.Fail(gin.NewErr("[2] 未授权"), 401)
			return
		}
		c.Next()
	}

}

var baseService *BaseService

func NewBaseService(e *xorm.Engine) *BaseService {
	if baseService != nil {
		return baseService
	}
	baseService = &BaseService{DB: e}
	return baseService
}
