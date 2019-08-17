package api

import (
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/api/dt"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/mgo/bson"
	"github.com/zhanghup/go-framework/pkg/xorm"
)

type DictParam struct {
}
type DictItemParam struct {
	Code string `json:"code"`
}

func WebBase(g *gin.RouterGroup, e *xorm.Engine) {
	{ // 数据字典
		dt.NewBaseService(e).CRUD(g, "/dict", new(app.Dict), dt.CRUD_CREATE, dt.CRUD_UPDATE,dt.CRUD_GET)
		g.Action("/dict/delete", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
			b := p.(*app.Bean)
			err = e.T(func(s *xorm.Session) error {
				_, err := e.SF(`delete from {{ table "dict" }} where id = :id`, bson.M{"id": b.Id}).Execute()
				if err != nil {
					return err
				}
				_, err = e.SF(`delete from {{ table "dict_item" }} where code = :code`, bson.M{"code": b.Id}).Execute()
				return err
			})
			return
		}, new(app.Bean))
		dt.NewBaseService(e).CRUD(g, "/dictitem", new(app.DictItem), dt.CRUD_CREATE, dt.CRUD_UPDATE, dt.CRUD_DELETE,dt.CRUD_GET)
		g.Action("/dict/list", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
			return dt.NewSysService(e).DictList()
		})
		g.Action("/dictitem/list", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
			param := p.(*DictItemParam)
			return dt.NewSysService(e).DictItems(param.Code)
		}, new(DictItemParam))
	}

	{ // 用户管理
		dt.NewBaseService(e).CRUD(g, "/user", new(app.User), dt.CRUD_GET, dt.CRUD_CREATE, dt.CRUD_UPDATE, dt.CRUD_DELETE)
		g.Action("/user/list", func(c *gin.Context, p interface{}) (obj interface{}, err error) {
			return dt.NewSysService(e).DictList()
		})
	}
}
