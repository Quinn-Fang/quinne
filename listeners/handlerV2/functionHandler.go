package handlerV2

import (
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/uspace"
)

func FunctionHandler(operandContext *parser.PrimaryExprContext, argumentsContext *parser.ArgumentsContext, scanner *scanner.Scanner) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionName)

	PrimaryExprContextHandler(operandContext, scanner)

	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionArgs)

	ArgumentsContextHandler(argumentsContext, scanner)

	curCursor.SetCursorContext(sym_tables.ContextTypeDefault)

	curSymTable := sym_tables.GetCurSymTable()
	fFunction := curSymTable.GetLastFunction()
	curNavigator := navigator.GetCurNavigator()
	if !curCursor.IsAppendingExpr() {
		// regular function
		curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)
	} else {
		curCursor.PushExpr(fFunction.ToString())
		curCursor.AddExprVarNames(fFunction.FName)
	}

	// curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)

	return nil
}

func ArgumentsContextHandler(contextParser *parser.ArgumentsContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ExpressionListContext:
			{
				ExpressionListContextHandler(parserContext, scanner)
			}
		}
	}

	return nil
}

func FunctionDeclHandler(contextParser *parser.FunctionDeclContext, scanner *scanner.Scanner) error {

	curNavigator := navigator.GetCurNavigator()
	curSymTable := sym_tables.GetCurSymTable()

	children := contextParser.GetChildren()
	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.BlockContext:
			{
				scopeContext := procedures.NewFunctionDecl("main")
				blockContext := sym_tables.NewScopeContext(sym_tables.ContextTypeFuncDecl)
				blockContext.SetScopeContext(scopeContext)

				curNavigator.AddEvent(uspace.EventTypeFunctionDecl, scopeContext, curSymTable)
				BlockContextHandler(parserContext, blockContext, scanner)
			}
		}
	}

	return nil
}
