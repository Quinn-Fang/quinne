package handler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/uspace"
)

func SourceFileHandler(antlrCtx antlr.ParserRuleContext) {
	children := antlrCtx.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *antlr.TerminalNodeImpl:
			{
				//terminalString, _ := utils.GetTerminalNodeText(parserContext)
				//fmt.Println(terminalString)
			}
		case *parser.FunctionDeclContext:
			{
				FunctionDeclHandler(parserContext)
			}
		}
	}

}

func FunctionDeclHandler(contextParser *parser.FunctionDeclContext) error {

	curNavigator := navigator.GetCurNavigator()
	curSymTable := sym_tables.GetCurSymTable()

	children := contextParser.GetChildren()
	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.BlockContext:
			{
				// scopeType := sym_tables.ContextTypeFuncDecl
				//	curNavigator.AddEvent(uspace.EventTypeIfElseExpr, newElseIfBranch, curSymTable)
				scopeContext := procedures.NewFunctionDecl("main")
				blockContext := sym_tables.NewScopeContext(sym_tables.ContextTypeFuncDecl)
				blockContext.SetScopeContext(scopeContext)

				curNavigator.AddEvent(uspace.EventTypeFunctionDecl, scopeContext, curSymTable)
				BlockContextHandler(parserContext, blockContext)
			}
		}
	}

	return nil
}