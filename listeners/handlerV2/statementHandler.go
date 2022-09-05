package handlerV2

import (
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/sym_tables"
)

func ExpressionStmtContextHandler(contextParser *parser.ExpressionStmtContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext, scanner)
			}
		}
	}

	return nil
}

func StatementListHandler(contextParser *parser.StatementListContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.StatementContext:
			{
				StatementHandler(parserContext, scanner)
			}

		}
	}
	return nil
}

func StatementHandler(contextParser *parser.StatementContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.SimpleStmtContext:
			{
				SimpleStmtContextHandler(parserContext, scanner)
			}
		case *parser.IfStmtContext:
			{
				IfElseStmtContextHandler(parserContext, scanner)
			}
		case *parser.ForStmtContext:
			{
				ForStatementContextHandler(parserContext, scanner)
			}
		}
	}
	return nil
}

func SimpleStmtContextHandler(contextParser *parser.SimpleStmtContext, scanner *scanner.Scanner) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.NewStatement()
	curStatement := curCursor.GetStatement()
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ShortVarDeclContext:
			{
				ShortVarDeclContextHandler(parserContext, scanner)
			}
		case *parser.VarSpecContext:
			{
				VarSpecContextHandler(parserContext, scanner)
			}
		case *parser.ExpressionStmtContext:
			{
				ExpressionStmtContextHandler(parserContext, scanner)
			}
		}
	}
	rightValues := curStatement.GetRightValues()
	curSymTable := sym_tables.GetCurSymTable()
	curStatement.Assign()

	// add assigned variable to symbol table
	for _, variable := range rightValues {
		curSymTable.AddVariable(variable)
	}

	return nil
}
