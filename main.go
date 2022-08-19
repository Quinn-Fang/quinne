package main

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/listeners"
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
	"quinn007.com/sym_tables/utils"
	"quinn007.com/uspace"
)

func runListener() {
	input, _ := antlr.NewFileStream("samples/sample_5.go")
	// Create First SymTable
	sym_tables.NewEntryTable()
	// Create Cursor
	navigator.InitCursor()
	curNavigator := navigator.NewNavigator()
	navigator.SetCurNavigator(curNavigator)

	// Create the Lexer
	lexer := parser.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGoParser(stream)

	//listen := GoListener{}
	tree := p.SourceFile()
	antlr.ParseTreeWalkerDefault.Walk(listeners.NewGoListener(p, tree), tree)
	utils.PrintAllSymTale()
	//curNavigator.PrintStack()

	//curNavigator.PrintCodeSegments()
}

func main() {
	runListener()
	fmt.Println("................... Start testing .......................")
	newNavigator := navigator.GetNavigator()
	event, err := newNavigator.GetNextEvent()
	for err == nil {
		fmt.Println("------------------  ***  --------------------")
		fmt.Printf("%+v \n", event)
		fmt.Printf("%+v \n", event.GetEventContext())
		if event.GetEventType() == uspace.EventTypeFunction {
			fFunction := event.GetFunction(event.GetEventContext())
			// fmt.Printf("| %+v %+v Executable: %+v\n", fFunction, fFunction.GetReturnValue(), event.GetSymTable().IsExecutable())
			fmt.Printf("| %+v %+v is executable ? : %+v \n", fFunction, fFunction.GetReturnValue(), event.GetSymTable().IsExecutable())
			if fFunction.GetFunctionName() == "BodylessFunction_3" {
				st := event.GetSymTable()
				x, err := st.GetVariableByName("var_2")
				if err != nil {
					panic(err)
				}
				fmt.Println(x)
				fFunction.SetReturnValue("returnValue 111")
				fmt.Println("9999999999")
				fmt.Println(x)
			}

		} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
			ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
			fmt.Printf("| %+v %+v \n", ifElseExpr, ifElseExprVarNames)
			if strings.Contains(ifElseExpr, "age>6") {
				//varMap := make(map[string]interface{}, 8)
				//varMap["age"] = 29
				//event.SetExpr(varMap)
				event.FillExpr()
			}
		} else if event.GetEventType() == uspace.EventTypeForLoop {
			event.GetSymTable().IsExecutable()
			fmt.Println()
		} else if event.GetEventType() == uspace.EventTypeFunctionDecl {
			event.GetSymTable().IsExecutable()
			fmt.Println()
		}

		//fmt.Printf("%+v ", event.GetEventType())
		//if event.GetEventType() == uspace.EventTypeFunction {
		//	fmt.Printf("%+v \n", event.GetFunction(event.GetEventContext()))
		//} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
		//	v1, v2 := event.GetExpr(event.GetEventContext())
		//	fmt.Printf("%+v %+v \n", v1, v2)
		//}
		event, err = newNavigator.GetNextEvent()
	}
}
