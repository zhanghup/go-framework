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

