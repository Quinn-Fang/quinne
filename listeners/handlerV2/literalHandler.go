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

		if scanner.GetMiddleType() == consts.MCTypeExpr {
			scanner.AppendExpr(terminalString)
		} else if scanner.GetInnerType() == consts.ICTypeFuncArgs {
			curFunction := curSymTable.GetLastFunction()
			curFunction.AddParam(curVariable)

		} else {
			curStatement.AddRightValue(curVariable)

		}

		/////////////////////////////////////// should be removed ////////////////////////////////////////
		//if cursor.GetCursorContext() == sym_tables.ContextTypeFunctionArgs {
		//	curFunction := curSymTable.GetLastFunction()
		//	curFunction.AddParam(curVariable)

		//	// curStatement.AddRightValue(curVariable)
		//	// cursor.PrintStatement()
		//	// } else if cursor.GetCursorContext() == sym_tables.ContextTypeIf || cursor.GetCursorContext() == sym_tables.ContextTypeElseIf {
		//} else if cursor.IsAppendingExpr() {
		//	cursor.PushExpr(terminalString)
		//} else {
		//	curStatement.AddRightValue(curVariable)

		//}
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

		if scanner.GetMiddleType() == consts.MCTypeExpr {
			scanner.AppendExpr(terminalString)
		} else if scanner.GetInnerType() == consts.ICTypeFuncArgs {
			curFunction := curSymTable.GetLastFunction()
			curFunction.AddParam(curVariable)
		} else {
			curStatement.AddRightValue(curVariable)
		}

		/////////////////////////////////////// should be removed ////////////////////////////////////////
		//if cursor.GetCursorContext() == sym_tables.ContextTypeFunctionArgs {
		//	curFunction := curSymTable.GetLastFunction()
		//	curFunction.AddParam(curVariable)

		//	// } else if cursor.GetCursorContext() == sym_tables.ContextTypeIf || cursor.GetCursorContext() == sym_tables.ContextTypeElseIf {
		//} else if cursor.IsAppendingExpr() {
		//	cursor.PushExpr(terminalString)
		//} else {
		//	curStatement.AddRightValue(curVariable)
		//}

	}

	return nil
}
