package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Quinn-Fang/quinne/variables"
)

const (
	funcReturnNotEqual string = "Func Return Not Equal: %+v != %+v \n"
	paramNotEqual             = "Parameter Not Equal: %+v != %+v \n"
	variablesNotEqual         = "Variables Not Equal: %+v != %+v"
)

func MakeSlice(elems ...interface{}) []interface{} {
	ret := make([]interface{}, 0)
	for _, elem := range elems {
		ret = append(ret, elem)
	}
	return ret
}

func MakeVarSlice(elems ...*variables.Variable) []*variables.Variable {
	ret := make([]*variables.Variable, 0)
	for _, elem := range elems {
		ret = append(ret, elem)
	}
	return ret
}

func CompareVars(t *testing.T, vars []*variables.Variable, expectedValues []interface{}) {
	if len(vars) != len(expectedValues) {
		panic("Number of vars and expected not same !")
	}

	for i := 0; i < len(vars); i++ {
		value := vars[i].GetVariableValue()
		expected := expectedValues[i]
		if !reflect.DeepEqual(value, expected) {
			errMsg := fmt.Sprintf(variablesNotEqual, value, expected)
			t.Error(errMsg)
			panic("Variable comparison wrong !")
		}
	}
	t.Log("Variables compare complete")
}
