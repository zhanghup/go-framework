package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/zhanghup/go-framework/pkg/mgo/bson"
	"time"
)

func ObjectString() *string {
	str := bson.NewObjectId().Hex()
	return &str
}
func Password(password, slat string) string {
	sh := sha256.New()
	sh.Write([]byte(password))
	bts := sh.Sum([]byte(slat))
	return fmt.Sprintf("%x", bts)
}
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

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

func StrContains(src []string, tag string) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}
