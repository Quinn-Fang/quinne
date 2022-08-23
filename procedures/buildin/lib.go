package buildin

import (
	"fmt"
	"reflect"
)

// type BuildInGetLength func(obj interface{}) int

// Build in functions
func GetLength(obj interface{}) int {
	reflectType := reflect.TypeOf(obj)
	switch reflectType.Kind() {
	case reflect.Slice:
		{
			return reflect.ValueOf(obj).Len()
		}
	case reflect.Array:
		{
			return reflect.ValueOf(obj).Len()
		}
	case reflect.String:
		{
			return reflect.ValueOf(obj).Len()
		}
	case reflect.Map:
		{
			return reflect.ValueOf(obj).Len()
		}
	default:
		errMsg := fmt.Sprintf("%+v %+v has no length.", obj, reflectType.Kind())
		panic(errMsg)
	}
}
