package handler

import (
	"fmt"

	"quinn007.com/parser"
)

func SimpleStmtContextHandler(contextParser *parser.SimpleStmtContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		fmt.Println("............................................")
		fmt.Printf("%T\n", child)
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

	return nil
}
