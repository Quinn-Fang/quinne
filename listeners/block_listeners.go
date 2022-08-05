package listeners

import (
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
)

func (this *GoListener) EnterBlock(c *parser.BlockContext) {
	newSymTable := sym_tables.NewSymTable(sym_tables.GetCurSymTable())
	sym_tables.SetCurSymTable(newSymTable)
}

func (this *GoListener) ExitBlock(c *parser.BlockContext) {
	curSymTable := sym_tables.GetCurSymTable()
	sym_tables.SetCurSymTable(curSymTable.GetPrev())
}
