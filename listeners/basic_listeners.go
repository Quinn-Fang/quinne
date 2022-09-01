package listeners

import (
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type GoListener struct {
	*parser.BaseGoParserListener
	goParser *parser.GoParser
	ast      antlr.Tree
	scanner  *scanner.Scanner
}

func NewGoListener(goParser *parser.GoParser, ast antlr.Tree, scanner *scanner.Scanner) *GoListener {
	return &GoListener{
		goParser: goParser,
		ast:      ast,
		scanner:  scanner,
	}
}

func (this *GoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	i := ctx.GetRuleIndex()
	ruleName := this.goParser.RuleNames[i]
	LexDispatcher(this, ctx, ruleName, this.scanner)
}
