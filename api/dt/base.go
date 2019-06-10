package dt

import "github.com/zhanghup/go-framework/pkg/xorm"

type BaseService struct {
	DB *xorm.Engine
}

func NewBaseService(){

}
