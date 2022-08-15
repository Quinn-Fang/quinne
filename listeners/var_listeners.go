package listeners

//func VarDeclListener(goListener *GoListener, antlrCtx antlr.ParserRuleContext) error {
//	fmt.Println("inside VarDeclListener ... .......")
//	children := antlrCtx.GetChildren()
//	for _, v := range children {
//		switch parserContext := v.(type) {
//		case *parser.VarSpecContext:
//			{
//				handler.VarSpecContextHandler(parserContext)
//			}
//		}
//	}
//	return nil
//}
//
//func IdentifierListListener(goListener *GoListener, antlrCtx antlr.ParserRuleContext) error {
//	return nil
//}
//
//func VarSpecListener(goListener *GoListener, antlrCtx antlr.ParserRuleContext) error {
//	return nil
//}
//
//func ShortVarDeclListener(goListener *GoListener, antlrCtx antlr.ParserRuleContext) error {
//	children := antlrCtx.GetChildren()
//	for _, v := range children {
//		switch parserContext := v.(type) {
//		case *parser.IdentifierListContext:
//			{
//				handler.IdentifierListContextHandler(parserContext)
//			}
//		case *parser.ExpressionListContext:
//			{
//				handler.ExpressionListContextHandler(parserContext)
//			}
//
//		}
//	}
//	return nil
//}
