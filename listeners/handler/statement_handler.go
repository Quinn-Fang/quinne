package handler

import (
	"github.com/Quinn-Fang/quinne/parser"
)

func StatementListHandler(contextParser *parser.StatementListContext) error {
	//fmt.Println("inside StatementListener ... .......")
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.StatementContext:
			{
				StatementHandler(parserContext)
			}

		}
	}
	return nil
}

func StatementHandler(contextParser *parser.StatementContext) error {
	//fmt.Println("inside StatementListener ... .......")
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.SimpleStmtContext:
			{
				SimpleStmtContextHandler(parserContext)
			}
		case *parser.IfStmtContext:
			{
				IfElseStmtContextHandler(parserContext)
			}
		case *parser.ForStmtContext:
			{
				ForStatementContextHandler(parserContext)
			}
		}
	}
	return nil
}
