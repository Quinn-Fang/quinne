package handlerV2

import (
	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/variables"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func Type_ContextHandler(contextParser *parser.Type_Context, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.TypeNameContext:
			{
				TypeNameContextHandler(parserContext, scanner)
			}
		}
	}
	return nil

}

func TypeNameContextHandler(contextParser *parser.TypeNameContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	for _, child := range children {
		switch parserContext := child.(type) {
		case *antlr.TerminalNodeImpl:
			{
				terminalString, _ := utils.GetTerminalNodeText(parserContext)
				vType := variables.StrToVType(terminalString)
				scanner.AddLambdaParamToDecl(vType)
				//cursor.IncreaseIndex()
			}
		}
	}
	return nil

}
