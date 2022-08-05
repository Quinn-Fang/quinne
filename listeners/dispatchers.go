package listeners

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/grammars"
)

func LexDispatcher(goListener *GoListener, antlrCtx antlr.ParserRuleContext, ruleName string) error {
	fmt.Println("33332222------- ", ruleName)
	if ruleName == grammars.Statement {
		StatementListener(goListener, antlrCtx)
	}
	return nil
}
