package sym_tables

type LogicSymbol string

const (
	LogicSymbolError LogicSymbol = "logicSymbolError"
	LogicSymbolIf                = "if"
	LogicSymbolElse              = "else"
)

type BranchType int

const (
	BranchTypeIf     BranchType = 1
	BranchTypeElseIf            = 2
	BranchTypeElse              = 3
)

type IfElseClause struct {
	ifElseBranch  []*IfElseBranch
	hasTrueBranch bool
	symTable      *SymTable
}

// pass in where this if else clause resides
func NewIfElseClause(curSymTable *SymTable) *IfElseClause {
	newIfElseClause := &IfElseClause{
		symTable: curSymTable,
	}
	return newIfElseClause
}

func (this *IfElseClause) AddBranch(curIfElseBranch *IfElseBranch) {
	this.ifElseBranch = append(this.ifElseBranch, curIfElseBranch)
}

func (this *IfElseClause) SetHasTrueBranch(hasTrueBranch bool) {
	this.hasTrueBranch = hasTrueBranch
}

type IfElseBranch struct {
	branchType   BranchType
	expr         string
	exprVarNames []string
	ifElseClause *IfElseClause
}

func NewIfElseBranch(curBranchType BranchType) *IfElseBranch {
	newIfElseBranch := &IfElseBranch{
		branchType: curBranchType,
	}
	return newIfElseBranch
}

func (this *IfElseBranch) SetExpr(curExpr string) {
	this.expr = curExpr
}

func (this *IfElseBranch) SetExprVarNames(varNames []string) {
	this.exprVarNames = varNames
}
