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
	//fmt.Println("999999")
	//fmt.Println("-----------------")
	//fmt.Println(ctx.GetText())
	i := ctx.GetRuleIndex()
	ruleName := this.p.RuleNames[i]
	//fmt.Println("the rule name is: ", ruleName)
	//fmt.Println(" and i is : ", i)
	//fmt.Println(" and  ruleName is : ", this.p.RuleNames)

	//fmt.Println(ctx.GetChildCount())
	LexDispatcher(this, ctx, ruleName)
}
