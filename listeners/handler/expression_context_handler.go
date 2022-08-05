package handler

import (
	"fmt"
	"strconv"

	"quinn007.com/listeners/utils"
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/procedures"
	"quinn007.com/sym_tables"
	"quinn007.com/variables"
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

	for _, child := range children {
		fmt.Println("&&&&&&&&&&&&&&&")
		fmt.Printf("%T\n", child)
		fmt.Printf("%+v\n", child)

		switch parserContext := child.(type) {
		case *parser.PrimaryExprContext:
			{
				PrimaryExprContextHandler(parserContext)
			}
		}
	}

	return nil
}

func PrimaryExprContextHandler(contextParser *parser.PrimaryExprContext) error {
	children := contextParser.GetChildren()

	if utils.IsFunction(children) {
		// curStatement := curCursor.GetStatement()
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
	if curCursor.GetCursorContext() == navigator.ContextTypeFunctionName {
		fmt.Println("Gettting Function Name: ", terminalString)
		curSymTable := sym_tables.GetCurSymTable()

		var newValue interface{}
		newReturnValue := variables.NewVariable(
			"",
			variables.VTypeFunctionReturned,
			newValue,
			curCursor.GetIndex())

		curCursor.IncreaseIndex()

		curStatement := curCursor.GetStatement()
		curStatement.AddRightValue(newReturnValue)

		newFunction := procedures.NewFunction(terminalString)
		newFunction.SetReturnValue(newReturnValue)
		curSymTable.AddFunction(newFunction)
	}

	//for _, child := range children {
	//	fmt.Println("7")
	//	fmt.Printf("%T\n", child)
	//	fmt.Printf("%+v\n", child)
	//	fmt.Println("7")
	//}

	return nil
}

func OperandContextHandler(contextParser *parser.OperandContext) error {
	fmt.Println("123456***")
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
		//fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^")
		//fmt.Printf("%T\n", child)
		//fmt.Printf("%+v\n", child)
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
		fmt.Println("************************")
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
		if cursor.GetCursorContext() == navigator.ContextTypeFunctionArgs {
			curFunction := curSymTable.GetLastFunction()
			curFunction.AddParam(curVariable)

			// curStatement.AddRightValue(curVariable)
			// cursor.PrintStatement()
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
		// curStatement := cursor.GetStatement()
		curVariable := variables.NewVariable(
			"",
			variables.VTypeString,
			terminalString,
			cursor.GetIndex())
		curSymTable := sym_tables.GetCurSymTable()

		if cursor.GetCursorContext() == navigator.ContextTypeFunctionArgs {
			curFunction := curSymTable.GetLastFunction()
			curFunction.AddParam(curVariable)

			//cursor.IncreaseIndex()
			//curStatement.AddRightValue(curVariable)
			fmt.Println("++++++++++++++ String", curVariable)
			//cursor.PrintStatement()
		} else {
			curStatement.AddRightValue(curVariable)
		}

		fmt.Println("************************")
	}

	return nil
}
