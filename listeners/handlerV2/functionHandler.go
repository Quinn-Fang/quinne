package handlerV2

import (
	"fmt"

	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/Quinn-Fang/quinne/scanner/consts"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/uspace"
)

func LambdaHandler(varSpecList *parser.VarSpecListContext, expressionListContext *parser.ExpressionListContext,
	lambdaIfStmt *parser.LambdaIfStmtContext, scanner *scanner.Scanner) error {
	scanner.NewInnerContext(consts.ICTypeLambdaParams)
	VarSpecListContextHandler(varSpecList, scanner)
	scanner.SetInnerType(consts.ICTypeLambdaExpr)
	scanner.AppendLambdaExpr(":")
	scanner.SetInnerType(consts.ICTypeLambdaRet)
	ExpressionListContextHandler(expressionListContext, scanner)

	// Optional if else statement of lambda statement
	if lambdaIfStmt != nil {
		scanner.SetInnerType(consts.ICTypeLambdaIfClause)
		LambdaIfElseStmtContextHandler(lambdaIfStmt, scanner)
	}
	y := scanner.GetLambdaContext().ToTernaryExpr()
	fmt.Println(y)
	entry := scanner.GetLambdaIfElseClauseEntry()
	x := entry.ToExprList()
	fmt.Println(x)
	return nil
}

func FunctionHandler(operandContext *parser.PrimaryExprContext, argumentsContext *parser.ArgumentsContext, scanner *scanner.Scanner) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionName)
	// use scanner instead
	scanner.NewInnerContext(consts.ICTypeFuncName)
	// scanner.SetInnerType(consts.ICTypeFuncName)

	PrimaryExprContextHandler(operandContext, scanner)
	// scanner.SetInnerType(consts.ICTypeFuncArgs)

	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionArgs)

	// use scanner instead
	// already set in function name (operandName handler)
	//scanner.SetInnerType(consts.ICTypeFuncArgs)

	if scanner.GetMiddleType() == consts.MCTypeExpr {
		scanner.AppendExpr("(")
	}
	ArgumentsContextHandler(argumentsContext, scanner)
	if scanner.GetMiddleType() == consts.MCTypeExpr {
		scanner.AppendExpr(")")
	}
	curSymTable := sym_tables.GetCurSymTable()
	var fFunction *procedures.FFunction
	if curCursor.GetCursorContext() == sym_tables.ContextTypeFunctionArgs {
		if !(scanner.GetInnerType() == consts.ICTypeLambdaCall) {
			fFunction = curSymTable.GetLastFunction()
		}
	}

	curCursor.SetCursorContext(sym_tables.ContextTypeDefault)

	// fFunction := curSymTable.GetLastFunction()
	curNavigator := navigator.GetCurNavigator()
	////////////////
	//if scanner
	////////////////

	if !(scanner.GetInnerType() == consts.ICTypeLambdaCall) {
		if !curCursor.IsAppendingExpr() {
			// regular function
			curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)
		} else {
			curCursor.PushExpr(fFunction.ToString())
			curCursor.AddExprVarNames(fFunction.FName)
		}

	}
	// curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)

	scanner.SetInnerType(consts.ICTypeUnset)

	return nil
}

func ArgumentsContextHandler(contextParser *parser.ArgumentsContext, scanner *scanner.Scanner) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ExpressionListContext:
			{
				ExpressionListContextHandler(parserContext, scanner)
			}
		}
	}

	return nil
}

func FunctionDeclHandler(contextParser *parser.FunctionDeclContext, scanner *scanner.Scanner) error {

	curNavigator := navigator.GetCurNavigator()
	curSymTable := sym_tables.GetCurSymTable()

	children := contextParser.GetChildren()
	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.BlockContext:
			{
				scopeContext := procedures.NewFunctionDecl("main")
				blockContext := sym_tables.NewScopeContext(sym_tables.ContextTypeFuncDecl)
				blockContext.SetScopeContext(scopeContext)

				curNavigator.AddEvent(uspace.EventTypeFunctionDecl, scopeContext, curSymTable)
				BlockContextHandler(parserContext, blockContext, scanner)
			}
		}
	}

	return nil
}
