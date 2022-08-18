package navigator

import (
	"fmt"

	"quinn007.com/navigator/utils"
	"quinn007.com/sym_tables"
	"quinn007.com/uspace"
)

var (
	curNavigator *Navigator
)

type Navigator struct {
	symTableCursorStack *utils.SymTableCursorStack
	codeSegment         *utils.CodeSegment
	eventQueue          *uspace.EventQueue
	//eventQueue          *EventQueue
	//iterator            *Iterator
}

func NewNavigator() *Navigator {
	newNavigator := &Navigator{
		symTableCursorStack: utils.NewStack(),
		codeSegment:         utils.NewCodeSegment(),
		//eventQueue:          NewEventQueue(),
		//iterator:            NewIterator(),
		eventQueue: uspace.NewEventQueue(),
	}
	return newNavigator
}

func (this *Navigator) AddEvent(eventType uspace.EventType, event interface{}, symTable *sym_tables.SymTable) {
	newEvent := uspace.NewEvent(eventType, symTable)
	newEvent.SetEvent(event)
	this.eventQueue.AddEvent(newEvent)
}

func (this *Navigator) GetNextEvent() (*uspace.Event, error) {
	event, err := this.eventQueue.PopFront()
	//if event.GetEventType() == sym_tables.ContextTypeFuncDecl {
	//	return this.GetNextEvent()
	//} else if event.GetEventType() == sym_tables.ContextTypeForLoop {
	//	return this.GetNextEvent()
	//}
	return event, err
}

func SetCurNavigator(navigator *Navigator) {
	curNavigator = navigator
}

func GetNavigator() *Navigator {
	return curNavigator
}

func GetCurNavigator() *Navigator {
	return curNavigator
}

func (this *Navigator) GetSymTableCursorStack() *utils.SymTableCursorStack {
	return this.symTableCursorStack
}

func (this *Navigator) GetCodeSegment() *utils.CodeSegment {
	return this.codeSegment
}

//func (this *Navigator) AddNewEvent(event *Event) {
//	this.eventQueue.Enqueue(event)
//}

func (this *Navigator) PrintStack() {
	fmt.Println("^^^^^^^^^^^^^^^^^ Symbol Table Stack : ^^^^^^^^^^^^^^^^^^")
	for _, v := range this.GetSymTableCursorStack().GetStack() {
		fmt.Println()
		fmt.Println(v)
		v.GetSymTable().PrintFunctions()
	}
}

func (this *Navigator) PrintCodeSegments() {
	fmt.Println("^^^^^^^^^^^^^^^^^ Code Segment : ^^^^^^^^^^^^^^^^^^")
	for _, v := range this.GetCodeSegment().GetQueue() {
		fmt.Println()
		fmt.Println(v)
		v.GetSymTable().PrintFunctions()
	}
}
