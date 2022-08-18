package handler

import (
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
	"quinn007.com/uspace"
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
