package handler

import (
	"fmt"
	"strconv"

	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/variables"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func ExpressionStmtContextHandler(contextParser *parser.ExpressionStmtContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext)
			}
		}
	}

	return nil
}

func ExpressionContextHandler(contextParser *parser.ExpressionContext) error {
	children := contextParser.GetChildren()

	curCursor, _ := navigator.GetCursor()
	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.PrimaryExprContext:
			{
				PrimaryExprContextHandler(parserContext)
			}
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext)
			}
		case *antlr.TerminalNodeImpl:
			{
				// if curCursor.GetCursorContext() == sym_tables.ContextTypeIf || curCursor.GetCursorContext() == sym_tables.ContextTypeElseIf {
				if curCursor.IsAppendingExpr() {
					terminalString, _ := utils.GetTerminalNodeText(parserContext)

					curCursor.PushExpr(terminalString)
				}

				//terminalString, _ := utils.GetTerminalNodeText(parserContext)
				//fmt.Println(terminalString)
			}
		}
	}

	return nil
}

func PrimaryExprContextHandler(contextParser *parser.PrimaryExprContext) error {
	children := contextParser.GetChildren()

	if utils.IsFunction(children) {
		FunctionHandler(children[0].(*parser.PrimaryExprContext), children[1].(*parser.ArgumentsContext))
		return nil
	}

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.IdentifierListContext:
			{
				IdentifierListContextHandler(parserContext)
			}
		case *parser.LiteralContext:
			{
				LiteralContextHandler(parserContext)
			}
		case *parser.PrimaryExprContext:
			{
				PrimaryExprContextHandler(parserContext)
			}
		case *parser.OperandContext:
			{
				OperandContextHandler(parserContext)
			}
		}
	}

	return nil
}

func OperandNameContextHandler(contextParser *parser.OperandNameContext) error {
	children := contextParser.GetChildren()
	terminalString, _ := utils.GetTerminalNodeText(children[0])
	curCursor, _ := navigator.GetCursor()
	curSymTable := sym_tables.GetCurSymTable()
	curStatement := curCursor.GetStatement()

	if curCursor.GetCursorContext() == sym_tables.ContextTypeFunctionName {

		var emptyValue interface{}
		newReturnValue := variables.NewVariable(
			"",
			variables.VTypeFunctionReturned,
			emptyValue,
			curCursor.GetIndex())

		curCursor.IncreaseIndex()

		curStatement.AddRightValue(newReturnValue)

		newFunction := procedures.NewFunction(terminalString)
		// newFunction.SetReturnValue(newReturnValue)
		newFunction.InitReturnValue(newReturnValue)
		curSymTable.AddFunction(newFunction)
	} else if curCursor.GetCursorContext() == sym_tables.ContextTypeFunctionArgs {
		variable, err := curSymTable.GetVariableByName(terminalString)
		if err != nil {
			errMsg := fmt.Sprintf("variable: %s does not exist", terminalString)
			panic(errMsg)
		}

		curFunction := curSymTable.GetLastFunction()
		curFunction.AddParam(variable)
		// } else if curCursor.GetCursorContext() == sym_tables.ContextTypeIf || curCursor.GetCursorContext() == sym_tables.ContextTypeElseIf {
	} else if curCursor.IsAppendingExpr() {
		curCursor.PushExpr(terminalString)
		curCursor.AddExprVarNames(terminalString)
	}

	return nil
}

func OperandContextHandler(contextParser *parser.OperandContext) error {
	//	fmt.Println("123456***")
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.IdentifierListContext:
			{
				IdentifierListContextHandler(parserContext)
			}
		case *parser.LiteralContext:
			{
				LiteralContextHandler(parserContext)
			}
		case *parser.OperandNameContext:
			{
				OperandNameContextHandler(parserContext)
			}

		}
	}

	return nil
}

func LiteralContextHandler(contextParser *parser.LiteralContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.BasicLitContext:
			{
				BasicLitContextHandler(parserContext)
			}

		}
	}

	return nil
}

func BasicLitContextHandler(contextParser *parser.BasicLitContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.String_Context:
			{
				StringContextHandler(parserContext)
			}
		case *parser.IntegerContext:
			{
				IntegerContextHandler(parserContext)
			}
		}
	}

	return nil
}

func IntegerContextHandler(contextParser *parser.IntegerContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		cursor, _ := navigator.GetCursor()
		terminalString, _ := utils.GetTerminalNodeText(child)
		curStatement := cursor.GetStatement()
		intVal, _ := strconv.Atoi(terminalString)
		curVariable := variables.NewVariable(
			"",
			variables.VTypeInt,
			intVal,
			cursor.GetIndex())

		//cursor.IncreaseIndex()
		curSymTable := sym_tables.GetCurSymTable()
		if cursor.GetCursorContext() == sym_tables.ContextTypeFunctionArgs {
			curFunction := curSymTable.GetLastFunction()
			curFunction.AddParam(curVariable)

			// curStatement.AddRightValue(curVariable)
			// cursor.PrintStatement()
			// } else if cursor.GetCursorContext() == sym_tables.ContextTypeIf || cursor.GetCursorContext() == sym_tables.ContextTypeElseIf {
		} else if cursor.IsAppendingExpr() {
			cursor.PushExpr(terminalString)
		} else {
			curStatement.AddRightValue(curVariable)

		}
	}

	return nil
}

func StringContextHandler(contextParser *parser.String_Context) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		cursor, _ := navigator.GetCursor()
		curStatement := cursor.GetStatement()
		terminalString, _ := utils.GetTerminalNodeText(child)
		// strip out quotes
		terminalString = terminalString[1 : len(terminalString)-1]
		curVariable := variables.NewVariable(
			"",
			variables.VTypeString,
			terminalString,
			cursor.GetIndex())
		curSymTable := sym_tables.GetCurSymTable()

		if cursor.GetCursorContext() == sym_tables.ContextTypeFunctionArgs {
			curFunction := curSymTable.GetLastFunction()
			curFunction.AddParam(curVariable)

			// } else if cursor.GetCursorContext() == sym_tables.ContextTypeIf || cursor.GetCursorContext() == sym_tables.ContextTypeElseIf {
		} else if cursor.IsAppendingExpr() {
			cursor.PushExpr(terminalString)
		} else {
			curStatement.AddRightValue(curVariable)
		}

	}

	return nil
}
