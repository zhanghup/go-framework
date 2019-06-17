package deploy

import (
	"github.com/zhanghup/go-framework/tools"
	"os"
)

func InitConfigFolder() {
	_ = os.MkdirAll("./resource/static", os.ModePerm)
	_, err := os.Open("conf")
	if os.IsExist(err) {
		return
	}
	tools.RiceWirteToLocal("conf")

}
