package uspace

import (
	"fmt"

	"quinn007.com/procedures"
	"quinn007.com/sym_tables"
	"quinn007.com/utils"
)

type EventType int

const (
	EventTypeIfElseExpr   EventType = 1
	EventTypeFunction               = 2
	EventTypeFunctionDecl           = 3
	EventTypeForLoop                = 4
)

type EventQueue struct {
	queue utils.Queue
}

type Event struct {
	eventType    EventType
	eventPointer interface{}
	symTable     *sym_tables.SymTable
}

func NewEventQueue() *EventQueue {
	NewEventQueue := &EventQueue{
		queue: *utils.NewQueue(),
	}
	return NewEventQueue
}

func NewEvent(eventType EventType, curSymTable *sym_tables.SymTable) *Event {
	newEvent := &Event{
		eventType: eventType,
		symTable:  curSymTable,
	}
	return newEvent
}

func (this *Event) GetSymTable() *sym_tables.SymTable {
	return this.symTable
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

func (this *Event) SetExpr(varMap map[string]interface{}) {
	if ifElseBranch, ok := this.eventPointer.(*sym_tables.IfElseBranch); ok {
		for _, varName := range ifElseBranch.GetExprVarNames() {
			if _, in := varMap[varName]; !in {
				err := fmt.Sprintf("Variable: %+v Unset", varName)
				panic(err)
			}
		}
		res := utils.ParseExpr(ifElseBranch.GetExpr(), varMap)
		ifElseBranch.SetJudgeRes(res)
	} else {
		panic("Not ifelse expr error ")
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
