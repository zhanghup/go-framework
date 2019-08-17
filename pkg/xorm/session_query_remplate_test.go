package xorm

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhanghup/go-framework/pkg/mgo/bson"
	"testing"
	"time"
)

func TestSFQueryString(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type User struct {
		Id      int64  `xorm:"autoincr pk"`
		Msg     string `xorm:"varchar(255)"`
		Age     int
		Money   float32
		Created time.Time `xorm:"created"`
	}

	assert.NoError(t, testEngine.Sync2(new(User)))

	var data = User{
		Msg:   "hi",
		Age:   28,
		Money: 1.5,
	}
	_, err := testEngine.InsertOne(data)
	assert.NoError(t, err)

	obj := new(User)
	_, err = testEngine.SF(`select * from %s where msg = :msg`, "user", bson.M{"msg": "hi"}).get(obj)
	assert.NoError(t, err)
	_,err = testEngine.SF("delete  from user where msg = :msg",bson.M{"msg":"hi"}).Execute()
	assert.NoError(t, err)
}
