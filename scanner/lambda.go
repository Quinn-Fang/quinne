package scanner

import (
	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/variables"
)

type LambdaIfElseContext struct {
	lBranchSymbol sym_tables.LogicSymbol
	lExpr         string
	lElseClause   *LambdaIfElseContext
	prevClause    *LambdaIfElseContext
}

// init with "if" means child if clause, "none" means no child, thus else clause
func NewLambdaIfElseContext() *LambdaIfElseContext {
	lambdaIfElseContext := &LambdaIfElseContext{}
	return lambdaIfElseContext
}

func (this *LambdaIfElseContext) AppendExpr(lExpr string) {
	this.lExpr += lExpr
}

func (this *LambdaIfElseContext) AddElseClause(elseClause *LambdaIfElseContext) {
	if elseClause == nil {
		panic("Can not set elseclause to nil !")
	}
	this.lElseClause = elseClause
	elseClause.SetPrevClause(this)
}

func (this *LambdaIfElseContext) SetPrevClause(prevLambdaIfElseCtx *LambdaIfElseContext) {
	this.prevClause = prevLambdaIfElseCtx
}

func (this *LambdaIfElseContext) SetBranchSymbol(logicSymbol sym_tables.LogicSymbol) {
	this.lBranchSymbol = logicSymbol
}

func (this *LambdaIfElseContext) GetBranchSymbol() sym_tables.LogicSymbol {
	return this.lBranchSymbol
}

// After fully parsed the whole lambdaifelseclause, recursived traverse children and
// return the parsed if else clause as string list expr
func (this *LambdaIfElseContext) ToExprList() []string {
	exprList := make([]string, 0)
	// exprList = append(exprList, string(this.lIfSymbol))
	exprList = append(exprList, this.lExpr)
	if this.lElseClause != nil {
		exprList = append(exprList, string(sym_tables.LogicSymbolElse))
		exprList = append(exprList, this.lElseClause.ToExprList()...)
	}
	return exprList
}

type LambdaContext struct {
	// Left value, parts before :
	lParams []string
	lRet    string
	// Right value, parts after : ,
	// usually expression if condition else condition
	lExprRaw string
	// Target expr string, translated using left and right
	// values
	lExpr    string
	lSubExpr string
	// temporary logic, 0: return value for if condition,
	// 1: if expression(true return value)
	// 2: else expression(false return value)
	lExprList             []string
	lambdaDecl            *procedures.LambdaDecl
	lIfElseClauseCtx      *LambdaIfElseContext
	lIfElseClauseCtxEntry *LambdaIfElseContext
}

func NewLambdaContext() *LambdaContext {
	newLambdaContext := &LambdaContext{
		lParams:    make([]string, 0),
		lambdaDecl: procedures.NewLambdaDecl(),
		lExprList:  make([]string, 0),
	}
	return newLambdaContext
}

func (this *LambdaContext) AddParam(paramName string) {
	this.lParams = append(this.lParams, paramName)
}

func (this *LambdaContext) AppendExprRaw(exprRaw string) {
	this.lExprRaw += exprRaw
}

func (this *LambdaContext) AppendExprList(exprStr string) {
	this.lExprList = append(this.lExprList, exprStr)
}

func (this *LambdaContext) SetLReturn(retValue string) {
	this.lRet = retValue
}

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
	//for _, vName := range this.lParams {
	//	curVariable := variables.NewVariable(
	//		vName,
	//		vType,
	//		"",
	//		-10,
	//	)
	//	this.lambdaDecl.AddParam(curVariable)
	//}
	//if len(this.lExprRaw) > 0 {
	//	this.AppendExprRaw(",")
	//}
	//this.AppendExprRaw(strings.Join(this.lParams, ","))
	//this.lParams = make([]string, 0)
}

func (this *LambdaContext) ToTernaryExpr() string {
	ret := ""
	ret += this.lExprList[1]
	ret += "?"
	ret += this.lRet
	ret += ":"
	ret += this.lExprList[3]
	return ret
}

func (this *LambdaContext) SetLambdaIfElseClause(lIfElseClause *LambdaIfElseContext) {
	this.lIfElseClauseCtx = lIfElseClause
}

func (this *LambdaContext) SetLambdaIfElseClauseEntry(lIfElseClause *LambdaIfElseContext) {
	this.lIfElseClauseCtxEntry = lIfElseClause
}

func (this *LambdaContext) GetLambdaIfElseClauseEntry() *LambdaIfElseContext {
	return this.lIfElseClauseCtxEntry
}

func (this *LambdaContext) NewLambdaIfElseClause() *LambdaIfElseContext {
	lambdaIfElseContext := NewLambdaIfElseContext()
	return lambdaIfElseContext
}
