package uspace

import (
	"quinn007.com/procedures"
	"quinn007.com/sym_tables"
	"quinn007.com/utils"
)

type EventType int

const (
	EventTypeIfElseExpr EventType = 1
	EventTypeFunction             = 2
)

type EventQueue struct {
	queue utils.Queue
}

type Event struct {
	eventType    EventType
	eventPointer interface{}
}

func NewEventQueue() *EventQueue {
	NewEventQueue := &EventQueue{
		queue: *utils.NewQueue(),
	}
	return NewEventQueue
}

func NewEvent(eventType EventType) *Event {
	newEvent := &Event{
		eventType: eventType,
	}
	return newEvent
}

func (this *Event) SetEvent(eventPointer interface{}) {
	this.eventPointer = eventPointer
}

func (this *Event) GetEventType() EventType {
	return this.eventType
}

func (this *Event) GetEventContext() interface{} {
	return this.eventPointer
}

func (this *Event) GetExpr(eventContext interface{}) (string, []string) {
	if ifElseBranch, ok := eventContext.(*sym_tables.IfElseBranch); ok {
		return ifElseBranch.GetExpr(), ifElseBranch.GetExprVarNames()
	} else {
		panic("Getting ifelse expr error ")
	}
}

func (this *Event) GetFunction(eventContext interface{}) *procedures.FFunction {
	if fFunction, ok := eventContext.(*procedures.FFunction); ok {
		return fFunction
	} else {
		panic("Getting function from event error")
	}
}

func (this *EventQueue) AddEvent(event *Event) {
	this.PushBack(event)
}

func (this *EventQueue) PushBack(event *Event) {
	this.queue.PushBack(event)
}

func (this *EventQueue) PopFront() (*Event, error) {
	ret, err := this.queue.PopFront()
	if err != nil {
		return nil, err
	}
	return ret.(*Event), err
}
