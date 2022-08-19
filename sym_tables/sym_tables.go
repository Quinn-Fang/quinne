package sym_tables

import (
	"errors"
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/Quinn-Fang/Quinne/procedures"
	"github.com/Quinn-Fang/Quinne/utils"
	"github.com/Quinn-Fang/Quinne/variables"
)

type ScopeContext struct {
	scopeType    ContextType
	scopeContext interface{}
}

func NewScopeContext(scopeType ContextType) *ScopeContext {
	newScopeContext := &ScopeContext{
		scopeType: scopeType,
	}
	return newScopeContext
}

func (this *ScopeContext) GetScopeType() ContextType {
	return this.scopeType
}

func (this *ScopeContext) GetScopeContext() interface{} {
	return this.scopeContext
}

func (this *ScopeContext) SetScopeContext(scopeContext interface{}) {
	this.scopeContext = scopeContext
}

var (
	rootSymTable *SymTable
	curSymTable  *SymTable
)

type SymTable struct {
	prev                    *SymTable
	children                []*SymTable
	variableMap             map[string]*variables.Variable
	functions               []*procedures.FFunction
	executable              bool
	executableStack         []bool
	ifElseStack             []LogicSymbol
	ifElseExprStack         []string
	ifElseExprVariableStack []map[string]interface{}
	hasTrueBranch           bool

	scopeQueue *utils.Queue

	ifElseClauseList *utils.Queue
	curScope         *ScopeContext
}

func (this *SymTable) IsExecutable() bool {
	// temporary handle scope unhandled like for loop ...
	if this.curScope == nil {
		return true
	}

	functionDecl, ok := this.curScope.GetScopeContext().(*procedures.FFunctionDecl)
	if ok {
		if functionDecl.GetFunctionName() == "main" {
			if this.GetExecutable() != true {
				this.SetExecutable(true)
			}
			return true
		}
	}

	_, ok = this.curScope.GetScopeContext().(*ForLoop)
	if ok {
		if this.GetExecutable() != true {
			this.SetExecutable(true)
		}
		return true
	}

	checkPrevExecutable := this.CheckPrevExecutable()
	if !checkPrevExecutable {
		return false
	}
	fmt.Println(checkPrevExecutable)

	if !(this.curScope.GetScopeType() == ContextTypeIf || this.curScope.GetScopeType() == ContextTypeElseIf) {
		return true
	}

	ifElseBranch, ok := this.curScope.GetScopeContext().(*IfElseBranch)
	if !ok {
		panic("Not IfElseBranch ... \n")
	}

	ifElseClause := ifElseBranch.GetParent()
	ifElseExecutable := false
	if ifElseClause.HasTrueBranch() {
		ifElseExecutable = false
	} else {
		// map unset !!!!!
		// exprRes := this.JudgeIfElseExpr(ifElseBranch.GetExpr(), make)
		canExecute := ifElseBranch.GetJudgeRes()
		ifElseExecutable = canExecute
	}
	if this.GetExecutable() != ifElseExecutable {
		this.SetExecutable(ifElseExecutable)
	}

	return ifElseExecutable
}

func (this *SymTable) PrintIfElseClauseList() {
	fmt.Printf("------------------------ %+v --------------------------\n", "IfElseClauseList Start")

	for elem_1 := this.ifElseClauseList.GetFront(); elem_1 != nil; elem_1 = elem_1.Next() {
		ifElseClause := elem_1.Value.(*IfElseClause)
		for _, v1 := range ifElseClause.GetBranches() {
			fmt.Println()
			fmt.Printf("%+v  %+v  %+v \n", v1.GetBranchType(), v1.GetExpr(), v1.GetExprVarNames())
			fmt.Println()
		}
	}

	fmt.Printf("------------------------ %+v --------------------------\n", "IfElseClauseList End")

}

func (this *SymTable) AddIfElseClause(ifElseClause *IfElseClause) {
	this.ifElseClauseList.PushBack(ifElseClause)
}

func (this *SymTable) GetLastIfElseClause() *IfElseClause {
	return this.ifElseClauseList.GetBack().Value.(*IfElseClause)
}

func (this *SymTable) SetScope(scopeContext *ScopeContext) {
	this.curScope = scopeContext
}

func (this *SymTable) GetScope() *ScopeContext {
	return this.curScope
}

// check if this symTable executable, should be called after table context being parsed
func (this *SymTable) CheckExecutable() bool {
	parent := this.GetPrev()
	// root sym table always executable
	if parent == nil {
		return true
	}

	// check parents executable
	if !this.CheckPrevExecutable() {
		return false
	}

	// check ifElseClause currently resides executable
	return false
}

// simply check if all parents are executable
func (this *SymTable) CheckPrevExecutable() bool {
	prev := this.GetPrev()
	for prev != nil {
		if !prev.GetExecutable() {
			return false
		} else {
			prev = prev.GetPrev()
		}
	}
	return true
}

