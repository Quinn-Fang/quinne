package buildin

import (
	"errors"
	"fmt"
)

var (
	curFuncTable *FuncTable
)

type FuncTable struct {
	functionMap map[string]interface{}
}

func NewFuncTable() *FuncTable {
	newFuncTable := &FuncTable{
		functionMap: make(map[string]interface{}),
	}
	return newFuncTable
}

func SetFuncTable(funcTable *FuncTable) {
	curFuncTable = funcTable
}

func InitFuncTable() {
	newFuncTable := NewFuncTable()
	SetFuncTable(newFuncTable)

	newFuncTable.AddFunction("len", GetLength)
}

func GetSystemFuncTable() *FuncTable {
	return curFuncTable
}

func (this *FuncTable) AddFunction(fName string, fFunction interface{}) {
	if _, ok := this.functionMap[fName]; ok {
		panic("Function already exists.")
	} else {
		this.functionMap[fName] = fFunction
	}
}

func (this *FuncTable) GetFunctionByName(fName string) (interface{}, error) {
	if function, ok := this.functionMap[fName]; !ok {
		errMsg := fmt.Sprintf("System function not found %+v \n", fName)
		return nil, errors.New(errMsg)
	} else {
		return function, nil
	}
}
