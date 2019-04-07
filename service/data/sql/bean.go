package sql

import (
	"gopkg.in/mgo.v2/bson"
)


type Bean struct {
	Id      string `json:"id" xorm:"pk"`
	Created string `json:"created" xorm:"created"`
	Updated string `json:"updated" xorm:"updated"`
	Status  int    `json:"status" xorm:"default(1)"`
}

func (bean *Bean) BeforeInsert() {
	if len(bean.Id) == 0 {
		bean.Id = bson.NewObjectId().Hex()
	}
}

//同步表结构
var tableBeans = make([]interface{}, 0)
//以下各分组的路由注册路口
func InitTableBeans(obj ... interface{}) {
	for _, o := range obj {
		tableBeans = append(tableBeans, o)
	}
}


