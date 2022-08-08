package handler

import (
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
)

func SimpleStmtContextHandler(contextParser *parser.SimpleStmtContext) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.NewStatement()
	curStatement := curCursor.GetStatement()
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ShortVarDeclContext:
			{
				ShortVarDeclContextHandler(parserContext)
			}
		case *parser.VarSpecContext:
			{
				VarSpecContextHandler(parserContext)
			}
		case *parser.ExpressionStmtContext:
			{
				ExpressionStmtContextHandler(parserContext)
			}
		}
	}
	//fmt.Println("##########################################################")
	//left_values := curStatement.GetLeftValues()
	rightValues := curStatement.GetRightValues()
	//fmt.Println(left_values)
	//// fmt.Println(rightValues)
	//fmt.Println("--------------------   Separator -------------------------")
	//for _, rightValue := range rightValues {
	//	fmt.Printf("%+v ", *rightValue)
	//}
	//fmt.Println()
	//fmt.Println("--------------------   Separator -------------------------")
	curSymTable := sym_tables.GetCurSymTable()
	//curFunctions := curSymTable.GetFunctions()
	//for _, curFunction := range curFunctions {
	//	fmt.Printf("%+v ", curFunction)
	//}
	//fmt.Println()

	//fmt.Println("##########################################################")

	// assign values
	curStatement.Assign()

	// add assigned variable to symbol table
	for _, variable := range rightValues {
		curSymTable.AddVariable(variable)
	}

	return nil
}
