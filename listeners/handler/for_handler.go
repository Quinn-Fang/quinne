package handler

import (
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/uspace"
)

func ForStatementContextHandler(contextParser *parser.ForStmtContext) error {

	curSymTable := sym_tables.GetCurSymTable()
	forLoop := sym_tables.NewForLoop()
	curNavigator := navigator.GetCurNavigator()

	curNavigator.AddEvent(uspace.EventTypeForLoop, forLoop, curSymTable)

	scope := sym_tables.NewScopeContext(sym_tables.ContextTypeForLoop)
	scope.SetScopeContext(forLoop)

	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.BlockContext:
			{
				BlockContextHandler(parserContext, scope)
			}
		}
	}
	return nil
}
