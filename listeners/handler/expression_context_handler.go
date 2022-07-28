package handler

import (
	"fmt"

	"quinn007.com/listeners/utils"
	"quinn007.com/parser"
)

func ExpressionStmtContextHandler(contextParser *parser.ExpressionStmtContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ExpressionContext:
			{
				fmt.Println("123456")
				ExpressionContextHandler(parserContext)

			}
		}
	}

	return nil
}

func ExpressionContextHandler(contextParser *parser.ExpressionContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
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
		fmt.Println("99999999")
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
	fmt.Println("777777777777777777777777")
	terminalString, _ := utils.GetTerminalNodeText(children[0])
	fmt.Println(terminalString)

	//for _, child := range children {
	//	fmt.Println("777777777777777777777777")
	//	fmt.Printf("%T\n", child)
	//	fmt.Printf("%+v\n", child)
	//	fmt.Println("777777777777777777777777")
	//	switch parserContext := child.(type) {
	//	case *parser.IdentifierListContext:
	//		{
	//			IdentifierListContextHandler(parserContext)
	//		}
	//	case *parser.LiteralContext:
	//		{
	//			fmt.Println("998998998")
	//			LiteralContextHandler(parserContext)
	//		}

	//	}
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
		switch parserContext := child.(type) {
		case *parser.String_Context:
			{
				StringContextHandler(parserContext)
			}
		}
	}

	return nil
}

func StringContextHandler(contextParser *parser.String_Context) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		fmt.Println("************************")
		fmt.Printf("%T\n", child)
		fmt.Printf("%+v\n", child)
		fmt.Println("************************")
		switch parserContext := child.(type) {
		case *parser.IdentifierListContext:
			{
				IdentifierListContextHandler(parserContext)
			}
		case *parser.ExpressionListContext:
			{
				ExpressionListContextHandler(parserContext)
			}

		}
	}

	return nil
}
