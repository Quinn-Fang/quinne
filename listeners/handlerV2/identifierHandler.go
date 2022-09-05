package handlerV2

import (
	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
)

func IdentifierListContextHandler(contextParser *parser.IdentifierListContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	splitter := ","

	identifierListStrings := make([]string, 0)
	for _, nodeContext := range children {
		terminalString, _ := utils.GetTerminalNodeText(nodeContext)
		if terminalString != splitter {
			identifierListStrings = append(identifierListStrings, terminalString)
		}
	}

	curCursor, _ := navigator.GetCursor()
	curStatement := curCursor.GetStatement()
	for _, variableName := range identifierListStrings {
		curStatement.AddLeftValue(variableName)
	}

	return nil
}

func VarDeclContextHandler(contextParser *parser.VarDeclContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.VarSpecContext:
			{
				VarSpecContextHandler(parserContext, scanner)
			}
		}
	}
	return nil
}

func ShortVarDeclContextHandler(contextParser *parser.ShortVarDeclContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.IdentifierListContext:
			{
				IdentifierListContextHandler(parserContext, scanner)
			}
		case *parser.ExpressionListContext:
			{
				ExpressionListContextHandler(parserContext, scanner)
			}

		}
	}
	return nil
}

func VarSpecContextHandler(contextParser *parser.VarSpecContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.IdentifierListContext:
			{
				IdentifierListContextHandler(parserContext, scanner)
			}
		case *parser.ExpressionListContext:
			{
				ExpressionListContextHandler(parserContext, scanner)
			}

		}
	}

	return nil
}
