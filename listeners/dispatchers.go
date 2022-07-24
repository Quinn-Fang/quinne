package listeners

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/grammars"
)

func LexDispatcher(goListener *GoListener, antlrCtx antlr.ParserRuleContext, ruleName string) error {
	if ruleName == grammars.IdentifierList {
		fmt.Println("33332222-------")
	} else if ruleName == grammars.VarDecl {
		VarDeclListener(goListener, antlrCtx)
	} else if ruleName == grammars.ShortVarDecl {
		ShortVarDeclListener(goListener, antlrCtx)
	}
	return nil
}
