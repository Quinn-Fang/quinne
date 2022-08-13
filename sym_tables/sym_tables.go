package sym_tables

import (
	"errors"
	"fmt"

	"github.com/Knetic/govaluate"
	"quinn007.com/procedures"
	"quinn007.com/variables"
)

var (
	rootSymTable *SymTable
	curSymTable  *SymTable
)

type LogicSymbol string

const (
	LogicSymbolError LogicSymbol = "logicSymbolError"
	LogicSymbolIf                = "if"
	LogicSymbolElse              = "else"
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
}

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
			prev.ParseIfElseClause()
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
func (this *SymTable) ParseIfElseClause() {
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
	newEntryTable := &SymTable{}
	SetCurSymTable(newEntryTable)
	SetRootSymTale(newEntryTable)
	return newEntryTable
}

func NewSymTable(prevSymTable *SymTable) *SymTable {
	var newSymTable *SymTable
	newSymTable = &SymTable{
		prev:        prevSymTable,
		variableMap: make(map[string]*variables.Variable),
		functions:   make([]*procedures.FFunction, 0),
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
