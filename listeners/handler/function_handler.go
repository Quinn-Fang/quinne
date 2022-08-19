package handler

import (
	"github.com/Quinn-Fang/Quinne/navigator"
	"github.com/Quinn-Fang/Quinne/parser"
	"github.com/Quinn-Fang/Quinne/sym_tables"
	"github.com/Quinn-Fang/Quinne/uspace"
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
	curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)
	// add new event to user space queue

	// curNavigator := navigator.GetCurNavigator()

	//event := navigator.NewEvent(navigator.EventTypeFunction)
	//curNavigator.AddNewEvent(event)
	return nil
}
