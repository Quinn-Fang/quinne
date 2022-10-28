package sym_tables

type LogicSymbol string

const (
	LogicSymbolError LogicSymbol = "logicSymbolError"
	LogicSymbolIf                = "if"
	LogicSymbolElse              = "else"
	LogicSymbolNil               = "nil"
)

type BranchType int

const (
	BranchTypeIf     BranchType = 1
	BranchTypeElseIf            = 2
	BranchTypeElse              = 3
)

type IfElseClause struct {
	ifElseBranches []*IfElseBranch
	hasTrueBranch  bool
	symTable       *SymTable
}

func NewIfElseClause() *IfElseClause {
	newIfElseClause := &IfElseClause{}
	return newIfElseClause
}

func (this *IfElseClause) AddBranch(curIfElseBranch *IfElseBranch) {
	this.ifElseBranches = append(this.ifElseBranches, curIfElseBranch)
}

func (this *IfElseClause) SetHasTrueBranch(hasTrueBranch bool) {
	this.hasTrueBranch = hasTrueBranch
}

func (this *IfElseClause) HasTrueBranch() bool {
	return this.hasTrueBranch
}

func (this *IfElseClause) GetBranches() []*IfElseBranch {
	return this.ifElseBranches
}

type IfElseBranch struct {
	branchType   BranchType
	expr         string
	exprVarNames []string
	ifElseClause *IfElseClause
	judgeRes     bool
}

func (this *IfElseBranch) SetJudgeRes(judgeRes bool) {
	this.judgeRes = judgeRes
}

func (this *IfElseBranch) GetJudgeRes() bool {
	return this.judgeRes
}

func NewIfElseBranch(curBranchType BranchType) *IfElseBranch {
	newIfElseBranch := &IfElseBranch{
		branchType: curBranchType,
	}
	return newIfElseBranch
}

func (this *IfElseBranch) SetParent(ifElseClause *IfElseClause) {
	this.ifElseClause = ifElseClause
}

func (this *IfElseBranch) GetParent() *IfElseClause {
	return this.ifElseClause
}

func (this *IfElseBranch) GetBranchType() BranchType {
	return this.branchType
}

func (this *IfElseBranch) SetExpr(curExpr string) {
	this.expr = curExpr
}

func (this *IfElseBranch) GetExpr() string {
	return this.expr
}

func (this *IfElseBranch) SetExprVarNames(varNames []string) {
	this.exprVarNames = varNames
}

func (this *IfElseBranch) GetExprVarNames() []string {
	return this.exprVarNames
}
