package tools

import (
	"encoding/json"
	"fmt"
	"time"
)

func Time() string {
	return time.Now().Format("15:04:05")
}
func Date() string {
	return time.Now().Format("2006-01-02")
}
func Month() string {
	return time.Now().Format("2006-01")
}
func Year() string {
	return time.Now().Format("2006")
}

func PtrOfInt(i int) *int {
	return &i

}

func PtrOfString(i string) *string {
	return &i

}

func PtrOfInt64(i int64) *int64 {
	return &i

}

func PtrOfFloat64(i float64) *float64 {
	return &i

}

// 以json格式输出struct对象
func PrintStruct(obj interface{}) string {

	r, err := json.Marshal(obj)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(r))
	return string(r)

}

// 以json格式输出struct对象
//func PrintStructFmt(obj interface{}) string {
//str, flag := sPrintStruct(obj)
//if !flag {
//fmt.Println(str)
//return str
//}
//r, err := json.Marshal(obj)
//if err != nil {
//panic(err)
//}

//var out bytes.Buffer
//err = json.Indent(&out, r, "", "\t")
//out.WriteTo(os.Stdout)
//return string(out.Bytes())
//}
func StrContains(src []string, tag string) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}
