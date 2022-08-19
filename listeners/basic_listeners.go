package listeners

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/Quinn-Fang/quinne/parser"
)

type GoListener struct {
	*parser.BaseGoParserListener
	p      *parser.GoParser
	t      antlr.Tree
	Test_1 bool
}

func NewGoListener(p *parser.GoParser, t antlr.Tree) *GoListener {
	return &GoListener{
		p: p,
		t: t,
	}
}

func (this *GoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	i := ctx.GetRuleIndex()
	ruleName := this.p.RuleNames[i]
	LexDispatcher(this, ctx, ruleName)
}
