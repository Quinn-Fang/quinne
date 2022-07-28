package navigator

import (
	"quinn007.com/procedures"
	"quinn007.com/sym_tables"
)

type Navigator struct {
	curFunction *procedures.FFunction
	curSymTable *sym_tables.SymTable
}

func NewNavigator() *Navigator {
	newNavigator := &Navigator{}
	return newNavigator
}

func (this *Navigator) GetCurFunction() *procedures.FFunction {
	return this.curFunction
}

func (this *Navigator) GetCurTable() *sym_tables.SymTable {
	return this.curSymTable
}
