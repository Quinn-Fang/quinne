package handlerV2

import (
	"github.com/Quinn-Fang/quinne/listeners/handler"
	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	scannerPkg "github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/scanner/consts"
	scannerConsts "github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func IfElseStmtContextHandler(contextParser *parser.IfStmtContext, scanner *scannerPkg.Scanner) error {
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
				// set middle context to append expr for both
				// if and else-if event and empty middle context for else
				scanner.NewMiddleContext(scannerConsts.MCTypeExpr)
				if terminalString == string(sym_tables.LogicSymbolIf) {
					if scanner.GetOuterType() == scannerConsts.OCTypeElse {
						// else-if event, the previous else event has already created
						// the outer context, so here only change the context type
						scanner.SetOuterType(scannerConsts.OCTypeElseIf)
					} else {
						// if event
						scanner.NewOuterContext(scannerConsts.OCTypeIf)
						scanner.AddIfElseEvent(sym_tables.LogicSymbolIf)
						// Create a new if-else clause and add it to current symbol table
						newIfElseClause := sym_tables.NewIfElseClause()
						curSymTable := sym_tables.GetCurSymTable()
						curSymTable.AddIfElseClause(newIfElseClause)
						scanner.SetOuterType(scannerConsts.OCTypeIf)
					}
				} else if terminalString == string(sym_tables.LogicSymbolElse) {
					// else or else-if event
					scanner.NewOuterContext(scannerConsts.OCTypeElse)
					scanner.AddIfElseEvent(sym_tables.LogicSymbolElse)
				}
			}
		}
	}

	return nil
}

func LambdaIfElseStmtContextHandler(contextParser *parser.LambdaIfStmtContext, scanner *scannerPkg.Scanner) error {
	// curCursor, _ := navigator.GetCursor()
	children := contextParser.GetChildren()
	scanner.SetInnerType(consts.ICTypeLambdaIfClause)
	newLambdaIfClauseCtx := scanner.NewLambdaIfElseClauseContext()
	oldLambdaIfClauseCtx := scanner.GetLambdaIfElseClause()
	if oldLambdaIfClauseCtx != nil {
		oldLambdaIfClauseCtx.AddElseClause(oldLambdaIfClauseCtx)
	} else {
		newLambdaIfClauseCtx.SetBranchSymbol(sym_tables.LogicSymbolIf)
		scanner.SetLambdaIfElseClauseContext(newLambdaIfClauseCtx)
		scanner.SetLambdaIfElseClauseContextEntry(newLambdaIfClauseCtx)
	}
	scanner.SetInnerType(consts.ICTypeLambdaIfClause)

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ExpressionContext:
			{
				// scanner.SetInnerType(consts.ICTypeLambdaIfExpr)
				scanner.SetInnerType(consts.ICTypeLambdaCondition)
				ExpressionContextHandler(parserContext, scanner)
				lambdaContext, _ := scanner.GetInnerContext().(*scannerPkg.LambdaContext)
				lambdaContext.AppendExprList(lambdaContext.GetSubExpr())
				lambdaContext.ClearSubExpr()
				scanner.SetInnerType(consts.ICTypeLambdaIfClause)
			}
		case *parser.LambdaIfStmtContext:
			{
				LambdaIfElseStmtContextHandler(parserContext, scanner)
			}
		case *antlr.TerminalNodeImpl:
			{
				terminalString, _ := utils.GetTerminalNodeText(child)
				if scanner.GetInnerType() == consts.ICTypeLambdaIfClause {
					scanner.AppendLambdaExprList(terminalString)
				}
			}
		}
	}

	return nil
}
