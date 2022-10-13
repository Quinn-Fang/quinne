package variables

import (
	"errors"
	"fmt"
)

type VTypeEnum int

const (
	VTypeUndefined        VTypeEnum = 0
	VTypeInt                        = 1
	VTypeString                     = 2
	VTypeFloat                      = 3
	VTypeBool                       = 4
	VTypeMap                        = 5
	VTypeFunctionReturned           = 6
	VTypeFunctionDecl               = 7
)

func StrToVType(vTypeString string) VTypeEnum {
	switch vTypeString {
	case "int":
		{
			return VTypeInt
		}
	case "string":
		{
			return VTypeString
		}
	default:
		{
			panic("Unrecognized variable type!")
		}
	}
}

type Variable struct {
	vName                  string
	vType                  VTypeEnum
	vValue                 interface{}
	vIndex                 int
	vFunctionReturnedIndex int
}

func NewEmptyVariable() *Variable {
	newVariable := Variable{}

	return &newVariable
}

func NewVariable(variableName string, variableType VTypeEnum, variableValue interface{}, index int) *Variable {
	newVariable := Variable{
		vName:  variableName,
		vType:  variableType,
		vValue: variableValue,
		vIndex: index,
	}

	return &newVariable
}

func (this *Variable) GetInt() (int, error) {
	if this.vType == VTypeInt {
		return this.vValue.(int), nil
	} else {
		return -1, errors.New("Wrong type")
	}
}

func (this *Variable) GetString() (string, error) {
	if this.vType == VTypeString {
		return this.vValue.(string), nil
	} else {
		return "", errors.New("Wrong type")
	}
}

func (this *Variable) GetFloat() (float64, error) {
	if this.vType == VTypeFloat {
		return this.vValue.(float64), nil
	} else {
		return -1.1, errors.New("Wrong type")
	}
}

func (this *Variable) GetBool() (bool, error) {
	if this.vType == VTypeFloat {
		return this.vValue.(bool), nil
	} else {
		return false, errors.New("Wrong type")
	}
}

func (this *Variable) ToString() string {
	if this.vType == VTypeInt || this.vType == VTypeBool || this.vType == VTypeString || this.vType == VTypeFloat {
		return fmt.Sprintf("%v", this.vValue)
	} else {
		errStr := fmt.Sprintf("can not convert %v %v to string", this.vType, this.vValue)
		panic(errStr)
	}
}

func (this *Variable) GetVariableName() string {
	return this.vName
}

func (this *Variable) SetVariableName(variableName string) {
	this.vName = variableName
}

func (this *Variable) SetVariableValue(variableValue interface{}) {
	this.vValue = variableValue
}

func (this *Variable) GetVariableType() VTypeEnum {
	return this.vType
}

func (this *Variable) GetVariableValue() interface{} {
	return this.vValue
}

func (this *Variable) GetVariableIndex() int {
	return this.vIndex
}
