package gql

import "github.com/graphql-go/graphql"

var Int = newScalar(graphql.Int)
var String = newScalar(graphql.String)
var Boolean = newScalar(graphql.Boolean)
var Float = newScalar(graphql.Float)
//var Int64 = newScalar(graphql.NewScalar(graphql.ScalarConfig{
//	Name:        "Int64",
//	Description: "`Int64 `标量类型表示非小数有符号整数值。 Int可以表示 - （2 ^ 63）和2 ^ 63 - 1之间的值。",
//	Serialize:   parseInt64,
//	ParseValue:  parseInt64,
//	ParseLiteral: func(valueAST ast.Value) interface{} {
//		switch valueAST := valueAST.(type) {
//		case *ast.IntValue:
//			if intValue, err := strconv.ParseInt(valueAST.Value, 64, 10); err == nil {
//				return intValue
//			}
//		}
//		return nil
//	},
//}))
//
//func parseInt64(value interface{}) interface{} {
//	switch value := value.(type) {
//	case bool:
//		if value == true {
//			return int64(1)
//		}
//		return int64(0)
//	case *bool:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	case int:
//		return int64(value)
//	case *int:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	case int8:
//		return int64(value)
//	case *int8:
//		if value == nil {
//			return nil
//		}
//		return int64(*value)
//	case int16:
//		return int64(value)
//	case *int16:
//		if value == nil {
//			return nil
//		}
//		return int64(*value)
//	case int32:
//		return int64(value)
//	case *int32:
//		if value == nil {
//			return nil
//		}
//		return int64(*value)
//	case int64:
//		return value
//	case *int64:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	case uint:
//		if value > math.MaxInt64 {
//			return nil
//		}
//		return int64(value)
//	case *uint:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	case uint8:
//		return int64(value)
//	case *uint8:
//		if value == nil {
//			return nil
//		}
//		return int64(*value)
//	case uint16:
//		return int64(value)
//	case *uint16:
//		if value == nil {
//			return nil
//		}
//		return int64(*value)
//	case uint32:
//		if value > uint32(math.MaxInt32) {
//			return nil
//		}
//		return int64(value)
//	case *uint32:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	case uint64:
//		if value > uint64(math.MaxInt64) {
//			return nil
//		}
//		return int64(value)
//	case *uint64:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	case float32:
//		return int64(value)
//	case *float32:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	case float64:
//		if value < float64(math.MinInt64) || value > float64(math.MaxInt64) {
//			return nil
//		}
//		return int64(value)
//	case *float64:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	case string:
//		val, err := strconv.ParseFloat(value, 0)
//		if err != nil {
//			return nil
//		}
//		return parseInt64(val)
//	case *string:
//		if value == nil {
//			return nil
//		}
//		return parseInt64(*value)
//	}
//
//	return nil
//}
