package utils

import (
	"github.com/Quinn-Fang/quinne/variables"
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

func StrToVType(vTypeString string) variables.VTypeEnum {
	switch vTypeString {
	case "int":
		{
			return variables.VTypeInt
		}
	case "string":
		{
			return variables.VTypeString
		}
	default:
		{
			panic("Unrecognized variable type!")
		}
	}
}
