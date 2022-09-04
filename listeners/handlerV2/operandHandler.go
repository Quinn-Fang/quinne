package handlerV2

import (
	"fmt"

	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/variables"
)

func OperandContextHandler(contextParser *parser.OperandContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

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
		case *parser.OperandNameContext:
			{
				OperandNameContextHandler(parserContext, scanner)
			}

		}
	}

	return nil
}

func OperandNameContextHandler(contextParser *parser.OperandNameContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()
	terminalString, _ := utils.GetTerminalNodeText(children[0])
	curCursor, _ := navigator.GetCursor()
	curSymTable := sym_tables.GetCurSymTable()
	curStatement := curCursor.GetStatement()

	if scanner.GetMiddleType() == consts.MCTypeExpr {
		scanner.AppendExpr(terminalString)
		scanner.AppendExprVarName(terminalString)
	} else if scanner.GetInnerType() == consts.ICTypeFuncName {
		var emptyValue interface{}
		newReturnValue := variables.NewVariable(
			"",
			variables.VTypeFunctionReturned,
			emptyValue,
			curCursor.GetIndex())

		curCursor.IncreaseIndex()

		curStatement.AddRightValue(newReturnValue)

		newFunction := procedures.NewFunction(terminalString)
		newFunction.InitReturnValue(newReturnValue)
		curSymTable.AddFunction(newFunction)
	} else if scanner.GetInnerType() == consts.ICTypeFuncArgs {
		variable, err := curSymTable.GetVariableByName(terminalString)
		if err != nil {
			errMsg := fmt.Sprintf("variable: %s does not exist", terminalString)
			panic(errMsg)
		}

		curFunction := curSymTable.GetLastFunction()
		curFunction.AddParam(variable)
	}

	//if curCursor.GetCursorContext() == sym_tables.ContextTypeFunctionName {

	//	var emptyValue interface{}
	//	newReturnValue := variables.NewVariable(
	//		"",
	//		variables.VTypeFunctionReturned,
	//		emptyValue,
	//		curCursor.GetIndex())

	//	curCursor.IncreaseIndex()

	//	curStatement.AddRightValue(newReturnValue)

	//	newFunction := procedures.NewFunction(terminalString)
	//	newFunction.InitReturnValue(newReturnValue)
	//	curSymTable.AddFunction(newFunction)
	//} else if curCursor.GetCursorContext() == sym_tables.ContextTypeFunctionArgs {
	//	variable, err := curSymTable.GetVariableByName(terminalString)
	//	if err != nil {
	//		errMsg := fmt.Sprintf("variable: %s does not exist", terminalString)
	//		panic(errMsg)
	//	}

	//	curFunction := curSymTable.GetLastFunction()
	//	curFunction.AddParam(variable)
	//} else if curCursor.IsAppendingExpr() {
	//	curCursor.PushExpr(terminalString)
	//	curCursor.AddExprVarNames(terminalString)
	//}

	return nil
}
