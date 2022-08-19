package listeners

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/grammars"
	"quinn007.com/listeners/handler"
)

func LexDispatcher(goListener *GoListener, antlrCtx antlr.ParserRuleContext, ruleName string) error {
	if ruleName == grammars.SourceFile {
		handler.SourceFileHandler(antlrCtx)
	} else if ruleName == grammars.Block {
	}
	return nil
}
