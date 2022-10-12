package handlerV2

import (
	"strconv"

	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/scanner/consts"
	scannerConsts "github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/variables"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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
				if curCursor.IsAppendingExpr() {
					terminalString, _ := utils.GetTerminalNodeText(parserContext)

					curCursor.PushExpr(terminalString)
				}

				if scanner.GetMiddleType() == scannerConsts.MCTypeExpr {
					terminalString, _ := utils.GetTerminalNodeText(parserContext)
					scanner.AppendExpr(terminalString)
				}
			}
		case *parser.LambdaContext:
			{
				//fmt.Println("1111")
				//terminalString, _ := utils.GetTerminalNodeText(parserContext)
				//fmt.Println(terminalString)

				///////////////////////////
				//x, _ := utils.GetTerminalNodeText(parserContext.GetChildren()[0])
				//x, _ = utils.GetTerminalNodeText(parserContext.GetChildren()[1])
				//x, _ = utils.GetTerminalNodeText(parserContext.GetChildren()[2])
				//x, _ = utils.GetTerminalNodeText(parserContext.GetChildren()[3])
				//x, _ = utils.GetTerminalNodeText(parserContext.GetChildren()[4])
				//fmt.Println(x)

				/////////////////////////////
				parserContextChildren := parserContext.GetChildren()
				if len(parserContextChildren) == 4 {
					// lambda expression without if else statement
					LambdaHandler(parserContextChildren[1].(*parser.VarSpecListContext), parserContextChildren[3].(*parser.ExpressionListContext),
						nil, scanner)
				} else {
					// lambda expression with if else statement
					LambdaHandler(parserContextChildren[1].(*parser.VarSpecListContext), parserContextChildren[3].(*parser.ExpressionListContext),
						parserContextChildren[4].(*parser.LambdaIfStmtContext), scanner)
				}

				cursor, _ := navigator.GetCursor()
				terminalString, _ := utils.GetTerminalNodeText(child)
				curStatement := cursor.GetStatement()
				intVal, _ := strconv.Atoi(terminalString)
				curVariable := variables.NewVariable(
					"",
					variables.VTypeFunctionDecl,
					intVal,
					cursor.GetIndex())

				//cursor.IncreaseIndex()
				curSymTable := sym_tables.GetCurSymTable()

				if scanner.GetInnerType() == consts.ICTypeFuncArgs {
					curFunction := curSymTable.GetLastFunction()
					curFunction.AddParam(curVariable)

				} else {
					curStatement.AddRightValue(curVariable)
				}

				if scanner.GetMiddleType() == consts.MCTypeExpr {
					scanner.AppendExpr(terminalString)
				}

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
