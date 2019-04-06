package tools

import (
	"fmt"
	"reflect"
	"strconv"

	"gopkg.in/ini.v1"
)

func IniToInterface(cfg *ini.File, obj interface{}) interface{} {
	RftStructDeep(obj, func(ty reflect.Type, vl reflect.Value, tg reflect.StructTag) bool {
		if ty.Kind() == reflect.Struct && StrContains(cfg.SectionStrings(), tg.Get("json")) {
			sec := cfg.Section(tg.Get("json"))
			for i := 0; i < ty.NumField(); i++ {
				t := ty.Field(i).Type
				tag := ty.Field(i).Tag
				v := vl.Field(i)
				if v.CanSet() {
					tgv, _ := sec.GetKey(tag.Get("json"))
					if t.Kind() == reflect.Ptr {
						v.Set(reflect.New(t.Elem()))
						t = t.Elem()
						v = v.Elem()

					}
					switch t.Kind() {
					case reflect.String:
						if sec.HasKey(tag.Get("json")) && len(tgv.String()) >= 0 {
							v.SetString(tgv.String())
						} else {
							fmt.Println(tag.Get("cfg"))
							v.SetString(tag.Get("cfg"))
						}

					case reflect.Int:
						if sec.HasKey(tag.Get("json")) && len(tgv.String()) >= 0 {
							vv, _ := tgv.Int64()
							v.SetInt(vv)
						} else if len(tag.Get("cfg")) > 0 {
							si, err := strconv.Atoi(tag.Get("cfg"))
							if err == nil {
								v.SetInt(int64(si))
							}
						}
					case reflect.Int64:
						if sec.HasKey(tag.Get("json")) && len(tgv.String()) >= 0 {
							vv, _ := tgv.Int64()
							v.SetInt(vv)
						} else if len(tag.Get("cfg")) > 0 {
							si, err := strconv.ParseInt(tag.Get("cfg"), 10, 64)
							if err == nil {
								v.SetInt(si)
							}
						}
					case reflect.Float64:
						if sec.HasKey(tag.Get("json")) && len(tgv.String()) >= 0 {
							vv, _ := tgv.Float64()
							v.SetFloat(vv)
						} else if len(tag.Get("cfg")) > 0 {
							si, err := strconv.ParseFloat(tag.Get("cfg"), 64)
							if err == nil {
								v.SetFloat(si)
							}
						}
					case reflect.Bool:
						if sec.HasKey(tag.Get("json")) && len(tgv.String()) >= 0 {
							vv, _ := tgv.Bool()
							v.SetBool(vv)
						} else {
							if tag.Get("cfg") == "true" {
								v.SetBool(true)
							} else {
								v.SetBool(false)
							}
						}
					}
				}
			}
		}
		return true
	})
	return obj
}