type ScopeQueue struct {
	queue *utils.Queue
}

func (this *ScopeQueue) PushBack(scopeContext *ScopeContext) {
	this.queue.PushBack(scopeContext)
}

func (this *ScopeQueue) PopFront() (*ScopeContext, error) {
	ret, err := this.queue.PopFront()
	return ret.(*ScopeContext), err
}

func (this *ScopeQueue) GetFront() (*ScopeContext, error) {
	ret, err := this.queue.PopFront()
	return ret.(*ScopeContext), err
}

type IfElseClauseQueue struct {
	queue *utils.Queue
}

func (this *IfElseClauseQueue) PushBack(ifElseClause *IfElseClause) {
	this.queue.PushBack(ifElseClause)
}

func (this *IfElseClauseQueue) PopFront() (*IfElseClause, error) {
	ret, err := this.queue.PopFront()
	return ret.(*IfElseClause), err
}

//////////////////////////////////////
func (this *SymTable) LookUpAndSetExecutable() bool {
	// parents up chain not executable
	if !this.GetChainExecutable() {
		return false
	} else {
		// check parent ifelse stack
		prev := this.GetPrev()
		if prev == nil {
			// root symtable
			return true
		} else {
			prev.ParseIfElseClause_1()
			if prev.ExecutableStackEmpty() {
				return true
			} else {
				return prev.PopFrontExecutableStack()
			}
		}
	}
}

func (this *SymTable) GetChainExecutable() bool {
	prev := this.GetPrev()
	for prev != nil {
		if prev.GetExecutable() == false {
			return false
		}
		prev = prev.GetPrev()
	}

	return true
}

// Called when entering block, on parent table
func (this *SymTable) ParseIfElseClause_1() {
	for !this.IfElseStackEmpty() {
		ifElseSymbol := this.PopFrontIfElseStack()
		if ifElseSymbol == LogicSymbolIf {
			if !this.hasTrueBranch {
				expr := this.PopFrontIfElseExprStack()
				exprVariable := this.PopFrontIfElseExprVariableStack()
				testRes := this.JudgeIfElseExpr(expr, exprVariable)
				if testRes == true {
					this.hasTrueBranch = true
					this.PushExecutableStack(true)
				} else {
					this.PushExecutableStack(false)
				}
			} else {
				this.PushExecutableStack(false)
			}
		} else if ifElseSymbol == LogicSymbolElse {
			// last else
			if len(this.ifElseStack) == 0 {
				if !this.hasTrueBranch {
					this.PushExecutableStack(true)
				} else {
					this.PushExecutableStack(false)
				}
				// set status to default
				this.hasTrueBranch = false
			} else {
				// else if branch, leave it to next if event, do nothing here

			}
		}
	}
}

////////////////////////////////////////////////
func (this *SymTable) GetIfElseExprVariableStack() []map[string]interface{} {
	return this.ifElseExprVariableStack
}

func (this *SymTable) PushIfElseExpVariableStack(ifElseExprVariable map[string]interface{}) {
	this.ifElseExprVariableStack = append(this.ifElseExprVariableStack, ifElseExprVariable)
}

func (this *SymTable) PopFrontIfElseExprVariableStack() map[string]interface{} {
	if !this.IfElseExprVariableStackEmpty() {
		ifElseExprVariable := this.ifElseExprVariableStack[0]
		this.ifElseExprVariableStack = this.ifElseExprVariableStack[1:]
		return ifElseExprVariable

	} else {
		panic("ifElseExprVariableStackEmpty")
	}
}

func (this *SymTable) IfElseExprVariableStackEmpty() bool {
	return len(this.ifElseExprVariableStack) == 0
}

////////////////////////////////////////////////
func (this *SymTable) GetIfElseExprStack() []string {
	return this.ifElseExprStack
}

func (this *SymTable) PushIfElseExprStack(ifElseExpr string) {
	this.ifElseExprStack = append(this.ifElseExprStack, ifElseExpr)
}

func (this *SymTable) PopFrontIfElseExprStack() string {
	if !this.IfElseExprStackEmpty() {
		ifElseExpr := this.ifElseExprStack[0]
		this.ifElseExprStack = this.ifElseExprStack[1:]
		return ifElseExpr

	} else {
		panic("ifElseExprStackStackEmpty")
	}
}

func (this *SymTable) IfElseExprStackEmpty() bool {
	return len(this.ifElseExprStack) == 0
}

////////////////////////////////////////////////
func (this *SymTable) GetExecutableStack() []bool {
	return this.executableStack
}

func (this *SymTable) PushExecutableStack(executable bool) {
	this.executableStack = append(this.executableStack, executable)
}

func (this *SymTable) PopFrontExecutableStack() bool {
	if !this.ExecutableStackEmpty() {
		executable := this.executableStack[0]
		this.executableStack = this.executableStack[1:]
		return executable

	} else {
		panic("ExecutableStackEmpty")
	}
}

