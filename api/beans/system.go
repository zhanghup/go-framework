package beans

import app "github.com/zhanghup/go-framework"

type Dict struct {
	app.Bean `xorm:"extends"`

	Name   *string `json:"name"`
	Remark *string `json:"remark"`
}

type DictItem struct {
	app.Bean `xorm:"extends"`

	Code      *string `json:"code"`
	Name      *string `json:"name"`
	Value     *string `json:"value"`
	Extension *string `json:"extension"`
}

type Menu struct {
	app.Bean `xorm:"extends"`

	Title  *string `json:"title"`
	Name   *string `json:"name"`
	NameEn *string `json:"name_en"`
	Index  *string `json:"index"`
	Icon   *string `json:"icon"`
	Parent *string `json:"parent"`
}

type Role struct {
	app.Bean `xorm:"extends"`

	Name *string `json:"name"`
	Desc *string `json:"desc"`
}

type RoleUser struct {
	app.Bean `xorm:"extends"`

	Role *string `json:"role"`
	User *string `json:"user"`
}

type Perm struct {
	app.Bean `xorm:"extends"`

	Type *string `json:"type"` // 类型（menu等）
	Role *string `json:"role"` // 角色ID
	Oid  *string `json:"oid"`  // 对象ID
	Mask *string `json:"mask"` // 权限
}
