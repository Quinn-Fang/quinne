package handler

import (
	"fmt"

	"quinn007.com/listeners/utils"
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/variables"
)

func VarDeclContextHandler(contextParser *parser.VarDeclContext) error {
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.VarSpecContext:
			{
				VarSpecContextHandler(parserContext)
			}
		}
	}
	return nil
}

func ShortVarDeclContextHandler(contextParser *parser.ShortVarDeclContext) error {
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
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

func VarSpecContextHandler(contextParser *parser.VarSpecContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
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

func IdentifierListContextHandler(contextParser *parser.IdentifierListContext) error {
	children := contextParser.GetChildren()
	splitter := ","

	identifierListStrings := make([]string, 0)
	for _, nodeContext := range children {
		terminalString, _ := utils.GetTerminalNodeText(nodeContext)
		if terminalString != splitter {
			identifierListStrings = append(identifierListStrings, terminalString)
		}
	}
	//	fmt.Println("55555555")
	//	fmt.Println(identifierListStrings)

	curCursor, _ := navigator.GetCursor()
	curStatement := curCursor.GetStatement()
	for _, variableName := range identifierListStrings {
		curStatement.AddLeftValue(variableName)
	}

	return nil
}

func ExpressionListContextHandler(contextParser *parser.ExpressionListContext) error {
	fmt.Println("inside ExpressionListContextHandler... ")
	children := contextParser.GetChildren()
	splitter := ","

	ExpressionListStrings := make([]string, 0)
	curCursor, _ := navigator.GetCursor()
	curStatement := curCursor.GetStatement()
	for _, nodeContext := range children {
		fmt.Println("*********************")
		fmt.Printf("%T\n", nodeContext)
		fmt.Printf("%+v\n", nodeContext)
		switch parserContext := nodeContext.(type) {
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext)
			}
		default:
			{
				// Splitter

				terminalString, _ := utils.GetTerminalNodeText(nodeContext)
				if terminalString != splitter {
					newVariable := variables.NewVariable(
						"",
						variables.VTypeUndefined,
						terminalString,
						curCursor.GetIndex())

					curStatement.AddRightValue(newVariable)
					curCursor.IncreaseIndex()

					ExpressionListStrings = append(ExpressionListStrings, terminalString)
				}
			}
		}

	}
	//	fmt.Println("666666666")
	//	fmt.Println(ExpressionListStrings)

	return nil
}
