package main

import (
	"fmt"
	"strings"

	"github.com/Quinn-Fang/quinne/quinne"
	"github.com/Quinn-Fang/quinne/uspace"
)

//func runListener() {
//	input, _ := antlr.NewFileStream("samples/sample_5.go")
//	// Create First SymTable
//	sym_tables.NewEntryTable()
//	// Create Cursor
//	navigator.InitCursor()
//	curNavigator := navigator.NewNavigator()
//	navigator.SetCurNavigator(curNavigator)
//
//	// Create the Lexer
//	lexer := parser.NewGoLexer(input)
//	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
//
//	// Create the Parser
//	p := parser.NewGoParser(stream)
//
//	//listen := GoListener{}
//	tree := p.SourceFile()
//	antlr.ParseTreeWalkerDefault.Walk(listeners.NewGoListener(p, tree), tree)
//	// utils.PrintAllSymTale()
//	//curNavigator.PrintStack()
//
//	//curNavigator.PrintCodeSegments()
//}

//func main_test_1() {
//	quinne.RunListener()
//	fmt.Println("................... Start testing .......................")
//	newNavigator := navigator.GetNavigator()
//	event, err := newNavigator.GetNextEvent()
//	for err == nil {
//		fmt.Println("------------------  ***  --------------------")
//		fmt.Printf("%+v \n", event)
//		fmt.Printf("%+v \n", event.GetEventContext())
//		if event.GetEventType() == uspace.EventTypeFunction {
//			fFunction := event.GetFunction(event.GetEventContext())
//			// fmt.Printf("| %+v %+v Executable: %+v\n", fFunction, fFunction.GetReturnValue(), event.GetSymTable().IsExecutable())
//			fmt.Printf("| %+v %+v is executable ? : %+v \n", fFunction, fFunction.GetReturnValue(), event.GetSymTable().IsExecutable())
//			if fFunction.GetFunctionName() == "BodylessFunction_3" {
//				st := event.GetSymTable()
//				x, err := st.GetVariableByName("var_2")
//				if err != nil {
//					panic(err)
//				}
//				fmt.Println(x)
//				fFunction.SetReturnValue("returnValue 111")
//				fmt.Println("9999999999")
//				fmt.Println(x)
//			}
//
//			if fFunction.GetFunctionName() == "secondLastFunction" {
//				event.GetSymTable().IsExecutable()
//			}
//		} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
//			ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
//			fmt.Printf("| %+v %+v \n", ifElseExpr, ifElseExprVarNames)
//			if strings.Contains(ifElseExpr, "age>6") {
//				//varMap := make(map[string]interface{}, 8)
//				//varMap["age"] = 29
//				//event.SetExpr(varMap)
//				event.FillExpr()
//			}
//		} else if event.GetEventType() == uspace.EventTypeForLoop {
//			event.GetSymTable().IsExecutable()
//			fmt.Println()
//		} else if event.GetEventType() == uspace.EventTypeFunctionDecl {
//			event.GetSymTable().IsExecutable()
//			fmt.Println()
//		}
//
//		//fmt.Printf("%+v ", event.GetEventType())
//		//if event.GetEventType() == uspace.EventTypeFunction {
//		//	fmt.Printf("%+v \n", event.GetFunction(event.GetEventContext()))
//		//} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
//		//	v1, v2 := event.GetExpr(event.GetEventContext())
//		//	fmt.Printf("%+v %+v \n", v1, v2)
//		//}
//		event, err = newNavigator.GetNextEvent()
//	}
//
//}

func main() {
	eventHandler := quinne.NewEventHandler("samples/sample_5.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		fmt.Printf("%+v\n", event)
		if event.GetEventType() == uspace.EventTypeFunction {
			fFunction := event.GetFunction(event.GetEventContext())
			if fFunction.GetFunctionName() == "BodylessFunction_3" {
				// deal with how function show be executed here
				// and provide the return value

				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
				fFunction.SetReturnValue("function 1's return value'")
				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
			} else if fFunction.GetFunctionName() == "secondLastFunction" {
				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
				fFunction.SetReturnValue(2)
				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
			}
		} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
			ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
			if strings.Contains(ifElseExpr, "age>6") {
				fmt.Printf("%+v %+v\n", ifElseExpr, ifElseExprVarNames)
				varMap := make(map[string]interface{}, 8)
				varMap["age"] = 5
				event.SetExpr(varMap)
				//event.FillExpr()
			}

		}

		event, err = eventHandler.GetNextEvent()

	}
}
