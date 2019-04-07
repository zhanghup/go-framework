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
