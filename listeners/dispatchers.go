package listeners

import (
	"github.com/Quinn-Fang/quinne/grammars"
	"github.com/Quinn-Fang/quinne/listeners/handler"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func LexDispatcher(goListener *GoListener, antlrCtx antlr.ParserRuleContext, ruleName string) error {
	if ruleName == grammars.SourceFile {
		handler.SourceFileHandler(antlrCtx)
	} else if ruleName == grammars.Block {
	}
	return nil
}
