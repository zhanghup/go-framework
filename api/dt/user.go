package dt

import (
	"errors"
	app "github.com/zhanghup/go-framework"
	"github.com/zhanghup/go-framework/pkg/gin"
	"github.com/zhanghup/go-framework/pkg/xorm"
	"github.com/zhanghup/go-framework/tools"
)

type UserService struct {
	DB   *xorm.Engine
	Base *BaseService
}

// 用户注册逻辑
type UserRegisterParam struct {
	Account    *string `json:"account" ck:"true"`
	Password   *string `json:"password" ck:"true"`
	RePassword *string `json:"re_password" ck:"true"`
}

func (this *UserService) Register(param UserRegisterParam) (err error) {
	s := this.DB.NewSession()
	user := new(app.User)
	ok, err := s.SF(`select * from {{ table "user" }} u where u.account = :account and u.status = 1`, map[string]interface{}{
		"account": param.Account,
	}).Get(user)
	if err != nil {
		return
	}
	if ok {
		err = errors.New("账户已经存在")
		return
	}

	user.Id = tools.ObjectString()
	user.Account = param.Account
	user.Slat = tools.ObjectString()
	user.Password = tools.PtrOfString(tools.Password(*param.Password, *user.Slat))
	user.Status = tools.PtrOfInt(1)
	_, err = s.Insert(user)
	return
}

// 用户名密码登录逻辑
type UserAccountLogin struct {
	Account  *string `json:"account" ck:"true"`
	Password *string `json:"password" ck:"true"`
	Code     *string `json:"code"` // 验证码
}

func (this *UserService) Login(c *gin.Context, param UserAccountLogin) (token interface{}, err error) {
	user := app.User{}
	ok, err := this.DB.Where("account = ? and status = 1", param.Account).Get(&user)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("用户不存在")
	}

	flag := false
	if user.Slat == nil {
		if *user.Password == *param.Password {
			flag = true
		}
	} else {
		if *user.Password == tools.Password(*param.Password, *user.Slat) {
			flag = true
		}
	}
	if flag {
		tok, err := this.Base.Token(c, *user.Id, "pc")
		return tok.Id, err
	}
	return nil, errors.New("登录失败")
}

type UserWxLogin struct {

}

var userService *UserService

func NewUserService(e *xorm.Engine) *UserService {
	if userService != nil {
		return userService
	}
	userService = &UserService{DB: e, Base: NewBaseService(e)}
	return userService
}
