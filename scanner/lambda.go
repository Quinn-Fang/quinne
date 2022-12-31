package scanner

import (
	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/variables"
)

type LambdaContext struct {
	// Left value, parts before :
	lParams []string
	lRet    string
	// Right value, parts after : ,
	// usually expression if condition else condition
	//lExprRaw string
	// Target expr string, translated using left and right
	// values
	//lExpr    string
	lSubExpr string
	// temporary logic, 0: return value for if condition,
	// 1: if expression(true return value)
	// 2: else expression(false return value)
	lExprList  []string
	lambdaDecl *procedures.LambdaDecl
	//lIfElseClauseCtx      *LambdaIfElseContext
	//lIfElseClauseCtxEntry *LambdaIfElseContext
}

func NewLambdaContext() *LambdaContext {
	newLambdaContext := &LambdaContext{
		lParams:    make([]string, 0),
		lambdaDecl: procedures.NewLambdaDecl(),
		lExprList:  make([]string, 0),
	}
	return newLambdaContext
}

func (this *LambdaContext) GetLambdaDecl() *procedures.LambdaDecl {
	return this.lambdaDecl
}

func (this *LambdaContext) AddParam(paramName string) {
	this.lParams = append(this.lParams, paramName)
}

func (this *LambdaContext) AppendExprList(exprStr string) {
	this.lExprList = append(this.lExprList, exprStr)
}

func (this *LambdaContext) AppendLReturn(retValue string) {
	this.lRet += retValue
}

//func (this *LambdaContext) SetLReturn(retValue string) {
//	this.lRet = retValue
//}

func (this *LambdaContext) AppendSubExpr(subExpr string) {
	this.lSubExpr += subExpr
}

func (this *LambdaContext) GetSubExpr() string {
	return this.lSubExpr
}

func (this *LambdaContext) ClearSubExpr() {
	this.lSubExpr = ""
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
	this.lParams = make([]string, 0)
}

func (this *LambdaContext) ToTernaryExpr() string {
	// simple: ifcond?return:else
	// clustered: ifcond?(ifcond?return:else):(ifcond?return:else)
	lambdaDecl := this.GetLambdaDecl()
	lambdaExpression := lambdaDecl.GetFirstLambdaExpression()
	if lambdaExpression == nil {
		panic("lambda expression does not exists!")
	}

	return ToLambdaExpressionString(lambdaExpression)
}

func ToLambdaExpressionString(lExpr *procedures.LambdaExpression) string {
	if lExpr == nil {
		// The caller of this function should've checked nullable of
		// lambda expression already.
		panic("LambdaExpression is nil !")
	}
	ret := ""
	lambdaReturnStr := lExpr.GetReturnValueString()
	ret += lambdaReturnStr
	ifCond := lExpr.GetIfCond()
	if ifCond == nil {
		// No more succeeding if condition
		return ret
	}
	// if condition exists.
	ifCondStr := lExpr.GetIfCondStr()
	ret = ifCondStr + "?" + ret

	// if ifCond is not nil and no more else, then we lose default
	// value, so next lambdaExpression(else) must exists.
	if lExpr.GetNextExpression() == nil {
		panic("No more succeeding LambdaExpression(else)!")
	}

	// LambdaExpression exists.
	ret += ":" + ToLambdaExpressionString(lExpr.GetNextExpression())
	return "(" + ret + ")"
}

type LambdaCallContext struct {
	lambdaCall *procedures.LambdaCall
}

func NewLambdaCallContext() *LambdaCallContext {
	return &LambdaCallContext{}
}

func (this *LambdaCallContext) SetLambdaCall(lCall *procedures.LambdaCall) {
	this.lambdaCall = lCall
}

func (this *LambdaCallContext) AddArgs(newArg *variables.Variable) {
	this.lambdaCall.AddArgs(newArg)
}

func (this *LambdaCallContext) SetRetValue(retValue *variables.Variable) {
	this.lambdaCall.SetReturnValue(retValue)
}
