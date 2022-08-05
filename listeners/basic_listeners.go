package listeners

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/parser"
)

type GoListener struct {
	*parser.BaseGoParserListener
	p *parser.GoParser
	t antlr.Tree
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
