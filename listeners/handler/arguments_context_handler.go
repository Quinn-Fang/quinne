package handler

import (
	"fmt"

	"quinn007.com/parser"
)

func ArgumentsContextHandler(contextParser *parser.ArgumentsContext) error {
	fmt.Println("789101112...")
	children := contextParser.GetChildren()

	for _, child := range children {
		fmt.Println("                  ")
		fmt.Printf("%T\n", child)
		fmt.Printf("%+v\n", child)
		fmt.Println("                  ")
		//switch parserContext := child.(type) {
		//case *parser.IdentifierListContext:
		//	{
		//	}
		//}
	}

	return nil
}
