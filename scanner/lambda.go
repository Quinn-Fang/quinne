package scanner

type LambdaContext struct {
	// Left value, parts before :
	lParams []string
	// Right value, parts after : ,
	// usually expression if condition else condition
	lExprRaw string
	// Target expr string, translated using left and right
	// values
	lExpr string
}

func NewLambdaContext() *LambdaContext {
	newLambdaContext := &LambdaContext{
		lParams: make([]string, 0),
	}
	return newLambdaContext
}

func (this *LambdaContext) AddParam(paramName string) {
	this.lParams = append(this.lParams, paramName)
}

func (this *LambdaContext) AppendExprRaw(exprRaw string) {
	this.lExprRaw += exprRaw
}