func (this *SymTable) ExecutableStackEmpty() bool {
	return len(this.executableStack) == 0
}

////////////////////////////////////////////////
func (this *SymTable) GetIfElseStack() []LogicSymbol {
	return this.ifElseStack
}

func (this *SymTable) PushIfElseStack(logicSymbol LogicSymbol) {
	this.ifElseStack = append(this.ifElseStack, logicSymbol)
}

func (this *SymTable) PopFrontIfElseStack() LogicSymbol {
	if !this.IfElseStackEmpty() {
		logicSymbol := this.ifElseStack[0]
		this.ifElseStack = this.ifElseStack[1:]
		return logicSymbol

	} else {
		return LogicSymbolError
	}
}

func (this *SymTable) PopBackIfElseStack() LogicSymbol {
	if !this.IfElseStackEmpty() {
		logicSymbol := this.ifElseStack[len(this.ifElseStack)-1]
		this.ifElseStack = this.ifElseStack[:len(this.ifElseStack)-1]
		return logicSymbol

	} else {
		return LogicSymbolError
	}
}

func (this *SymTable) IfElseStackEmpty() bool {
	return len(this.ifElseStack) == 0
}

func SetRootSymTale(symTable *SymTable) {
	rootSymTable = symTable
}

func GetRootSymTale() *SymTable {
	return rootSymTable
}

func SetCurSymTable(symTable *SymTable) {
	curSymTable = symTable
}

func GetCurSymTable() *SymTable {
	return curSymTable
}

func NewEntryTable() *SymTable {
	newEntryTable := &SymTable{
		executable: true,
	}
	SetCurSymTable(newEntryTable)
	SetRootSymTale(newEntryTable)
	return newEntryTable
}

func NewSymTable(prevSymTable *SymTable) *SymTable {
	var newSymTable *SymTable
	newSymTable = &SymTable{
		prev:                    prevSymTable,
		variableMap:             make(map[string]*variables.Variable),
		functions:               make([]*procedures.FFunction, 0),
		executableStack:         make([]bool, 0),
		ifElseStack:             make([]LogicSymbol, 0),
		ifElseExprStack:         make([]string, 0),
		ifElseExprVariableStack: make([]map[string]interface{}, 0),
		ifElseClauseList:        utils.NewQueue(),

		scopeQueue: utils.NewQueue(),
	}
	// prevSymTable.next = newSymTable
	prevSymTable.children = append(prevSymTable.children, newSymTable)

	return newSymTable
}

func (this *SymTable) AddVariable(newVariable *variables.Variable) error {
	if _, ok := this.variableMap[newVariable.GetVariableName()]; ok {
		return errors.New("Variable exists")
	} else {
		this.variableMap[newVariable.GetVariableName()] = newVariable
		return nil
	}
}

func (this *SymTable) GetVariableByName(variableName string) (*variables.Variable, error) {
	// look for the variable up chain
	if this == rootSymTable {
		return variables.NewEmptyVariable(), errors.New("variable does not exist")
	}
	if variable, ok := this.variableMap[variableName]; !ok {
		prevSymTable := this.GetPrev()
		if parentVariable, err := prevSymTable.GetVariableByName(variableName); err != nil {
			return variables.NewEmptyVariable(), err
		} else {
			return parentVariable, nil
		}
	} else {
		return variable, nil
	}
}

func (this *SymTable) GetVariables() map[string]*variables.Variable {
	return this.variableMap
}

func (this *SymTable) AddFunction(newFunction *procedures.FFunction) error {
	this.functions = append(this.functions, newFunction)
	return nil
}

func (this *SymTable) GetFunctions() []*procedures.FFunction {
	return this.functions
}

func (this *SymTable) GetLastFunction() *procedures.FFunction {
	return this.functions[len(this.functions)-1]
}

func (this *SymTable) GetPrev() *SymTable {
	return this.prev
}

func (this *SymTable) GetChildren() []*SymTable {
	return this.children
}

func (this *SymTable) SetExecutable(isExecutable bool) {
	this.executable = isExecutable
}

func (this *SymTable) GetExecutable() bool {
	return this.executable
}

//func (this *SymTable) LookUpExecutable() bool {
//	if this.IfElseStackEmpty() {
//		return true
//	}
//
//	ifElseSymbol := this.PopFrontIfElseStack()
//	if ifElseSymbol == LogicSymbolIf {
//
//	}
//
//}

func (this *SymTable) JudgeIfElseExpr(expr string, parameters map[string]interface{}) bool {
	expression, err := govaluate.NewEvaluableExpression(expr)

	// parameters := make(map[string]interface{}, 8)

	result, err := expression.Evaluate(parameters)
	if err != nil {
		panic(err)
	}

	if ret, ok := result.(bool); ok {
		return ret
	} else {
		panic("result is not bool")
	}
}

func (this *SymTable) PrintFunctions() {
	for _, v := range this.GetFunctions() {
		fmt.Printf("%+v\n", v)
	}
}
