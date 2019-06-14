package xorm

import (
	"database/sql"
	"github.com/go-xorm/core"
	"reflect"
)

func (engine *Engine) T(fn func(s *Session) error) error {
	session := engine.NewSession()
	err := fn(session)
	if err != nil {
		err := session.Rollback()
		return err
	}
	return session.Commit()
}

func (engine *Engine) TO(fn func(s *Session) (interface{}, error)) (interface{}, error) {
	session := engine.NewSession()
	obj, err := fn(session)
	if err != nil {
		er := session.Rollback()
		if er != nil {
			return obj, er
		}
		return obj, err
	}
	return obj, session.Commit()
}

func (session *Session) Execute() (sql.Result, error) {
	defer session.resetStatement()
	if session.isAutoClose {
		defer session.Close()
	}

	sqlStr := session.statement.RawSQL
	params := session.statement.RawParams

	i := len(params)
	if i == 1 {
		vv := reflect.ValueOf(params[0])
		if vv.Kind() != reflect.Ptr || vv.Elem().Kind() != reflect.Map {
			return session.exec(sqlStr, params...)
		} else {
			sqlStr1, args, _ := core.MapToSlice(sqlStr, params[0])
			return session.exec(sqlStr1, args...)
		}
	} else {
		return session.exec(sqlStr, params...)
	}
}
