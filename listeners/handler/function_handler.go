package handler

import (
	"quinn007.com/parser"
)

func FunctionHandler(operandContext *parser.PrimaryExprContext, argumentsContext *parser.ArgumentsContext) error {
	PrimaryExprContextHandler(operandContext)
	ArgumentsContextHandler(argumentsContext)
	return nil
}
