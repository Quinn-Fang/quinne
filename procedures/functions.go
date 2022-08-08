package procedures

import (
	"quinn007.com/variables"
)

type FFunction struct {
	FName    string
	FParams  []*variables.Variable
	FReturn  *variables.Variable
	Assigned bool
}

func NewFunction(functionName string) *FFunction {
	newFunction := &FFunction{
		FName:    functionName,
		FParams:  make([]*variables.Variable, 0),
		Assigned: false,
	}

	return newFunction
}

func (this *FFunction) GetParams() []*variables.Variable {
	return this.FParams
}

func (this *FFunction) AddParam(param *variables.Variable) {
	this.FParams = append(this.FParams, param)
}

func (this *FFunction) GetReturnValue() *variables.Variable {
	return this.FReturn
}

func (this *FFunction) SetReturnValue(returnValue *variables.Variable) {
	this.FReturn = returnValue
}
