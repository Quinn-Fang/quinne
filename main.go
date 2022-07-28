package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/listeners"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
)

func runListener() {
	// Setup the input
	// is := antlr.NewInputStream("package main \n 1 +     2 * 3+1+2344")
	// input, _ := antlr.NewFileStream(os.Args[1])
	input, _ := antlr.NewFileStream("samples/sample_4.go")
	fmt.Println("1111111")

	// Create the Lexer
	// lexer := parser.NewGoLexer(is)
	lexer := parser.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGoParser(stream)

	//listen := GoListener{}
	// Finally parse the expression
	// antlr.ParseTreeWalkerDefault.Walk(&listen, p.SourceFile())
	tree := p.SourceFile()
	// antlr.ParseTreeWalkerDefault.Walk(NewGoListener(p, tree), tree)
	antlr.ParseTreeWalkerDefault.Walk(listeners.NewGoListener(p, tree), tree)
	sym_tables.NewSymTable(nil)

}

func main() {
	runListener()

}
