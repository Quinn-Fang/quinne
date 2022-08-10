package handler

import (
	"quinn007.com/navigator"
	"quinn007.com/parser"
)

func FunctionHandler(operandContext *parser.PrimaryExprContext, argumentsContext *parser.ArgumentsContext) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.SetCursorContext(navigator.ContextTypeFunctionName)

	PrimaryExprContextHandler(operandContext)

	curCursor.SetCursorContext(navigator.ContextTypeFunctionArgs)

	ArgumentsContextHandler(argumentsContext)

	curCursor.SetCursorContext(navigator.ContextTypeDefault)
	// add new event to user space queue
	curNavigator := navigator.GetCurNavigator()
	event := navigator.NewEvent(navigator.EventTypeFunction)
	curNavigator.AddNewEvent(event)
	return nil
}
