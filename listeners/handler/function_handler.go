package handler

import (
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
)

func FunctionHandler(operandContext *parser.PrimaryExprContext, argumentsContext *parser.ArgumentsContext) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionName)

	PrimaryExprContextHandler(operandContext)

	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionArgs)

	ArgumentsContextHandler(argumentsContext)

	curCursor.SetCursorContext(sym_tables.ContextTypeDefault)
	// add new event to user space queue
	curNavigator := navigator.GetCurNavigator()
	event := navigator.NewEvent(navigator.EventTypeFunction)
	curNavigator.AddNewEvent(event)
	return nil
}
