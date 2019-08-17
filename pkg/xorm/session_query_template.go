package xorm

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const TEMP_KEY = `{{\s*%v\s*"\s*(\S+)\s*"\s*}}`
const TEMP_OPERATE = `{{\s*%v\s*"\s*(\S+)\s*"\s*}}(.*){{\s*end\s*}}`

var (
	temp_table, _    = regexp.Compile(fmt.Sprintf(TEMP_KEY, "table"))
	temp_orderby, _  = regexp.Compile(fmt.Sprintf(TEMP_KEY, "orderby"))
	temp_sqlblock, _ = regexp.Compile(fmt.Sprintf(TEMP_KEY, "sqlblock"))
	temp_params, _   = regexp.Compile(`:(\S+)`)
	//temp_if, _       = regexp.Compile(fmt.Sprintf(TEMP_OPERATE, "if"))
	temp_NotEmpty, _ = regexp.Compile(fmt.Sprintf(TEMP_OPERATE, "notempty"))
)

var sqltemplateCatch = NewMemoryStore()

func (engine *Engine) _teble(name string) string {
	return engine.TableMapper.Obj2Table(name)
}
func (engine *Engine) _orderby(name string, param map[string]interface{}) string {
	ordermap, ok := param[name]
	if !ok {
		return ""
	}
	var orders []string
	orders, ok = ordermap.([]string)
	if !ok {
		t, ok := ordermap.(*[]string)
		if !ok {
			return ""
		}
		orders = *t
	}

	if len(orders) == 0 {
		return ""
	}

	vs := []string{}

	for i := range orders {

		if strings.HasPrefix(orders[i], "-") {
			vs = append(vs, engine.Quote(orders[i][1:])+" DESC")
		} else {
			vs = append(vs, engine.Quote(orders[i])+" ASC")
		}
	}

	return "ORDER BY " + strings.Join(vs, ",")
}

func (engine *Engine) _sqlblock(name string) string {
	str, err := sqltemplateCatch.Get(fmt.Sprintf("_sqlblock_%s", name))
	if err != nil {
		return ""
	}
	return str.(string)
}

func (engine *Engine) _sql_format_params(name string, p interface{}) (string, []interface{}) {
	if p == nil {
		return fmt.Sprintf("# %s nil #", name), nil
	}
	ty := reflect.TypeOf(p)
	vl := reflect.ValueOf(p)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
		vl = vl.Elem()
	}

	if ty.Kind() == reflect.Slice {
		a, b := make([]string, 0), make([]interface{}, 0)
		c := vl.Len()
		for i := 0; i < c; i++ {
			a = append(a, "?")
			b = append(b, vl.Index(i).Interface())
		}
		return fmt.Sprintf("(%s)", strings.Join(a, ",")), b
	}

	return "?", []interface{}{p}
}

func (this *Engine) _sqlFormatParams(sql string, param map[string]interface{}) (string, []interface{}) {
	// 格式化参数模块
	params := make([]interface{}, 0)
	for {
		rs := temp_params.FindStringSubmatch(sql)
		if len(rs) > 1 && len(rs[1]) > 1 {
			a, b := this._sql_format_params(rs[1], param[rs[1]])
			sql = strings.Replace(sql, rs[0], a, 1)
			params = append(params, b...)
		} else {
			break
		}
	}
	return sql, params
}
func (this *Engine) _sqlFormatOperate(sql string, param map[string]interface{}) string {
	//  格式化 notempty 表达式
	for {
		rs := temp_NotEmpty.FindStringSubmatch(sql)
		if len(rs) > 2 {
			_, b := this._sql_format_params(rs[1], param[rs[1]])
			if len(b) == 0 {
				sql = strings.Replace(sql, rs[0], "", 1)
			} else {
				sql = strings.Replace(sql, rs[0], rs[2], 1)
			}
		} else {
			break
		}
	}
	return sql
}
func (this *Engine) _sqlFormatAction(sql string, param map[string]interface{}) string {

	//  格式化table
	for {
		rs := temp_table.FindStringSubmatch(sql)
		if len(rs) > 1 && len(rs[1]) > 1 {
			sql = strings.Replace(sql, rs[0], this._teble(rs[1]), 1)
		} else {
			break
		}
	}

	// 格式化sqlblock
	for {
		rs := temp_sqlblock.FindStringSubmatch(sql)
		if len(rs) > 1 && len(rs[1]) > 1 {
			sql = strings.Replace(sql, rs[0], this._sqlblock(rs[1]), 1)
		} else {
			break
		}
	}

	// 格式化orderby
	for {
		rs := temp_orderby.FindStringSubmatch(sql)
		if len(rs) > 1 && len(rs[1]) > 1 {
			sql = strings.Replace(sql, rs[0], this._orderby(rs[1], param), 1)
		} else {
			break
		}
	}
	return sql
}

/*
Two types to search
	The No.1
		sql:	select * from myuser where id = ? and name = ?
		args:  	[]string{"5cc79f78e1382321ef2a5ae9","zander"}

	The No.2
		sql: 	select * from {{ table "myuser" }} where id = :id and name = :name
		args:	[]interfaces{}{
					map[string]interface{}{
						"id": 	"5cc79f78e1382321ef2a5ae9",
						"name": "zander",
					}
				}

Template types:
		{{ table "user" }}
			return:		"test_user"
		{{ orderby "users" }}
			param:		map[string]interface{}{"users":[]string{"id","-name"}}
			return:		" ORDER BY id ASC,name DESC"
*/
func (engine *Engine) SF(sqlstr string, args ...interface{}) *Session {
	session := engine.NewSession()
	session.isAutoClose = true
	return session.SF(sqlstr, args...)
}
func (this *Session) sf(sqlstr string, args ...interface{}) (string, map[string]interface{}) {
	param := map[string]interface{}{}
	pp := make([]interface{}, 0)
	for _, p := range args {
		ty := reflect.TypeOf(p)
		vl := reflect.ValueOf(p)
		if ty.Kind() == reflect.Ptr {
			ty = ty.Elem()
			vl = vl.Elem()
		}
		if ty.Kind() == reflect.Map {
			for _, key := range vl.MapKeys() {
				param[key.String()] = vl.MapIndex(key).Interface()
			}
		} else {
			pp = append(pp, p)
		}
	}
	return fmt.Sprintf(sqlstr, pp...), param
}
func (this *Session) SF(sqlstr string, args ...interface{}) *Session {

	param := map[string]interface{}{}
	if len(args) > 0 {
		//param = args[0].(map[string]interface{})
		sqlstr, param = this.sf(sqlstr, args...)
	}
	sqltemplateKey := fmt.Sprintf("__%x__", md5.Sum([]byte(sqlstr)))
	sqlobj, err := sqltemplateCatch.Get(sqltemplateKey)
	if err != nil { // 如果库里没有操作过这里SQL语句，那么格式化sql语句
		sqlstr = this.engine._sqlFormatAction(sqlstr, param)
		sqlstr = this.engine._sqlFormatOperate(sqlstr, param)
	} else {
		sqlstr = sqlobj.(string)
	}
	_ = sqltemplateCatch.Put(sqltemplateKey, sqlstr)
	sqlstr, params := this.engine._sqlFormatParams(sqlstr, param)
	return this.SQL(sqlstr, params...)
}

func SqlBlockTemplateString(name, sql string) {
	_ = sqltemplateCatch.Put(fmt.Sprintf("_sqlblock_%s", name), sql)
}
