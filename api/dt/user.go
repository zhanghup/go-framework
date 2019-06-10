package dt

import "github.com/zhanghup/go-framework/pkg/xorm"

type UserService struct {
	DB   *xorm.Engine
	Base *BaseService
}


