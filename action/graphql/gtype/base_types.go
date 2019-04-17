package gtype

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strconv"
)

var Int = graphql.Int
var String = graphql.String
var Boolean = graphql.Boolean
var Float = graphql.Float
var Int64 = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Int64",
	Description: "`Int64 `标量类型表示非小数有符号整数值。 Int可以表示 - （2 ^ 63）和2 ^ 63 - 1之间的值。",
	Serialize:   parseInt64,
	ParseValue:  parseInt64,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.ParseInt(valueAST.Value, 64, 10); err == nil {
				return intValue
			}
		}
		return nil
	},
})
