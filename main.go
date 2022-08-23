package main

import (
	"fmt"

	"github.com/Quinn-Fang/quinne/quinne"
)

func main() {
	// quinne.TNewListener("tests/data/sample_if_else.go")
	eventHandler := quinne.NewEventHandler("tests/data/sample_if_else.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		fmt.Printf("%+v\n", event)
		event, err = eventHandler.GetNextEvent()
	}
}
