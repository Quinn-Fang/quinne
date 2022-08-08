package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/listeners"
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
	"quinn007.com/sym_tables/utils"
)

func runListener() {
	input, _ := antlr.NewFileStream("samples/sample_4.go")
	// Create First SymTable
	sym_tables.NewEntryTable()
	// Create Cursor
	navigator.InitCursor()
	curNavigator := navigator.NewNavigator()
	navigator.SetCurNavigator(curNavigator)

	// Create the Lexer
	lexer := parser.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGoParser(stream)

	//listen := GoListener{}
	tree := p.SourceFile()
	antlr.ParseTreeWalkerDefault.Walk(listeners.NewGoListener(p, tree), tree)
	utils.PrintAllSymTale()
	curNavigator.PrintStack()

	curNavigator.PrintCodeSegments()
}

func main() {
	runListener()

}
