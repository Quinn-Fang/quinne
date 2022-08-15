package handler

import "quinn007.com/parser"

func ForStatementContextHandler(contextParser *parser.ForStmtContext) error {
	children := contextParser.GetChildren()
	for _, v := range children {
		switch parserContext := v.(type) {
		case *parser.BlockContext:
			{
				BlockContextHandler(parserContext)
			}
		}
	}
	return nil
}
