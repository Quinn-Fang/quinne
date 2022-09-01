package scanner

import (
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/uspace"
	"github.com/Quinn-Fang/quinne/utils"
)

type IfElseContext struct {
	queue utils.Queue
}

func NewIfElseContext() *IfElseContext {
	newIfElseContext := &IfElseContext{
		queue: *utils.NewQueue(),
	}
	return newIfElseContext
}

func (this *IfElseContext) toLogicSymbol(elem interface{}) sym_tables.LogicSymbol {
	if ret, ok := elem.(sym_tables.LogicSymbol); ok {
		return ret
	} else {
		panic("can not get int value of LogicSymbol")
	}
}

// check the internal queue and return currrent branch
// context, as if, else-if, else
func (this *IfElseContext) GetCurrentBranch(exprString string, exprVarNames []string) (consts.OCType, *sym_tables.IfElseBranch) {
	// no if else event
	if this.queue.IsEmpty() {
		panic("if else queue Empty !")
	}

	curSymTable := sym_tables.GetCurSymTable()
	curNavigator := navigator.GetCurNavigator()

	ifElseClause := curSymTable.GetLastIfElseClause()
	firstSymbol, _ := this.queue.PopBack()

	var curEventType consts.OCType
	var curEventContext *sym_tables.IfElseBranch

	if this.toLogicSymbol(firstSymbol) == sym_tables.LogicSymbolIf {

		// if curSymTable.IfElseStackEmpty() {
		if this.queue.IsEmpty() {
			// if event
			newIfBranch := sym_tables.NewIfElseBranch(sym_tables.BranchTypeIf)
			newIfBranch.SetExpr(exprString)
			newIfBranch.SetExprVarNames(exprVarNames)
			newIfBranch.SetParent(ifElseClause)

			ifElseClause.AddBranch(newIfBranch)
			// add user space event
			curNavigator.AddEvent(uspace.EventTypeIfElseExpr, newIfBranch, curSymTable)

			curEventType = consts.OCTypeIf
			curEventContext = newIfBranch

		} else {
			// else-if event
			// secondSymbol := curSymTable.PopBackIfElseStack()
			secondSymbol, _ := this.queue.PopBack()

			if this.toLogicSymbol(secondSymbol) != sym_tables.LogicSymbolElse {
				panic("else not followed by if")
			}
			newElseIfBranch := sym_tables.NewIfElseBranch(sym_tables.BranchTypeElseIf)
			newElseIfBranch.SetExpr(exprString)
			newElseIfBranch.SetExprVarNames(exprVarNames)
			newElseIfBranch.SetParent(ifElseClause)

			ifElseClause.AddBranch(newElseIfBranch)
			// add user space event
			curNavigator.AddEvent(uspace.EventTypeIfElseExpr, newElseIfBranch, curSymTable)

			curEventType = consts.OCTypeElseIf
			curEventContext = newElseIfBranch
		}
	} else if this.toLogicSymbol(firstSymbol) == sym_tables.LogicSymbolElse {
		// else event
		newElseBranch := sym_tables.NewIfElseBranch(sym_tables.BranchTypeElse)
		newElseBranch.SetParent(ifElseClause)

		ifElseClause.AddBranch(newElseBranch)

		curEventType = consts.OCTypeElse
		curEventContext = newElseBranch

	} else {
		panic("Unknown error")
	}

	return curEventType, curEventContext
}
