package sym_tables

import (
	"errors"

	"quinn007.com/variables"
)

type symTable struct {
	prev        *symTable
	next        *symTable
	variableMap map[string]variables.Variable
}

func NewSymTable(prevSymTable *symTable) *symTable {
	var newSymTable *symTable
	if prevSymTable == nil {
		newSymTable = &symTable{}
	} else {
		newSymTable = &symTable{
			prev: prevSymTable,
		}
		prevSymTable.next = newSymTable
	}

	return newSymTable
}

func (this *symTable) InsertSymTable(newVariable *variables.Variable) error {
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
