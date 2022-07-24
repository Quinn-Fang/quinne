package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/listeners"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
)

//type GoListener struct {
//	*parser.BaseGoParserListener
//	p *parser.GoParser
//	t antlr.Tree
//}
//
//func NewGoListener(p *parser.GoParser, t antlr.Tree) *GoListener {
//	return &GoListener{
//		p: p,
//		t: t,
//	}
//}
//
//func (this *GoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	fmt.Println("999999")
//	fmt.Println("-----------------")
//	fmt.Println(ctx.GetText())
//	i := ctx.GetRuleIndex()
//	ruleName := this.p.RuleNames[i]
//	fmt.Println("the rule name is: ", ruleName)
//
//	fmt.Println(ctx.GetChildCount())
//}
//
//func (this *GoListener) EnterShortVarDecl(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText())
//}

//func (this *GoListener) EnterVarDecl(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText())
//}

//func (this *GoListener) EnterDeclaration(ctx antlr.ParserRuleContext) {
//	fmt.Println("3333333")
//	fmt.Println(ctx.GetText())
//}
//func (this *GoListener) EnterAssignment(ctx antlr.ParserRuleContext) {
//	fmt.Println("333333")
//	fmt.Println(ctx.GetText())
//}
//
////func (this *GoListener) EnterVarSpec(ctx antlr.ParserRuleContext) {
////	fmt.Println(ctx.GetText())
////}
//func (this *GoListener) EnterSimpleStmt(ctx antlr.ParserRuleContext) {
//	fmt.Println("222222")
//	fmt.Println(ctx.GetText())
//}
//
//func (this *GoListener) ExitFunctionDecl(ctx antlr.ParserRuleContext) {
//	fmt.Println("444444")
//	fmt.Println(ctx.GetText())
//}

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
