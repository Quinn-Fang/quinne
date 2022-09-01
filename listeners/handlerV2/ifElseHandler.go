package handlerV2

import (
	"github.com/Quinn-Fang/quinne/listeners/handler"
	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	scannerConsts "github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func IfElseStmtContextHandler(contextParser *parser.IfStmtContext, scanner *scanner.Scanner) error {
	// fmt.Println("Inside IfElseStmtContextHandler .........................")

	curNavigator := navigator.GetCurNavigator()

	curCursor, _ := navigator.GetCursor()
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.BlockContext:
			{

				curEventType, curEventContext := scanner.ConsumeIfElseEvent()
				blockContext := sym_tables.NewScopeContext(curEventType)
				blockContext.SetScopeContext(curEventContext)

				scanner.ClearAllContext()

				//////////////////////////////////  below should be removed ///////////////////////////////////////////

				curCursor.InitIfElseClause()
				curCursor.ClearExpr()
				curCursor.InitExprVarNames()
				handler.BlockContextHandler(parserContext, blockContext)
			}
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext, scanner)
			}
		case *parser.IfStmtContext:
			{
				IfElseStmtContextHandler(parserContext, scanner)
			}
		case *antlr.TerminalNodeImpl:
			{
				terminalString, _ := utils.GetTerminalNodeText(child)
				if terminalString == string(sym_tables.LogicSymbolIf) {
					//if curSymTable.IfElseStackEmpty() {
					//	// if event
					//	// create if-else clause
					//	newIfElseClause := sym_tables.NewIfElseClause()
					//	curSymTable.AddIfElseClause(newIfElseClause)
					//	curCursor.SetIfElseClause(newIfElseClause)
					//}

					//curCursor.SetCursorContext(sym_tables.ContextTypeIf)
					//curCursor.SetAppendingExpr(true)

					//curSymTable.PushIfElseStack(sym_tables.LogicSymbolIf)

					scanner.AddIfElseEvent(sym_tables.LogicSymbolIf)
					// set middle context to appeding expr status for both
					// if and else-if event
					scanner.SetMiddleContext(scannerConsts.MCTypeExpr)

					if scanner.GetOuterType() == scannerConsts.OCTypeElse {
						// else-if event
						scanner.SetOuterContext(scannerConsts.OCTypeElseIf)
					} else {
						// if event
						// Create a new if-else clause and add it to current symbol table
						newIfElseClause := sym_tables.NewIfElseClause()
						curSymTable := sym_tables.GetCurSymTable()
						curSymTable.AddIfElseClause(newIfElseClause)
						scanner.SetOuterContext(scannerConsts.OCTypeIf)
					}
				} else if terminalString == string(sym_tables.LogicSymbolElse) {
					// else or else-if event
					scanner.AddIfElseEvent(sym_tables.LogicSymbolElse)
					scanner.SetOuterContext(scannerConsts.OCTypeElse)
				}
			}
		}
	}

	return nil
}
