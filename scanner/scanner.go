package scanner

import (
	"github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/variables"
)

const (
	errorOCUnknownTypeMsg = "Unknown outer context type"
	errorMCUnknownTypeMsg = "Unknown middle context type"
	errorICUnknownTypeMsg = "Unknown inner context type"
)

// Outer: if else, funcDecl ...
type OuterContext struct {
	contextType consts.OCType
	context     interface{}
}

// expr:
type MiddleContext struct {
	contextType consts.MCType
	context     interface{}
}

// Inner: funcName, funcArgs
type InnerContext struct {
	contextType consts.ICType
	context     interface{}
}

type Scanner struct {
	lineNum       int
	outerContext  *OuterContext
	middleContext *MiddleContext
	innerContext  *InnerContext
}

func NewScanner() *Scanner {
	newScanner := &Scanner{}
	return newScanner
}

func (this *Scanner) ClearAllContext() {
	var (
		nilOuterContext  *OuterContext
		nilMiddleContext *MiddleContext
		nilInnerContext  *InnerContext
	)

	this.outerContext = nilOuterContext
	this.middleContext = nilMiddleContext
	this.innerContext = nilInnerContext
}

func (this *Scanner) NewOuterContext(ocType consts.OCType) {
	newOuterContext := &OuterContext{}
	switch ocType {
	case consts.OCTypeIf, consts.OCTypeElseIf, consts.OCTypeElse:
		{
			newOuterContext.contextType = ocType
			newOuterContext.context = NewIfElseContext()
			this.outerContext = newOuterContext
		}
	default:
		{
			panic(errorOCUnknownTypeMsg)
		}
	}
}

func (this *Scanner) SetOuterType(ocType consts.OCType) {
	if this.outerContext == nil {
		panic("OuterContext not set !")
	}

	outerContext := this.outerContext
	switch ocType {
	case consts.OCTypeIf, consts.OCTypeElseIf, consts.OCTypeElse:
		{
			outerContext.contextType = ocType
		}
	default:
		{
			panic(errorOCUnknownTypeMsg)
		}
	}
}

func (this *Scanner) GetOuterType() consts.OCType {
	if this.outerContext == nil {
		return consts.OCTypeUnSet
	}
	return this.outerContext.contextType
}

func (this *Scanner) NewMiddleContext(mcType consts.MCType) {
	newMiddleContext := &MiddleContext{
		contextType: mcType,
	}
	switch mcType {
	case consts.MCTypeExpr:
		{
			newMiddleContext.context = NewExprContext()
			this.middleContext = newMiddleContext
		}
	default:
		{
			panic(errorMCUnknownTypeMsg)
		}
	}
}

func (this *Scanner) SetMiddleType(mcType consts.MCType) {
	if this.middleContext == nil {
		panic("middleContext not set !")
	}
	switch mcType {
	case consts.MCTypeExpr:
		{
			this.middleContext.contextType = mcType
		}
	default:
		{
			panic(errorMCUnknownTypeMsg)
		}
	}
}

func (this *Scanner) GetMiddleType() consts.MCType {
	if this.middleContext == nil {
		return consts.MCTypeUnset
	}
	return this.middleContext.contextType
}

func (this *Scanner) isOuterIfElse() bool {
	if this.outerContext.contextType == consts.OCTypeIf ||
		this.outerContext.contextType == consts.OCTypeElseIf ||
		this.outerContext.contextType == consts.OCTypeElse {
		return true
	}
	return false
}

func (this *Scanner) isMiddleExpr() bool {
	if this.middleContext.contextType == consts.MCTypeExpr {
		return true
	}
	return false
}

// add if or else key words at parsing phase
func (this *Scanner) AddIfElseEvent(logicSymbol sym_tables.LogicSymbol) {
	if !this.isOuterIfElse() {
		panic("not in if else outer context !")
	}

	ifElseContext, _ := this.outerContext.context.(*IfElseContext)
	ifElseContext.queue.PushBack(logicSymbol)

}

func (this *Scanner) AppendExpr(subExpr string) {
	if !this.isMiddleExpr() {
		panic("not in expr middle context !")
	}

	exprContext, _ := this.middleContext.context.(*ExprContext)
	exprContext.AppendExpr(subExpr)
}

func (this *Scanner) AppendExprVarName(exprVarName string) {
	if !this.isMiddleExpr() {
		panic("not in expr middle context !")
	}

	exprContext, _ := this.middleContext.context.(*ExprContext)
	exprContext.AppendVarName(exprVarName)
}

// called before entering if else branch block
func (this *Scanner) ConsumeIfElseEvent() (sym_tables.ContextType, *sym_tables.IfElseBranch) {
	if !this.isOuterIfElse() && !this.isMiddleExpr() {
		panic("not in if else outer context and expr middle context !")
	}

	ifElseContext, _ := this.outerContext.context.(*IfElseContext)
	exprContext, _ := this.middleContext.context.(*ExprContext)
	branchType, ifElseBranch := ifElseContext.GetCurrentBranch(exprContext.exprString, exprContext.exprVarNames)
	var contextType sym_tables.ContextType
	switch branchType {
	case consts.OCTypeIf:
		{
			contextType = sym_tables.ContextTypeIf
		}
	case consts.OCTypeElseIf:
		{
			contextType = sym_tables.ContextTypeElseIf
		}
	case consts.OCTypeElse:
		{
			contextType = sym_tables.ContextTypeElse
		}
	default:
		{
			panic("Unknown if else context type !")
		}
	}

	return contextType, ifElseBranch
}

