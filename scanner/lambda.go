package scanner

import (
	"strings"

	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/variables"
)

type LambdaContext struct {
	// Left value, parts before :
	lParams []string
	// Right value, parts after : ,
	// usually expression if condition else condition
	lExprRaw string
	// Target expr string, translated using left and right
	// values
	lExpr      string
	lambdaDecl *procedures.LambdaDecl
}

func NewLambdaContext() *LambdaContext {
	newLambdaContext := &LambdaContext{
		lParams:    make([]string, 0),
		lambdaDecl: procedures.NewLambdaDecl(),
	}
	return newLambdaContext
}

func (this *LambdaContext) AddParam(paramName string) {
	this.lParams = append(this.lParams, paramName)
}

func (this *LambdaContext) AppendExprRaw(exprRaw string) {
	this.lExprRaw += exprRaw
}

func (this *LambdaContext) AddLambdaDeclParams(vType variables.VTypeEnum) {
	// 1. Create new variable by scanner's saved parameter names
	// 2. Add those variables to actual lambdaDecl
	// 3. Clear current lParams list
	// 4. Add parameter Names to raw string waiting for conversion
	for _, vName := range this.lParams {
		curVariable := variables.NewVariable(
			vName,
			vType,
			"",
			-10,
		)
		this.lambdaDecl.AddParam(curVariable)
	}
	this.AppendExprRaw(strings.Join(this.lParams, ","))
	this.lParams = make([]string, 0)
}
