package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Quinn-Fang/quinne/quinne"
	"github.com/Quinn-Fang/quinne/uspace"
)

func Test_4(t *testing.T) {
	eventHandler := quinne.NewEventHandler("data/sample_if_else.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		fmt.Printf("%+v\n", event)
		if event.GetEventType() == uspace.EventTypeFunction {
			// deal with how function show be executed here
			// and provide the return value
			fFunction := event.GetFunction(event.GetEventContext())
			if fFunction.GetFunctionName() == "bodyLessFunction_1" {
				t.Log("Executed bodyLessFunction_1")
			}
			if fFunction.GetFunctionName() == "bodyLessFunction_2" {
				t.Log("Executed bodyLessFunction_2")
			}
		} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
			ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
			if strings.Contains(ifElseExpr, "len") {
				fmt.Printf("%+v %+v\n", ifElseExpr, ifElseExprVarNames)
				varMap := make(map[string]interface{}, 8)
				//varMap["age"] = 7
				event.SetExpr(varMap)

				// event.FillExpr()
			}
			//} else if strings.Contains(ifElseExpr, "566") {
			//	expected := "566>test1"
			//	if ifElseExpr != expected {
			//		t.Error(formatError(ifElseExprDoesNotMatch, ifElseExpr, expected))
			//	}
			//} else if strings.Contains(ifElseExpr, "naveen") {
			//	expected := "name==naveen"
			//	if ifElseExpr != expected {
			//		t.Error(formatError(ifElseExprDoesNotMatch, ifElseExpr, expected))
			//	}

			//}

		}

		event, err = eventHandler.GetNextEvent()

		fmt.Println()
	}
}
