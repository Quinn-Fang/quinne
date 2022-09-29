package listeners

import (
	"github.com/Quinn-Fang/quinne/grammars"
	"github.com/Quinn-Fang/quinne/listeners/handlerV2"
	"github.com/Quinn-Fang/quinne/scanner"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func LexDispatcher(goListener *GoListener, antlrCtx antlr.ParserRuleContext, ruleName string, scanner *scanner.Scanner) error {
	if ruleName == grammars.SourceFile {
		// handler.SourceFileHandler(antlrCtx)
		handlerV2.SourceFileHandler(antlrCtx, scanner)
	} else if ruleName == grammars.Block {
	}
	return nil
}
