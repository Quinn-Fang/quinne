package handler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/listeners/utils"
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
	"quinn007.com/uspace"
)

func IfElseStmtContextHandler(contextParser *parser.IfStmtContext) error {
	fmt.Println("Inside IfElseStmtContextHandler .........................")

	curNavigator := navigator.GetCurNavigator()

	curCursor, _ := navigator.GetCursor()
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.BlockContext:
			{
				// cursorContext := curCursor.GetCursorContext()

				// Set up current block context without creating new sym table
				curSymTable := sym_tables.GetCurSymTable()
				if curSymTable.IfElseStackEmpty() {
					panic("ifElseStack Empty")
				}
				ifElseClause := curSymTable.GetLastIfElseClause()

				// check stack see what event are we facing here
				firstSymbol := curSymTable.PopBackIfElseStack()
				var curEvent sym_tables.ContextType
				if firstSymbol == sym_tables.LogicSymbolIf {

					if curSymTable.IfElseStackEmpty() {
						// if event
						newIfBranch := sym_tables.NewIfElseBranch(sym_tables.BranchTypeIf)
						newIfBranch.SetExpr(curCursor.GetExpr())
						newIfBranch.SetExprVarNames(curCursor.GetExprVarNames())

						ifElseClause.AddBranch(newIfBranch)
						// add user space event
						curNavigator.AddEvent(uspace.EventTypeIfElseExpr, newIfBranch)

						curEvent = sym_tables.ContextTypeIf
					} else {
						// else-if event
						secondSymbol := curSymTable.PopBackIfElseStack()
						if secondSymbol != sym_tables.LogicSymbolElse {
							panic("else not followed by if")
						}
						newElseIfBranch := sym_tables.NewIfElseBranch(sym_tables.BranchTypeElseIf)
						newElseIfBranch.SetExpr(curCursor.GetExpr())
						newElseIfBranch.SetExprVarNames(curCursor.GetExprVarNames())

						ifElseClause.AddBranch(newElseIfBranch)
						// add user space event
						curNavigator.AddEvent(uspace.EventTypeIfElseExpr, newElseIfBranch)

						curEvent = sym_tables.ContextTypeElseIf
					}
				} else if firstSymbol == sym_tables.LogicSymbolElse {
					// else event
					newElseBranch := sym_tables.NewIfElseBranch(sym_tables.BranchTypeElse)

					ifElseClause.AddBranch(newElseBranch)

					curEvent = sym_tables.ContextTypeElse

				} else {
					panic("Unknown error")
				}

				blockContext := sym_tables.NewScopeContext(curEvent)
				curCursor.InitIfElseClause()
				curCursor.ClearExpr()
				curCursor.InitExprVarNames()
				BlockContextHandler(parserContext, blockContext)
			}
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext)
			}
		case *parser.IfStmtContext:
			{
				IfElseStmtContextHandler(parserContext)
			}
		case *antlr.TerminalNodeImpl:
			{
				curSymTable := sym_tables.GetCurSymTable()

				terminalString, _ := utils.GetTerminalNodeText(child)
				if terminalString == string(sym_tables.LogicSymbolIf) {
					fmt.Println("Context IF ........................")
					if curSymTable.IfElseStackEmpty() {
						// if event
						// create if-else clause
						newIfElseClause := sym_tables.NewIfElseClause()
						curSymTable.AddIfElseClause(newIfElseClause)
						curCursor.SetIfElseClause(newIfElseClause)
					}

					curCursor.SetCursorContext(sym_tables.ContextTypeIf)

					curSymTable.PushIfElseStack(sym_tables.LogicSymbolIf)
				} else if terminalString == string(sym_tables.LogicSymbolElse) {
					fmt.Println("Context ELSE ......................")
					curSymTable.PushIfElseStack(sym_tables.LogicSymbolElse)
				}
			}
		}
	}
	fmt.Println("Exiting IfElseStmtContextHandler .........................")
	curSymTable := sym_tables.GetCurSymTable()
	curSymTable.PrintFunctions()
	curSymTable.PrintIfElseClauseList()
	return nil
}
