package main

import (
	"github.com/Quinn-Fang/quinne/tests/examples"
)

func main() {
	// quinne.TNewListener("tests/data/sample_if_else.go")
	//eventHandler := quinne.NewEventHandler("tests/data/sample_if_else.go")
	//event, err := eventHandler.GetNextEvent()
	//for err == nil {
	//	fmt.Printf("%+v\n", event)
	//	event, err = eventHandler.GetNextEvent()
	//}
	examples.CreateBulbDiagram()
}
