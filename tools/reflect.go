package tools

import (
	"reflect"
)

func RftStructDeep(obj interface{}, fn func(t1 reflect.Type, v1 reflect.Value, tg reflect.StructTag) bool) {
	if fn == nil || obj == nil {
		panic("参数不能为空")
	}
	ty := reflect.TypeOf(obj)
	vl := reflect.ValueOf(obj)
	rftSelfDeep(ty, vl, "", fn)
}

func rftSelfDeep(ty reflect.Type, vl reflect.Value, tg reflect.StructTag, fn func(rty reflect.Type, rvl reflect.Value, rtg reflect.StructTag) bool) {
	switch ty.Kind() {
	case reflect.Ptr:
		if vl.Pointer() == 0 && vl.CanSet() {
			if !fn(ty, vl, tg) {
				return
			}
		}
		ty = ty.Elem()
		vl = vl.Elem()
		rftSelfDeep(ty, vl, tg, fn)

	case reflect.Struct:
		if !vl.CanSet() {
			return
		}
		if !fn(ty, vl, tg) {
			return
		}
		for i := 0; i < ty.NumField(); i++ {
			t := ty.Field(i).Type
			tag := ty.Field(i).Tag
			v := vl.Field(i)
			rftSelfDeep(t, v, tag, fn)

		}
	case reflect.Map:
	case reflect.Slice:
	case reflect.Func:
	case reflect.Chan:
	default:
		if vl.CanSet() {
			fn(ty, vl, tg)
		}
	}

}
