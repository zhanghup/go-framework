package gtype

import (
	"math"
	"strconv"
)

func parseInt64(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return int64(1)
		}
		return int64(0)
	case *bool:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	case int:
		return int64(value)
	case *int:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	case int8:
		return int64(value)
	case *int8:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int16:
		return int64(value)
	case *int16:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int32:
		return int64(value)
	case *int32:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int64:
		return value
	case *int64:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	case uint:
		if value > math.MaxInt64 {
			return nil
		}
		return int64(value)
	case *uint:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	case uint8:
		return int64(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return int64(*value)
	case uint16:
		return int64(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return int64(*value)
	case uint32:
		if value > uint32(math.MaxInt32) {
			return nil
		}
		return int64(value)
	case *uint32:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	case uint64:
		if value > uint64(math.MaxInt64) {
			return nil
		}
		return int64(value)
	case *uint64:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	case float32:
		return int64(value)
	case *float32:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	case float64:
		if value < float64(math.MinInt64) || value > float64(math.MaxInt64) {
			return nil
		}
		return int64(value)
	case *float64:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return parseInt64(val)
	case *string:
		if value == nil {
			return nil
		}
		return parseInt64(*value)
	}

	return nil
}
