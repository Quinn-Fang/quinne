package handler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/listeners/utils"
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
)

func IfElseStmtContextHandler(contextParser *parser.IfStmtContext) error {
	fmt.Println("Inside IfElseStmtContextHandler .........................")
	curCursor, _ := navigator.GetCursor()
	children := contextParser.GetChildren()

	for _, child := range children {
		fmt.Println()
		fmt.Printf("%T\n", child)
		fmt.Printf("%+v\n", child)
		fmt.Println()
		switch parserContext := child.(type) {
		case *parser.BlockContext:
			{

				fmt.Println("Context BLOCK .................")
				curCursor.SetCursorContext(navigator.ContextTypeDefault)
			}
		case *parser.ExpressionContext:
			{
				ExpressionContextHandler(parserContext)
			}
		case *parser.IfStmtContext:
			{
				IfElseStmtContextHandler(parserContext)
			}
		case *antlr.TerminalNodeImpl:
			{
				terminalString, _ := utils.GetTerminalNodeText(child)
				if terminalString == string(sym_tables.LogicSymbolIf) {
					fmt.Println("Context IF ........................")
					//curCursor.PushIfElseStack(navigator.ContextTypeIfBlock)
					curCursor.SetCursorContext(navigator.ContextTypeIfBlock)
				} else if terminalString == string(sym_tables.LogicSymbolElse) {
					fmt.Println("Context ELSE ......................")
					//curCursor.PushIfElseStack(navigator.ContextTypeElseBlock)
					curCursor.SetCursorContext(navigator.ContextTypeElseBlock)
				}
			}
		}
	}
	fmt.Println("Exiting IfElseStmtContextHandler .........................")
	return nil
}
