package sym_tables

import (
	"errors"

	"quinn007.com/procedures"
	"quinn007.com/variables"
)

var (
	curSymTable *SymTable
)

type SymTable struct {
	prev        *SymTable
	children    []*SymTable
	variableMap map[string]variables.Variable
	functions   []*procedures.FFunction
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
	return newEntryTable
}

func NewSymTable(prevSymTable *SymTable) *SymTable {
	var newSymTable *SymTable
	newSymTable = &SymTable{
		prev: prevSymTable,
	}
	// prevSymTable.next = newSymTable
	prevSymTable.children = append(prevSymTable.children, newSymTable)

	return newSymTable
}

func (this *SymTable) AddVariable(newVariable *variables.Variable) error {
	if _, ok := this.variableMap[newVariable.GetVariableName()]; ok {
		return errors.New("Variable exists")
	} else {
		this.variableMap[newVariable.GetVariableName()] = *newVariable
		return nil
	}
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
