package handlerV2

import (
	"strconv"

	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/variables"
)

func LiteralContextHandler(contextParser *parser.LiteralContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.BasicLitContext:
			{
				BasicLitContextHandler(parserContext, scanner)
			}

		}
	}

	return nil
}

func BasicLitContextHandler(contextParser *parser.BasicLitContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.String_Context:
			{
				StringContextHandler(parserContext, scanner)
			}
		case *parser.IntegerContext:
			{
				IntegerContextHandler(parserContext, scanner)
			}
		}
	}

	return nil
}

func IntegerContextHandler(contextParser *parser.IntegerContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		cursor, _ := navigator.GetCursor()
		terminalString, _ := utils.GetTerminalNodeText(child)
		curStatement := cursor.GetStatement()
		intVal, _ := strconv.Atoi(terminalString)
		curVariable := variables.NewVariable(
			"",
			variables.VTypeInt,
			intVal,
			cursor.GetIndex())

		//cursor.IncreaseIndex()
		curSymTable := sym_tables.GetCurSymTable()

		if scanner.GetInnerType() == consts.ICTypeFuncArgs {
			curFunction := curSymTable.GetLastFunction()
			curFunction.AddParam(curVariable)

		} else if scanner.GetInnerType() == consts.ICTypeLambdaExpr {
			scanner.AppendLambdaExpr(terminalString)
		} else if scanner.GetInnerType() == consts.ICTypeLambdaIfExpr {
			scanner.AppendLambdaExprList(terminalString)
			//lambdaIfElseContext := scanner.GetLambdaIfElseClause()
			//lambdaIfElseContext.AppendIfExpr(terminalString)
		} else if scanner.GetInnerType() == consts.ICTypeLambdaCondition {
			lambdaContext := scanner.GetLambdaContext()
			lambdaContext.AppendSubExpr(terminalString)
		} else {
			curStatement.AddRightValue(curVariable)
		}

		if scanner.GetMiddleType() == consts.MCTypeExpr {
			scanner.AppendExpr(terminalString)
		}
	}

	return nil
}

func StringContextHandler(contextParser *parser.String_Context, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		cursor, _ := navigator.GetCursor()
		curStatement := cursor.GetStatement()
		terminalString, _ := utils.GetTerminalNodeText(child)
		// strip out quotes
		terminalString = terminalString[1 : len(terminalString)-1]
		curVariable := variables.NewVariable(
			"",
			variables.VTypeString,
			terminalString,
			cursor.GetIndex())
		curSymTable := sym_tables.GetCurSymTable()

		if scanner.GetInnerType() == consts.ICTypeFuncArgs {
			curFunction := curSymTable.GetLastFunction()
			curFunction.AddParam(curVariable)
		} else if scanner.GetInnerType() == consts.ICTypeLambdaExpr {
			scanner.AppendLambdaExpr(terminalString)
		} else if scanner.GetInnerType() == consts.ICTypeLambdaCondition {
			// scanner.AppendLambdaExprList(terminalString)
			lambdaContext := scanner.GetLambdaContext()
			lambdaContext.AppendSubExpr(terminalString)
			//lambdaIfElseContext := scanner.GetLambdaIfElseClause()
			//lambdaIfElseContext.AppendIfExpr(terminalString)
		} else {
			curStatement.AddRightValue(curVariable)
		}

		if scanner.GetMiddleType() == consts.MCTypeExpr {
			scanner.AppendExpr(terminalString)
		}

	}

	return nil
}
