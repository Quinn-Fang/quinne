package handler

import (
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
	"quinn007.com/uspace"
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
	curNavigator.AddEvent(uspace.EventTypeFunction, fFunction)
	// add new event to user space queue

	// curNavigator := navigator.GetCurNavigator()

	//event := navigator.NewEvent(navigator.EventTypeFunction)
	//curNavigator.AddNewEvent(event)
	return nil
}
