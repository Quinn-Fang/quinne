package procedures

import (
	"fmt"

	"github.com/Quinn-Fang/quinne/variables"
)

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

type FunctionCallDecl struct {
	fName   string
	fType   FDeclType
	fParams []*variables.Variable
	fReturn *variables.Variable
}

func NewFunctionDecl(functionName string) *FunctionCallDecl {
	newFunctionDecl := &FunctionCallDecl{
		fName: functionName,
	}
	return newFunctionDecl
}

func (this *FunctionCallDecl) GetFunctionName() string {
	return this.fName
}

func (this *FunctionCallDecl) SetFType(fType FDeclType) {
	this.fType = fType
}

func (this *FunctionCallDecl) GetFType() FDeclType {
	return this.fType
}

func (this *FunctionCallDecl) GetFParams() []*variables.Variable {
	return this.fParams
}

func (this *FunctionCallDecl) AddFParams(variable *variables.Variable) {
	this.fParams = append(this.fParams, variable)
}

func (this *FunctionCallDecl) SetFReturn(variable *variables.Variable) {
	this.fReturn = variable
}

func (this *FunctionCallDecl) GetFReturn() *variables.Variable {
	return this.fReturn
}

type FunctionCall struct {
	FName    string
	FParams  []*variables.Variable
	FReturn  *variables.Variable
	Assigned bool
}

func NewFunction(functionName string) *FunctionCall {
	newFunction := &FunctionCall{
		FName:    functionName,
		FParams:  make([]*variables.Variable, 0),
		Assigned: false,
	}

	return newFunction
}

func (this *FunctionCall) GetFunctionName() string {
	return this.FName
}

func (this *FunctionCall) GetParams() []*variables.Variable {
	return this.FParams
}

func (this *FunctionCall) AddParam(param *variables.Variable) {
	this.FParams = append(this.FParams, param)
}

func (this *FunctionCall) GetReturnValue() interface{} {
	return this.FReturn.GetVariableValue()
}

func (this *FunctionCall) SetReturnValue(returnValue interface{}) {
	fReturnVar := this.FReturn
	fReturnVar.SetVariableValue(returnValue)
}

func (this *FunctionCall) InitReturnValue(variable *variables.Variable) {
	this.FReturn = variable
}

func (this *FunctionCall) ToString() string {
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
