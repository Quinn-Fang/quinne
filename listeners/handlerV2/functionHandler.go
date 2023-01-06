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

	// Here wether it is a function call or a lambda call is determined in
	// the operand name handler, by looking up it's definition in current
	// symbol table.
	if scanner.GetMiddleType() == consts.MCTypeExpr {
		scanner.AppendExpr("(")
	}
	ArgumentsContextHandler(argumentsContext, scanner)
	if scanner.GetMiddleType() == consts.MCTypeExpr {
		scanner.AppendExpr(")")
	}

	curSymTable := sym_tables.GetCurSymTable()
	if !(scanner.GetInnerType() == consts.ICTypeLambdaCall) {
		fFunction := curSymTable.GetLastFunction()
		curNavigator := navigator.GetCurNavigator()
		curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)
	}
	//curSymTable := sym_tables.GetCurSymTable()
	//var fFunction *procedures.FFunction

	//if !(scanner.GetInnerType() == consts.ICTypeLambdaCall) {
	//	fFunction = curSymTable.GetLastFunction()
	//}
	//curCursor.SetCursorContext(sym_tables.ContextTypeDefault)

	//curNavigator := navigator.GetCurNavigator()

	//if !(scanner.GetInnerType() == consts.ICTypeLambdaCall) {
	//	if !curCursor.IsAppendingExpr() {
	//		// regular function
	//		curNavigator.AddEvent(uspace.EventTypeFunction, fFunction, curSymTable)
	//	} else {
	//		curCursor.PushExpr(fFunction.ToString())
	//		curCursor.AddExprVarNames(fFunction.FName)
	//	}

	//}

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
