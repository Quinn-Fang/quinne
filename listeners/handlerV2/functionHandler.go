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

func LambdaHandler(varSpecList *parser.VarSpecListContext, expressionListContext *parser.LambdaExpressionContext, scanner *scanner.Scanner) error {
	// For Lambda definition, lambda should be called in the form of a regular function.
	scanner.NewInnerContext(consts.ICTypeLambdaParams)
	VarSpecListContextHandler(varSpecList, scanner)
	scanner.SetInnerType(consts.ICTypeLambdaExpr)

	scanner.SetInnerType(consts.ICTypeLambdaExpression)
	LambdaExpressionContextHandler(expressionListContext, scanner)

	y := scanner.GetLambdaContext().ToTernaryExpr()
	fmt.Println(y)
	return nil
}

func FunctionHandler(operandContext *parser.PrimaryExprContext, argumentsContext *parser.ArgumentsContext, scanner *scanner.Scanner) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.SetCursorContext(sym_tables.ContextTypeFunctionName)
	scanner.NewInnerContext(consts.ICTypeFuncName)

	PrimaryExprContextHandler(operandContext, scanner)

	// curCursor.SetCursorContext(sym_tables.ContextTypeFunctionArgs)
	// scanner.SetInnerType(consts.ICTypeFuncArgs)

	// Here we do not know whether it is function call or lambda call,
	// in operand handler, see if function name has been defined, if so,
	// it a lambda call, otherwise it it a function call. So the function
	// arg or lambdacall inner type should be set there.

	if scanner.GetMiddleType() == consts.MCTypeExpr {
		scanner.AppendExpr("(")
	}
	ArgumentsContextHandler(argumentsContext, scanner)
	if scanner.GetMiddleType() == consts.MCTypeExpr {
		scanner.AppendExpr(")")
	}
	curSymTable := sym_tables.GetCurSymTable()
	var fFunction *procedures.FFunction
	// if curCursor.GetCursorContext() == sym_tables.ContextTypeFunctionArgs {
	//if scanner.GetInnerType() == consts.ICTypeFuncArgs {
	//	if !(scanner.GetInnerType() == consts.ICTypeLambdaCall) {
	//		fFunction = curSymTable.GetLastFunction()
	//	}
	//}

	if !(scanner.GetInnerType() == consts.ICTypeLambdaCall) {
		fFunction = curSymTable.GetLastFunction()
	}
	curCursor.SetCursorContext(sym_tables.ContextTypeDefault)

	curNavigator := navigator.GetCurNavigator()

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
