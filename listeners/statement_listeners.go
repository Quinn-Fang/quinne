package listeners

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/listeners/handler"
	"quinn007.com/parser"
)

func StatementListener(goListener *GoListener, antlrCtx antlr.ParserRuleContext) error {
	fmt.Println("inside StatementListener ... .......")
	children := antlrCtx.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.SimpleStmtContext:
			{
				handler.SimpleStmtContextHandler(parserContext)
			}
		case *parser.IfStmtContext:
			{
				handler.IfElseStmtContextHandler(parserContext)
			}
		}
	}
	return nil
}
