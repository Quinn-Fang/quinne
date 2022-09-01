package scanner

import (
	"github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
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

func (this *Scanner) SetOuterContext(ocType consts.OCType) {
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
	return this.outerContext.contextType
}

func (this *Scanner) SetMiddleContext(mcType consts.MCType) {
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

func (this *Scanner) GetMiddleType() consts.MCType {
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
