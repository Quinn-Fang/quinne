package scanner

type ExprContext struct {
	exprString   string
	exprVarNames []string
}

func NewExprContext() *ExprContext {
	newExprContext := &ExprContext{
		exprVarNames: make([]string, 0),
	}
	return newExprContext
}

func (this *ExprContext) AppendExpr(subExpr string) {
	this.exprString += subExpr
}

func (this *ExprContext) AppendVarName(varName string) {
	this.exprVarNames = append(this.exprVarNames, varName)
}

func (this *ExprContext) GetExpr() string {
	return this.exprString
}

func (this *ExprContext) GetExprVarNames() []string {
	return this.exprVarNames
}
