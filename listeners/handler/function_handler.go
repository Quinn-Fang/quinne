package handler

import (
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/uspace"
)

func FunctionHandler(operandContext *parser.PrimaryExprContext, argumentsContext *parser.ArgumentsContext) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionName)

	PrimaryExprContextHandler(operandContext)

	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionArgs)

	ArgumentsContextHandler(argumentsContext)

	curCursor.SetCursorContext(sym_tables.ContextTypeDefault)

	curSymTable := sym_tables.GetCurSymTable()
	fFunction := curSymTable.GetLastFunction()
	curNavigator := navigator.GetCurNavigator()
	if !curCursor.IsAppendingExpr() {
		// regular function
		curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)
	} else {
		curCursor.PushExpr(fFunction.ToString())
		curCursor.AddExprVarNames(fFunction.FName)
	}

	// curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)

	return nil
}
