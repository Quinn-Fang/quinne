package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Quinn-Fang/quinne/quinne"
	"github.com/Quinn-Fang/quinne/uspace"
)

func Test_1(t *testing.T) {
	eventHandler := quinne.NewEventHandler("data/sample_001.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		if event.GetEventType() == uspace.EventTypeFunction {
			// deal with how function show be executed here
			// and provide the return value
			fFunction := event.GetFunction(event.GetEventContext())
			if fFunction.GetFunctionName() == "create_battery" {

				fFunction.SetReturnValue(true)
				fReturn := fFunction.GetReturnValue().(bool)
				expected := true
				if fReturn != expected {
					errMsg := fmt.Sprintf(funcReturnNotEqual, fReturn, expected)
					t.Error(errMsg)
				}
				//		fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
			} else if fFunction.GetFunctionName() == "create_switch" {
				fFunction.SetReturnValue("success")
				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
			} else if fFunction.GetFunctionName() == "create_bulb" {
				// Get Function arguments, maybe put some check on them ...
				params := fFunction.GetParams()
				for _, v := range params {
					value := v.GetVariableValue().(string)
					expected := "bulb-2"
					if value != expected {
						errMsg := fmt.Sprintf(paramNotEqual, value, expected)
						t.Error(errMsg)
					}
					// fmt.Printf("%+v ", v.GetVariableValue())
				}

				// Set the return value for this particular function
				fFunction.SetReturnValue("success")
				fReturn := fFunction.GetReturnValue().(string)
				expected := "success"
				if fReturn != expected {
					errMsg := fmt.Sprintf(funcReturnNotEqual, fReturn, expected)
					t.Error(errMsg)
				}
				// fmt.Printf("%+v %+v\n\n", fFunction, fFunction.GetReturnValue())
			}
		} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
			// Get the if expression or else-if expression and variables within that you
			// should provide value or assigned automatically if has defined previously

			ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
			if strings.Contains(ifElseExpr, "SWITCH_ON") {
				fmt.Printf("%+v %+v\n", ifElseExpr, ifElseExprVarNames)
				varMap := make(map[string]interface{}, 8)
				varMap["SWITCH_ON"] = false
				event.SetExpr(varMap)

				// filled variables automatically if already defined and can be accessed
				// by scope rules

				//event.FillExpr()
			}

		}

		event, err = eventHandler.GetNextEvent()

		fmt.Println()
	}
}
