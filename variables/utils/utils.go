package utils

import (
	"github.com/Quinn-Fang/Quinne/variables"
)

func TypeOf(v interface{}) variables.VTypeEnum {
	switch v.(type) {
	case int, int32, int64:
		{
			return variables.VTypeInt
		}
	case string:
		{
			return variables.VTypeString
		}
	case float32, float64:
		{
			return variables.VTypeFloat
		}
	case bool:
		{
			return variables.VTypeBool
		}
	default:
		{
			return variables.VTypeUndefined
		}
	}
}
