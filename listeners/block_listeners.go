package listeners

import (
	"fmt"

	"quinn007.com/navigator"
	"quinn007.com/navigator/utils"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
)

func (this *GoListener) EnterBlock(c *parser.BlockContext) {
	fmt.Println("ENTERING BLOCK $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	newSymTable := sym_tables.NewSymTable(sym_tables.GetCurSymTable())
	// Navigator start

	curNavigator := navigator.GetCurNavigator()
	symTableCursorStack := curNavigator.GetSymTableCursorStack()

	fmt.Println("11111111111111111111")
	fmt.Printf("%p\n", sym_tables.GetCurSymTable())
	sym_tables.GetCurSymTable().PrintFunctions()
	for _, vvv := range symTableCursorStack.GetStack() {
		fmt.Printf("%p %+v ", vvv.GetSymTable(), vvv)
	}
	fmt.Println()

	curSymTableCursor, err := symTableCursorStack.Peek()
	if err == nil {
		curSymTableCursor.SetFuncEndIndex(
			len(sym_tables.GetCurSymTable().GetFunctions()) - 1)
		curNavigator.GetCodeSegment().InsertBack(curSymTableCursor)
	}
	newSymTableCursor := utils.NewSymTableCursor()
	newSymTableCursor.SetSymTable(newSymTable)

	symTableCursorStack.Push(newSymTableCursor)
	// Navigator end

	sym_tables.SetCurSymTable(newSymTable)
}

func (this *GoListener) ExitBlock(c *parser.BlockContext) {
	fmt.Println("EXITING BLOCK $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	curSymTable := sym_tables.GetCurSymTable()
	// Navigator start
	// 1: Extend the last code block
	// 2: Create new code block in case more code

	curNavigator := navigator.GetCurNavigator()
	symTableCursorStack := curNavigator.GetSymTableCursorStack()
	codeSegment := curNavigator.GetCodeSegment()

	fmt.Println("2222222222222222222")
	fmt.Printf("%p\n", sym_tables.GetCurSymTable())
	sym_tables.GetCurSymTable().PrintFunctions()
	for _, vvv := range symTableCursorStack.GetStack() {
		fmt.Printf("%p %+v ", vvv.GetSymTable(), vvv)
	}
	fmt.Println()
	fmt.Println()

	if curSymTableCursor, err := symTableCursorStack.Peek(); err != nil {
		panic("Unknown err stack empty")
	} else {
		// curSymTableCursor.SetFuncEndIndex(curSymTableCursor.GetFuncStartIndex() + len(curSymTable.GetFunctions()) - 1)
		// curSymTableCursor.SetFuncEndIndex(curSymTableCursor.GetFuncStartIndex() + len(curSymTable.GetFunctions()) - 1)
		newEnd := len(curSymTable.GetFunctions()) - 1

		fmt.Println("------------------------------")
		fmt.Printf("%+v %+v\n\n", curSymTableCursor, newEnd)

		curSymTableCursor.SetFuncEndIndex(newEnd)
		codeSegment.InsertBack(curSymTableCursor)
		_, err_2 := symTableCursorStack.Pop()
		if err_2 != nil {
			panic("err_2")
		}
		curSymTableCursor, _ := symTableCursorStack.Peek()
		curSymTable.PrintFunctions()

		newSymTableCursor := utils.NewSymTableCursor()
		newSymTableCursor.SetSymTable(curSymTable.GetPrev())
		newSymTableCursor.SetFuncStartIndex(curSymTableCursor.GetFuncEndIndex() + 1)
		symTableCursorStack.Push(newSymTableCursor)
	}

	// Navigator end
	sym_tables.SetCurSymTable(curSymTable.GetPrev())
}
