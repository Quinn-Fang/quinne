package handler

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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

	curCursor, _ := navigator.GetCursor()
	cursorContext := curCursor.GetCursorContext()
	for _, child := range children {
		fmt.Println("10001 ]]]]]]]]]]")
		fmt.Println(cursorContext)
		fmt.Printf("%T\n", child)
		if cursorContext == navigator.ContextTypeIfBlock || cursorContext == navigator.ContextTypeElseBlock {
		}
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
				if curCursor.GetCursorContext() == navigator.ContextTypeIfBlock {
					terminalString, _ := utils.GetTerminalNodeText(parserContext)
					// curExpr := curCursor.GetExpr()
					// curExpr.SetOperator(navigator.OperatorType(terminalString))
					curCursor.PushExpr(terminalString)
				}
				fmt.Println("10002 ]]]]]]]]]]")
				terminalString, _ := utils.GetTerminalNodeText(parserContext)
				fmt.Println(terminalString)
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
		//curCursor, _ := navigator.GetCursor()
		//if curCursor.GetCursorContext()
		//fmt.Println("####################################################")
		//fmt.Println()
		//fmt.Printf("%T\n", child)
		//fmt.Printf("%+v\n", child)
		//fmt.Println()

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

	for _, child := range children {
		fmt.Println("7----7")
		fmt.Printf("%T\n", child)
		fmt.Printf("%+v\n", child)
		fmt.Println("7----7")
	}

	if curCursor.GetCursorContext() == navigator.ContextTypeFunctionName {
		//		fmt.Println("Gettting Function Name: ", terminalString)

		var newValue interface{}
		newReturnValue := variables.NewVariable(
			"",
			variables.VTypeFunctionReturned,
			newValue,
			curCursor.GetIndex())

		curCursor.IncreaseIndex()

		curStatement.AddRightValue(newReturnValue)

		newFunction := procedures.NewFunction(terminalString)
		newFunction.SetReturnValue(newReturnValue)
		curSymTable.AddFunction(newFunction)
	} else if curCursor.GetCursorContext() == navigator.ContextTypeFunctionArgs {
		variable, err := curSymTable.GetVariableByName(terminalString)
		if err != nil {
			errMsg := fmt.Sprintf("variable: %s does not exist", terminalString)
			panic(errMsg)
		}

		curFunction := curSymTable.GetLastFunction()
		curFunction.AddParam(variable)
	} else if curCursor.GetCursorContext() == navigator.ContextTypeIfBlock {
		// curCursor.PushExpr(terminalString)

		//curVariable, err := curSymTable.GetVariableByName(terminalString)
		//if err != nil {
		//	panic(err)
		//}
		//curCursor.AddExprVariable(curVariable)

		//curExpr := curCursor.GetExpr()
		//curExpr.PushValue(curVariable)
	}

	return nil
}

func OperandContextHandler(contextParser *parser.OperandContext) error {
	//	fmt.Println("123456***")
	children := contextParser.GetChildren()

	for _, child := range children {
		fmt.Println("oooooooooooooooo")
		fmt.Printf("%T\n", child)
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
		//		fmt.Println("************************")
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
		} else if cursor.GetCursorContext() == navigator.ContextTypeIfBlock {
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
			//			fmt.Println("++++++++++++++ String", curVariable)
			//cursor.PrintStatement()
		} else if cursor.GetCursorContext() == navigator.ContextTypeIfBlock {
			cursor.PushExpr(terminalString)
		} else {
			curStatement.AddRightValue(curVariable)
		}

		//		fmt.Println("************************")
	}

	return nil
}
