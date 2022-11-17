package tests

import (
	"fmt"
	"testing"

	"github.com/Quinn-Fang/quinne/quinne"
	"github.com/Quinn-Fang/quinne/uspace"
)

func Test_Infra_1(t *testing.T) {
	eventHandler := quinne.NewEventHandler("data/infra/infra_sample_1.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		fmt.Printf("%+v\n", event)
		if event.GetEventType() == uspace.EventTypeFunction {
			// deal with how function show be executed here
			// and provide the return value
			fFunction := event.GetFunction(event.GetEventContext())
			switch fFunction.GetFunctionName() {
			case "getCurrentNodeCount":
				{
					fFunction.SetReturnValue(2)
				}
			case "CreateEKSCluster":
				{
					params := fFunction.GetParams()
					for _, v := range params {
						value := v.GetVariableValue()
						fmt.Println("The Value is : ", value)
						//expected := "bulb-2"
						//if value != expected {
						//	errMsg := fmt.Sprintf(paramNotEqual, value, expected)
						//	t.Error(errMsg)
						//}
						// fmt.Printf("%+v ", v.GetVariableValue())
					}

				}
			case "expandEKSCluster":
				{
					params := fFunction.GetParams()
					for _, v := range params {
						value := v.GetVariableValue()
						fmt.Println("The Value is : ", value)
						//expected := "bulb-2"
						//if value != expected {
						//	errMsg := fmt.Sprintf(paramNotEqual, value, expected)
						//	t.Error(errMsg)
						//}
						// fmt.Printf("%+v ", v.GetVariableValue())
					}

				}
			}
		}

		//if event.GetEventType() == uspace.EventTypeFunction {

		//	fmt.Println()
		//	retVar, err := event.GetVarByName("ret")
		//	if err != nil {
		//		panic(err)
		//	}
		//	retValue, _ := retVar.GetVariableValue().(bool)
		//	expected := false
		//	if retValue != expected {
		//		errMsg := fmt.Sprintf(funcReturnNotEqual, retValue, expected)
		//		t.Error(errMsg)
		//	}
		//}
		event, err = eventHandler.GetNextEvent()
	}
}
