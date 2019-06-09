package gin

import "reflect"

func (group *RouterGroup) Action(relativePath string, fn func() (interface{}, error)) IRoutes {
	return group.handle("POST", relativePath, []HandlerFunc{func(c *Context) {
		obj, err := fn()
		if err != nil {
			ty := reflect.TypeOf(err)
			if ty.Kind() == reflect.Ptr {
				ty = ty.Elem()
			}
			if ty.String() == "exception" {
				c.JSON(200, err)
				return
			} else {
				c.JSON(200, exception{
					Code:    9999,
					Message: "ok",
					Data:    obj,
				})
			}
		}
		c.JSON(200, exception{
			Code:    0,
			Message: "ok",
			Data:    obj,
		})
	}})
}
