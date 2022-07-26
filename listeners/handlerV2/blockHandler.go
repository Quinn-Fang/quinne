package handlerV2

import (
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/navigator/utils"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func BlockContextHandler(antlrCtx antlr.ParserRuleContext, blockContext *sym_tables.ScopeContext, scanner *scanner.Scanner) {
	EnterBlockHandler(blockContext, scanner)
	children := antlrCtx.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.StatementListContext:
			{
				StatementListHandler(parserContext, scanner)
			}
		}
	}
	ExitBlockHandler(scanner)
}

func EnterBlockHandler(blockContext *sym_tables.ScopeContext, scanner *scanner.Scanner) {
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

func ExitBlockHandler(scanner *scanner.Scanner) {
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
