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
	InnerContext  *InnerContext
}

func NewScanner() *Scanner {
	newScanner := &Scanner{}
	return newScanner
}

func (this *Scanner) InitOuterContext(ocType consts.OCType) {
	newOuterContext := OuterContext{
		contextType: ocType,
	}
	switch ocType {
	case consts.OCTypeIf, consts.OCTypeElseIf, consts.OCTypeElse:
		{
			newOuterContext.context = NewIfElseContext
			this.InitMiddleContext(consts.MCTypeExpr)
		}
	default:
		{
			panic(errorOCUnknownTypeMsg)
		}
	}
}

func (this *Scanner) InitMiddleContext(mcType consts.MCType) {
	newMiddleContext := MiddleContext{
		contextType: mcType,
	}
	switch mcType {
	case consts.MCTypeExpr:
		{
			newMiddleContext.context = NewExprContext()
		}
	default:
		{
			panic(errorMCUnknownTypeMsg)
		}
	}
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
func (this *Scanner) ConsumeIfElseEvent() (consts.LogicContextType, *sym_tables.IfElseBranch) {
	if !this.isOuterIfElse() && !this.isMiddleExpr() {
		panic("not in if else outer context and expr middle context !")
	}

	ifElseContext, _ := this.outerContext.context.(*IfElseContext)
	exprContext, _ := this.middleContext.context.(*ExprContext)
	return ifElseContext.GetCurrentBranch(exprContext.exprString, exprContext.exprVarNames)

}
