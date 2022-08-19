package handler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/navigator/utils"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/sym_tables"
)

func BlockContextHandler(antlrCtx antlr.ParserRuleContext, blockContext *sym_tables.ScopeContext) {
	EnterBlockHandler(blockContext)
	children := antlrCtx.GetChildren()

	for _, child := range children {
		//fmt.Println()
		//fmt.Printf("%T\n", child)
		//fmt.Println()
		switch parserContext := child.(type) {
		case *parser.StatementListContext:
			{
				StatementListHandler(parserContext)
			}
		}
	}
	ExitBlockHandler()

	/////////////////////////
	//curSymTable := sym_tables.GetCurSymTable()
	//if curSymTable != nil {
	//	curSymTable.PrintIfElseClauseList()
	//}
}

func EnterBlockHandler(blockContext *sym_tables.ScopeContext) {
	// fmt.Println("ENTERING BLOCK $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	newSymTable := sym_tables.NewSymTable(sym_tables.GetCurSymTable())
	// Navigator start
	newSymTable.SetScope(blockContext)

	curNavigator := navigator.GetCurNavigator()
	symTableCursorStack := curNavigator.GetSymTableCursorStack()

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

	// Setting Executable

}

func ExitBlockHandler() {
	// 	fmt.Println("EXITING BLOCK $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	curSymTable := sym_tables.GetCurSymTable()
	// Navigator start
	// 1: Extend the last code block
	// 2: Create new code block in case more code

	curNavigator := navigator.GetCurNavigator()
	symTableCursorStack := curNavigator.GetSymTableCursorStack()
	codeSegment := curNavigator.GetCodeSegment()

	if curSymTableCursor, err := symTableCursorStack.Peek(); err != nil {
		panic("Unknown err stack empty")
	} else {
		newEnd := len(curSymTable.GetFunctions()) - 1

		curSymTableCursor.SetFuncEndIndex(newEnd)
		codeSegment.InsertBack(curSymTableCursor)

		_, err_2 := symTableCursorStack.Pop()
		if err_2 != nil {
			panic("err_2")
		}
		curSymTableCursor, _ := symTableCursorStack.Peek()
		// curSymTable.PrintFunctions()

		newSymTableCursor := utils.NewSymTableCursor()
		newSymTableCursor.SetSymTable(curSymTable.GetPrev())
		newSymTableCursor.SetFuncStartIndex(curSymTableCursor.GetFuncEndIndex() + 1)
		symTableCursorStack.Push(newSymTableCursor)
	}

	// Navigator end
	sym_tables.SetCurSymTable(curSymTable.GetPrev())
}
