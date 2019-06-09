package beans

import app "github.com/zhanghup/go-framework"

type User struct {
	app.Bean `xorm:"extends"`

	Type     *string `json:"type"`
	Account  *string `json:"account"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	Avatar   *string `json:"avatar"`
	IdCard   *string `json:"id_card"`
	Birth    *int64  `json:"birth"`
	Sex      *string `json:"sex"`   // 0：未知，1：男，2：女
	Mobile   *string `json:"phone"` // 联系电话
	Admin    *int    `json:"admin"`
}

type UserToken struct {
	app.Bean `xorm:"extends"`

	Token  *string `json:"token"`  // 授权码
	Ops    *int64  `json:"ops"`    // 接口调用次数
	Type   *string `json:"type"`   // 授权类型 [pc,微信小程序，微信公众号]
	Expire *string `json:"expire"` // 到期时间
	Agent  *string `json:"agent"` // User-Agent
}
