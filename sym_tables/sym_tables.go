package sym_tables

import (
	"errors"

	"quinn007.com/procedures"
	"quinn007.com/variables"
)

type SymTable struct {
	prev        *SymTable
	next        *SymTable
	variableMap map[string]variables.Variable
	functions   []*procedures.FFunction
}

func NewSymTable(prevSymTable *SymTable) *SymTable {
	var newSymTable *SymTable
	if prevSymTable == nil {
		newSymTable = &SymTable{}
	} else {
		newSymTable = &SymTable{
			prev: prevSymTable,
		}
		prevSymTable.next = newSymTable
	}

	return newSymTable
}

func (this *SymTable) AddVariable(newVariable *variables.Variable) error {
	if this == nil {
		return errors.New("Nil symbol table assigned ...")
	}

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

func (this *SymTable) GetPrev() *SymTable {
	return this.prev
}

func (this *SymTable) GetNext() *SymTable {
	return this.next
}
