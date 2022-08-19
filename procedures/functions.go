package procedures

import (
	"quinn007.com/variables"
)

type FFunctionDecl struct {
	fName   string
	fParams []*variables.Variable
	fReturn []variables.VTypeEnum
}

func NewFunctionDecl(functionName string) *FFunctionDecl {
	newFunctionDecl := &FFunctionDecl{
		fName: functionName,
	}
	return newFunctionDecl
}

func (this *FFunctionDecl) GetFunctionName() string {
	return this.fName
}

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

func (this *FFunction) GetFunctionName() string {
	return this.FName
}

func (this *FFunction) GetParams() []*variables.Variable {
	return this.FParams
}

func (this *FFunction) AddParam(param *variables.Variable) {
	this.FParams = append(this.FParams, param)
}

func (this *FFunction) GetReturnValue() interface{} {
	return this.FReturn.GetVariableValue()
}

func (this *FFunction) SetReturnValue(returnValue interface{}) {
	fReturnVar := this.FReturn
	fReturnVar.SetVariableValue(returnValue)
}

func (this *FFunction) InitReturnValue(variable *variables.Variable) {
	this.FReturn = variable
}
