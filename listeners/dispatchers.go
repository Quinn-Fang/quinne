package listeners

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/grammars"
	"quinn007.com/listeners/handler"
)

func LexDispatcher(goListener *GoListener, antlrCtx antlr.ParserRuleContext, ruleName string) error {
	fmt.Println("33332222------- ", ruleName)
	//if ruleName == grammars.Statement {
	//	handler.StatementListener(antlrCtx)
	//} else if ruleName == grammars.Block {
	//	handler.
	//}
	//if ruleName == grammars.Block {
	//	fmt.Println("1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	//	if goListener.Test_1 {
	//		return nil
	//	}
	//	goListener.Test_1 = true
	//	handler.BlockContextHandler(antlrCtx, nil)
	//} else if ruleName == grammars.SourceFile {
	//	fmt.Println("2222222222222222222222222222222222222222222222222222222222222222222222222222222222222222")
	//	handler.SourceFileHandler(antlrCtx)
	//}
	if ruleName == grammars.SourceFile {
		fmt.Println("1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
		handler.SourceFileHandler(antlrCtx)
	} else if ruleName == grammars.Block {
		//if goListener.Test_1 {
		//	return nil
		//}
		//goListener.Test_1 = true
		fmt.Println("2222222222222222222222222222222222222222222222222222222222222222222222222222222222222222")
		//handler.BlockContextHandler(antlrCtx, nil)
	}
	return nil
}
