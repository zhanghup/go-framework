package app

type Bean struct {
	Id      *string `json:"id" xorm:"pk"`
	Created *int64  `json:"created" xorm:"created"`
	Updated *int64  `json:"updated" xorm:"updated"`
	Weight  *int    `json:"weight" xorm:"weight"`
	Status  *int    `json:"status" xorm:"status"`
}
