package procedures

import (
	"fmt"

	"github.com/Quinn-Fang/quinne/variables"
)

//type FFunctionDecl struct {
//	fName   string
//	fParams []*variables.Variable
//	fReturn []variables.VTypeEnum
//}

// Lambda functions

// lambdaExpression, return value + if else stmt
type LambdaExpression struct {
	// string for now, mostly for expr engine
	lReturn string
	// make it a pointer so it's easier to check wether succeeding
	// if condition exists
	lIfCond *variables.Variable
	// make it a pointer so to check wether succeeding Lambda Expression
	// exists.
	// else stmt exists if LambdaExpression is not Nil
	lNextExpression *LambdaExpression
}

func (this *LambdaExpression) GetIfCond() *variables.Variable {
	return this.lIfCond
}

func (this *LambdaExpression) GetIfCondStr() string {
	return this.lIfCond.GetVariableValue().(string)
}

func (this *LambdaExpression) GetNextExpression() *LambdaExpression {
	return this.lNextExpression
}

func (this *LambdaExpression) GetReturnValueString() string {
	return this.lReturn
}

// Lambda declaration
type LambdaDecl struct {
	lExpr        string
	lTernaryExpr string
	// specify parameters name and type
	// without value
	lParams          []*variables.Variable
	lFirstExpression *LambdaExpression
	lLastExpression  *LambdaExpression
}

func (this *LambdaDecl) AddParam(param *variables.Variable) {
	this.lParams = append(this.lParams, param)
}

// Actual lambda function
type LambdaCall struct {
	lDecl   *LambdaDecl
	lArgs   []*variables.Variable
	lReturn *variables.Variable
}

func NewLambdaCall(lambdaDecl *LambdaDecl) *LambdaCall {
	return &LambdaCall{
		lDecl: lambdaDecl,
	}
}

func (this *LambdaCall) GetArgs() []*variables.Variable {
	return this.lArgs
}

func (this *LambdaCall) GetParams() []*variables.Variable {
	return this.lDecl.lParams
}

func (this *LambdaCall) AddArgs(newArg *variables.Variable) {
	this.lArgs = append(this.lArgs, newArg)
}

func (this *LambdaCall) GetLambdaExpr() string {
	return this.lDecl.lExpr
}

func (this *LambdaCall) GetLambdaTernaryExpr() string {
	return this.lDecl.lTernaryExpr
}

func (this *LambdaCall) GetReturnValue() *variables.Variable {
	return this.lReturn
}

func (this *LambdaCall) SetReturnValue(newVariable *variables.Variable) {
	this.lReturn = newVariable
}

func NewLambdaDecl() *LambdaDecl {
	newLambdaExpression := &LambdaExpression{}
	newLambdaDecl := &LambdaDecl{
		lFirstExpression: newLambdaExpression,
		lLastExpression:  newLambdaExpression,
	}
	return newLambdaDecl
}

func (this *LambdaDecl) NewLambdaExpression() *LambdaExpression {
	newLambdaExpression := &LambdaExpression{}
	this.lLastExpression.lNextExpression = newLambdaExpression
	this.lLastExpression = newLambdaExpression
	return newLambdaExpression
}

func (this *LambdaDecl) GetFirstLambdaExpression() *LambdaExpression {
	return this.lFirstExpression
}

//func (this *LambdaDecl) SetRet(retValue string) {
//	this.lLastExpression.lReturn = retValue
//}

func (this *LambdaDecl) AppendRet(retValue string) {
	this.lLastExpression.lReturn += retValue
}

func (this *LambdaDecl) SetIfCond(ifCond string) {
	newStringVariable := variables.NewVariable("", variables.VTypeString, ifCond, -1)
	this.lLastExpression.lIfCond = newStringVariable
}

func (this *LambdaDecl) AppendExpr(exprSubString string) {
	this.lExpr += exprSubString
}

//func (this *LambdaDecl) AddParam(param *variables.Variable) {
//	this.lParams = append(this.lParams, param)
//}

func (this *LambdaDecl) SetTernaryExpr(ternaryExpr string) {
	this.lTernaryExpr = ternaryExpr
}

func (this *LambdaDecl) GetTernaryExpr() string {
	return this.lTernaryExpr
}

type FDeclType int

const (
	FDeclTypeDefault FDeclType = 1
	FDeclTypeLib               = 2
)

type FFunctionDecl struct {
	fName   string
	fType   FDeclType
	fParams []*variables.Variable
	fReturn *variables.Variable
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

func (this *FFunctionDecl) SetFType(fType FDeclType) {
	this.fType = fType
}

func (this *FFunctionDecl) GetFType() FDeclType {
	return this.fType
}

func (this *FFunctionDecl) GetFParams() []*variables.Variable {
	return this.fParams
}

func (this *FFunctionDecl) AddFParams(variable *variables.Variable) {
	this.fParams = append(this.fParams, variable)
}

func (this *FFunctionDecl) SetFReturn(variable *variables.Variable) {
	this.fReturn = variable
}

func (this *FFunctionDecl) GetFReturn() *variables.Variable {
	return this.fReturn
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

func (this *FFunction) ToString() string {
	ret := this.FName
	ret += "("
	for _, variable := range this.FParams {
		vValue := variable.GetVariableValue()
		if _, ok := vValue.(string); ok {
			ret += fmt.Sprintf("%q", vValue)
		} else {
			ret += fmt.Sprintf("%v", vValue)
		}
	}
	ret += ")"

	return ret
}
