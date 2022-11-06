package tests

import (
	"fmt"
	"testing"

	"github.com/Quinn-Fang/quinne/quinne"
	"github.com/Quinn-Fang/quinne/uspace"
)

func Test_7(t *testing.T) {
	eventHandler := quinne.NewEventHandler("data/sample_lambda_2.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		fmt.Printf("%+v\n", event)

		if event.GetEventType() == uspace.EventTypeFunction {

			fmt.Println()
			retVar, err := event.GetVarByName("ret")
			if err != nil {
				panic(err)
			}
			retValue, _ := retVar.GetVariableValue().(bool)
			expected := false
			if retValue != expected {
				errMsg := fmt.Sprintf(funcReturnNotEqual, retValue, expected)
				t.Error(errMsg)
			}
		}
		event, err = eventHandler.GetNextEvent()
	}
}
