package buildin

import (
	"fmt"
	"reflect"

	"github.com/Quinn-Fang/quinne/procedures"
)

var (
	curFuncTable *FuncTable
)

type FuncTable struct {
	functionMap map[string]*procedures.FFunction
}

func NewFuncTable() *FuncTable {
	newFuncTable := &FuncTable{
		functionMap: make(map[string]*procedures.FFunction),
	}
	return newFuncTable
}

func SetFuncTable(funcTable *FuncTable) {
	curFuncTable = funcTable
}

func InitFuncTable() {
	newFuncTable := NewFuncTable()
	SetFuncTable(newFuncTable)
}

func (this *FuncTable) AddFunction(fFunction *procedures.FFunction) {
	if _, ok := this.functionMap[fFunction.GetFunctionName()]; ok {
		panic("Function already exists.")
	} else {
		this.functionMap[fFunction.GetFunctionName()] = fFunction
	}
}

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
