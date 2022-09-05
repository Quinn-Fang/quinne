package handlerV2

import (
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func SourceFileHandler(antlrCtx antlr.ParserRuleContext, scanner *scanner.Scanner) {
	children := antlrCtx.GetChildren()

	// utils.PrintChildren(children)

	for _, child := range children {
		switch parserContext := child.(type) {
		case *antlr.TerminalNodeImpl:
			{
			}
		case *parser.FunctionDeclContext:
			{
				FunctionDeclHandler(parserContext, scanner)
			}
		}
	}

}
