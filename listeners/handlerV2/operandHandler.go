package handlerV2

import (
	"fmt"

	"github.com/Quinn-Fang/quinne/listeners/utils"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/procedures"
	scannerPkg "github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/uspace"
	"github.com/Quinn-Fang/quinne/variables"
)

func OperandContextHandler(contextParser *parser.OperandContext, scanner *scannerPkg.Scanner) error {
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

func OperandNameContextHandler(contextParser *parser.OperandNameContext, scanner *scannerPkg.Scanner) error {
	children := contextParser.GetChildren()
	terminalString, _ := utils.GetTerminalNodeText(children[0])
	curCursor, _ := navigator.GetCursor()
	curSymTable := sym_tables.GetCurSymTable()
	curStatement := curCursor.GetStatement()

	if scanner.GetInnerType() == consts.ICTypeFuncName {
		symTable := sym_tables.GetCurSymTable()
		// check if function has been defined
		if variable, err := symTable.GetVariableByName(terminalString); err == nil {
			// check if it is lambda call
			if variable.GetVariableType() == variables.VTypeLambdaFunctionDecl {
				var emptyValue interface{}
				newReturnValue := variables.NewVariable(
					"",
					variables.VTypeLambdaReturned,
					emptyValue,
					-1,
				)
				lambdaDecl, _ := variable.GetVariableValue().(*procedures.LambdaDecl)
				lambadCall := procedures.NewLambdaCall(lambdaDecl)
				lambadCall.SetReturnValue(newReturnValue)
				curNavigator := navigator.GetCurNavigator()
				curNavigator.AddEvent(uspace.EventTypeLambdaCall, lambadCall, curSymTable)
				curStatement.AddRightValue(newReturnValue)

				scanner.NewInnerContext(consts.ICTypeLambdaCall)
				lambdaCallContext, _ := scanner.GetInnerContext().(*scannerPkg.LambdaCallContext)
				lambdaCallContext.SetLambdaCall(lambadCall)
				lambdaCallContext.SetRetValue(newReturnValue)
			}
		} else {
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
			scanner.SetInnerType(consts.ICTypeFuncArgs)

		}
	} else if scanner.GetInnerType() == consts.ICTypeFuncArgs {
		variable, err := curSymTable.GetVariableByName(terminalString)
		if err != nil {
			errMsg := fmt.Sprintf("variable: %s does not exist", terminalString)
			panic(errMsg)
		}

		curFunction := curSymTable.GetLastFunction()
		curFunction.AddParam(variable)
	} else if scanner.GetInnerType() == consts.ICTypeLambdaExpr {
		scanner.AppendLambdaExprList(terminalString)
	} else if scanner.GetInnerType() == consts.ICTypeLambdaIfExpr {
		//scanner.AppendLambdaExprList(terminalString)
		//curLambdaIfClauseCtx := scanner.GetLambdaIfElseClause()
		//curLambdaIfClauseCtx.AppendIfExpr(terminalString)
	} else if scanner.GetInnerType() == consts.ICTypeLambdaCondition {
		lambdaContext := scanner.GetLambdaContext()
		lambdaContext.AppendSubExpr(terminalString)
		//scanner.AppendLambdaExprList(terminalString)
		//lambdaIfElseContext := scanner.GetLambdaIfElseClause()
		//lambdaIfElseContext.AppendIfExpr(terminalString)
	} else if scanner.GetInnerType() == consts.ICTypeLambdaRet {
		scanner.SetLambdaReturnValue(terminalString)
	} else if scanner.GetInnerType() == consts.ICTypeLambdaCall {
		variable, err := curSymTable.GetVariableByName(terminalString)
		if err != nil {
			errMsg := fmt.Sprintf("variable: %s does not exist", terminalString)
			panic(errMsg)
		}
		lambdaCallContext := scanner.GetInnerContext().(*scannerPkg.LambdaCallContext)
		lambdaCallContext.AddArgs(variable)
	}

	if scanner.GetMiddleType() == consts.MCTypeExpr {
		scanner.AppendExpr(terminalString)
		scanner.AppendExprVarName(terminalString)
	}

	return nil
}
