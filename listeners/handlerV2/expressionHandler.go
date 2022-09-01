package handlerV2

import (
	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	scannerConsts "github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/variables"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ExpressionListContextHandler(contextParser *parser.ExpressionListContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	splitter := ","

	ExpressionListStrings := make([]string, 0)
	curCursor, _ := navigator.GetCursor()
	curStatement := curCursor.GetStatement()
	for _, nodeContext := range children {
		switch parserContext := nodeContext.(type) {
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext, scanner)
			}
		default:
			{
				// Splitter

				terminalString, _ := utils.GetTerminalNodeText(nodeContext)
				if terminalString != splitter {
					newVariable := variables.NewVariable(
						"",
						variables.VTypeUndefined,
						terminalString,
						curCursor.GetIndex())

					curStatement.AddRightValue(newVariable)
					curCursor.IncreaseIndex()

					ExpressionListStrings = append(ExpressionListStrings, terminalString)
				}
			}
		}

	}

	return nil
}

func ExpressionContextHandler(contextParser *parser.ExpressionContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	curCursor, _ := navigator.GetCursor()
	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.PrimaryExprContext:
			{
				PrimaryExprContextHandler(parserContext, scanner)
			}
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext, scanner)
			}
		case *antlr.TerminalNodeImpl:
			{
				// if curCursor.GetCursorContext() == sym_tables.ContextTypeIf || curCursor.GetCursorContext() == sym_tables.ContextTypeElseIf {
				if curCursor.IsAppendingExpr() {
					terminalString, _ := utils.GetTerminalNodeText(parserContext)

					curCursor.PushExpr(terminalString)
				}

				if scanner.GetMiddleType() == scannerConsts.MCTypeExpr {
					terminalString, _ := utils.GetTerminalNodeText(parserContext)
					scanner.AppendExpr(terminalString)
				}
				//terminalString, _ := utils.GetTerminalNodeText(parserContext)
				//fmt.Println(terminalString)
			}
		}
	}

	return nil
}

func PrimaryExprContextHandler(contextParser *parser.PrimaryExprContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	if utils.IsFunction(children) {
		FunctionHandler(children[0].(*parser.PrimaryExprContext), children[1].(*parser.ArgumentsContext), scanner)
		return nil
	}

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.IdentifierListContext:
			{
				IdentifierListContextHandler(parserContext, scanner)
			}
		case *parser.LiteralContext:
			{
				LiteralContextHandler(parserContext, scanner)
			}
		case *parser.PrimaryExprContext:
			{
				PrimaryExprContextHandler(parserContext, scanner)
			}
		case *parser.OperandContext:
			{
				OperandContextHandler(parserContext, scanner)
			}
		}
	}

	return nil
}