func (this *Scanner) NewInnerContext(icType consts.ICType) {
	newInnerContext := &InnerContext{
		contextType: icType,
	}
	switch icType {
	case consts.ICTypeFuncName:
		{
			newInnerContext.context = NewFunctionContext()
			this.innerContext = newInnerContext
		}
	case consts.ICTypeLambdaParams, consts.ICTypeLambdaExpr:
		{
			newInnerContext.context = NewLambdaContext()
			this.innerContext = newInnerContext
		}
	case consts.ICTypeLambdaCall:
		{
			newInnerContext.context = NewLambdaCallContext()
			this.innerContext = newInnerContext
		}
	default:
		{
			panic(errorICUnknownTypeMsg)
		}
	}
}

func (this *Scanner) SetInnerType(icType consts.ICType) {
	if this.innerContext == nil {
		panic("innerContext not set !")
	}
	this.innerContext.contextType = icType
	//switch icType {
	//case consts.ICTypeFuncName, consts.ICTypeFuncArgs, consts.ICTypeUnset:
	//	{
	//		this.innerContext.contextType = icType
	//	}
	//case consts.ICTypeLambdaParams, consts.ICTypeLambdaExpr, consts.ICTypeLambdaIfClause,
	//	consts.ICTypeLambdaIfExpr, consts.ICTypeLambdaRet:
	//	{
	//		this.innerContext.contextType = icType
	//	}
	//default:
	//	{
	//		panic(errorICUnknownTypeMsg)
	//	}
	//}
}

func (this *Scanner) GetInnerType() consts.ICType {
	if this.innerContext == nil {
		return consts.ICTypeUnset
	}
	return this.innerContext.contextType
}

// Lambda context

func (this *Scanner) AddLambdaParam(paramName string) {
	if this.innerContext == nil {
		panic("innerContext not set !")
	}
	if !(this.innerContext.contextType == consts.ICTypeLambdaParams) {
		panic("not lambda inner context!")
	}
	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
	lambdaContext.AddParam(paramName)
}

func (this *Scanner) AppendLambdaExpr(exprStrRaw string) {
	if this.innerContext == nil {
		panic("innerContext not set !")
	}
	if !(this.innerContext.contextType == consts.ICTypeLambdaExpr) {
		panic("not lambda inner context!")
	}
	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
	lambdaContext.AppendExprRaw(exprStrRaw)
}

func (this *Scanner) AppendLambdaExprList(exprStr string) {
	if this.innerContext == nil {
		panic("innerContext not set !")
	}
	//if !(this.innerContext.contextType == consts.ICTypeLambdaIfClause) {
	//	panic("not lambda inner context!")
	//}
	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
	lambdaContext.AppendExprList(exprStr)
}

func (this *Scanner) SetLambdaReturnValue(retValue string) {
	if this.innerContext == nil {
		panic("innerContext not set !")
	}
	//if !(this.innerContext.contextType == consts.ICTypeLambdaIfClause) {
	//	panic("not lambda inner context!")
	//}
	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
	lambdaContext.SetLReturn(retValue)
}

func (this *Scanner) AppendLambdaReturnValue(retValue string) {
	if this.innerContext == nil {
		panic("innerContext not set !")
	}
	//if !(this.innerContext.contextType == consts.ICTypeLambdaIfClause) {
	//	panic("not lambda inner context!")
	//}
	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
	lambdaContext.AppendLReturn(retValue)
}

func (this *Scanner) AddLambdaParamToDecl(vType variables.VTypeEnum) {
	if this.innerContext == nil {
		panic("innerContext not set !")
	}
	if !(this.innerContext.contextType == consts.ICTypeLambdaParams) {
		panic("not lambda inner context!")
	}
	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
	lambdaContext.AddLambdaDeclParams(vType)
}

//func (this *Scanner) NewLambdaIfElseClauseContext() *LambdaIfElseContext {
//	return NewLambdaIfElseContext()
//}

//func (this *Scanner) SetLambdaIfElseClauseContext(lIfElseContext *LambdaIfElseContext) {
//	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
//	lambdaContext.lIfElseClauseCtx = lIfElseContext
//}
//
//func (this *Scanner) GetLambdaIfElseClause() *LambdaIfElseContext {
//	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
//	return lambdaContext.lIfElseClauseCtx
//}

//func (this *Scanner) SetLambdaIfElseClauseContextEntry(lIfElseContext *LambdaIfElseContext) {
//	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
//	lambdaContext.lIfElseClauseCtxEntry = lIfElseContext
//}
//
//func (this *Scanner) GetLambdaIfElseClauseEntry() *LambdaIfElseContext {
//	lambdaContext, _ := this.innerContext.context.(*LambdaContext)
//	return lambdaContext.lIfElseClauseCtxEntry
//}

func (this *Scanner) GetInnerContext() interface{} {
	return this.innerContext.context
}

func (this *Scanner) GetLambdaContext() *LambdaContext {
	if lambdaContext, ok := this.innerContext.context.(*LambdaContext); ok {
		return lambdaContext
	} else {
		panic("innerContext is not LambdaContext")
	}
}
