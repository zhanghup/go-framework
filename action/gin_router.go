package action

/*************************************************
 *	@author zander
 *	@time 	2018年1月2日14:12:37
 *	@brief 	gin路由控制器
 *************************************************/

import (
	"github.com/zhanghup/go-framework/pkg/gin"
	"reflect"
)

type RouteGroup struct {
	Path  string
	Func  []func(c *gin.Context)
	Group *gin.RouterGroup
}

//所有路由的集合
var groups = make([]*RouteGroup, 0)

//以下各分组的路由注册路口
func RegisterGroup(path string, fn ...func(c *gin.Context)) *gin.RouterGroup {
	group := new(gin.RouterGroup)
	rc := &RouteGroup{Path: path, Func: fn, Group: group}
	groups = append(groups, rc)
	return group
}

// 在初始化所有路由之前
var beforeRegisterGroups = make([]func(), 0)

func BeforeRegisterGroups(fn func()) {
	beforeRegisterGroups = append(beforeRegisterGroups, fn)
}

// 在初始化所有路由之后
var afterRegisterGroups = make([]func(), 0)

func AfterRegisterGroups(fn func()) {
	afterRegisterGroups = append(afterRegisterGroups, fn)
}

// 在初始化所有路由之后
var registerRouters = make([]func(), 0)

func RegisterRouters(fn func()) {
	registerRouters = append(registerRouters, fn)
}

// 实例化所有的实现的接口
func ginRoute() {
	for _, fn := range beforeRegisterGroups {
		fn()
	}
	for _, route := range groups {
		g := Gin.Group(route.Path)
		if route.Func != nil && len(route.Func) != 0 {
			for _, fn := range route.Func {
				g.Use(fn)
			}
		}
		reflect.ValueOf(route.Group).Elem().Set(reflect.ValueOf(g).Elem())
	}
	for _, fn := range afterRegisterGroups {
		fn()
	}
	for _, fn := range registerRouters {
		fn()
	}
}
