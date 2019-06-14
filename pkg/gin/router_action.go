package gin

import (
	"github.com/zhanghup/go-framework/tools"
	"reflect"
)

func (group *RouterGroup) Action(relativePath string, fn func(c *Context, p interface{}) (obj interface{}, err error), param ...interface{}) IRoutes {
	return group.handle("POST", relativePath, []HandlerFunc{func(c *Context) {
		var obj interface{}
		var err error

		// 若不需要绑定任何参数
		if len(param) == 0 {
			obj, err = fn(c, nil)
		} else {
			// 若包含参数
			ty := reflect.TypeOf(param[0])
			if ty.Kind() == reflect.Ptr {
				ty = ty.Elem()
			}
			pp := reflect.New(ty).Interface()
			// 读取参数
			err = c.BindJSON(pp)
			if err != nil {
				Fail(c, NewError(9998, err.Error(), pp), pp)
				return
			}

			flag := false
			tools.RftStructDeep(pp, func(ty reflect.Type, vl reflect.Value, tg reflect.StructTag) bool {
				if tg.Get("ck") == "true" {
					switch ty.Kind() {
					case reflect.Ptr:
						if vl.Pointer() == 0 {
							flag = true
							return false
						}
					case reflect.String:
						if len(vl.String()) == 0 {
							flag = true
							return false
						}
					}
					return true
				} else {
					if ty.Kind() == reflect.Ptr && vl.Pointer() == 0 && vl.CanSet() {
						vl.Set(reflect.New(ty.Elem()))
					}
					return true
				}
			})

			if flag {
				Fail(c, NewError(9998, "参数校验错误", pp), pp)
				return
			}
			obj, err = fn(c, reflect.ValueOf(pp).Elem().Interface())
			if err != nil {
				Fail(c, err, obj)
				return
			}
		}

		c.JSON(200, exception{
			Code:    0,
			Message: "ok",
			Data:    obj,
		})
	}})
}

func Fail(c *Context, err error, obj interface{}) {
	ty := reflect.TypeOf(err)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	if ty.String() == "gin.exception" {
		c.JSON(200, err)
		return
	} else {
		c.JSON(200, exception{
			Code:    9999,
			Message: "系统异常",
			Data:    map[string]interface{}{"error": err.Error()},
		})
	}
}
