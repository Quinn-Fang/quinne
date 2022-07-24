package handler

import (
	"fmt"

	"quinn007.com/listeners/utils"
	"quinn007.com/parser"
)

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

	return nil
}

func ExpressionListContextHandler(contextParser *parser.ExpressionListContext) error {
	fmt.Println("inside ExpressionListContextHandler... ")
	children := contextParser.GetChildren()
	splitter := ","

	identifierListStrings := make([]string, 0)
	for _, nodeContext := range children {
		terminalString, _ := utils.GetTerminalNodeText(nodeContext)
		if terminalString != splitter {
			identifierListStrings = append(identifierListStrings, terminalString)
		}
	}

	return nil
}
