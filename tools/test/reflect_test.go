package test

import (
	"reflect"
	"testing"

	"github.com/zhanghup/go-framework/tools"
)

func TestReflect(t *testing.T) {
	type Ga struct {
		FF string `json:"ff"`
		AD string `json:"ad"`
	}
	tt := struct {
		Name  *string `json:"name"`
		Value *int    `json:"value"`
		Key   float64 `json:"key"`
		Sa    Ga      `json:"sa"`
	}{
		Key: 0.2,
		Sa: Ga{
			FF: "123",
		},
	}
	tools.RftStructDeep(&tt, func(t1 reflect.Type, v1 reflect.Value, tg reflect.StructTag) bool {
		if tg.Get("json") == "sa" {
			v1.Set(reflect.ValueOf(Ga{FF: "helo", AD: "world"}))
			return false
		}
		return true
	})

	tools.PrintStruct(tt)
}
