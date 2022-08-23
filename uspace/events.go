package uspace

import (
	"fmt"

	"github.com/Quinn-Fang/quinne/procedures"
	"github.com/Quinn-Fang/quinne/procedures/buildin"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/utils"
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

func (this *Event) FillExpr() {
	if ifElseBranch, ok := this.eventPointer.(*sym_tables.IfElseBranch); ok {
		varMap := make(map[string]interface{})
		for _, varName := range ifElseBranch.GetExprVarNames() {
			if variable, err := this.GetSymTable().GetVariableByName(varName); err != nil {
				errInfo := fmt.Sprintf("Variable: '%s' does not exist.", varName)
				panic(errInfo)
			} else {
				varMap[varName] = variable.GetVariableValue()
			}
		}
		res := utils.ParseExpr(ifElseBranch.GetExpr(), varMap)
		ifElseBranch.SetJudgeRes(res)
	} else {
		panic("Not ifelse expr error ")
	}
}

func (this *Event) FillExprV2(varNames []string) map[string]interface{} {
	// fill in variables that have not been filled by user
	ret := make(map[string]interface{})
	symTable := this.GetSymTable()
	funcTable := buildin.GetSystemFuncTable()

	for _, varName := range varNames {
		found := false
		// first look up from system build in functions
		if function, err := funcTable.GetFunctionByName(varName); err == nil {
			found = true
			ret[varName] = function
		} else {
			// then look up variable from symbol table
			if variable, err := symTable.GetVariableByName(varName); err == nil {
				found = true
				ret[varName] = variable.GetVariableValue()
			}

		}

		if !found {
			errInfo := fmt.Sprintf("Variable: '%s' does not exist.", varName)
			panic(errInfo)
		}

	}
	return ret
}

func (this *Event) SetExpr(varMap map[string]interface{}) {
	// varMap includes user defined functions and variables only
	if ifElseBranch, ok := this.eventPointer.(*sym_tables.IfElseBranch); ok {
		userVarMap := make(map[string]interface{})
		systemVarNames := make([]string, 0)
		for _, varName := range ifElseBranch.GetExprVarNames() {
			if varValue, in := varMap[varName]; in {
				userVarMap[varName] = varValue
			} else {
				systemVarNames = append(systemVarNames, varName)
			}
		}

		if len(userVarMap) != len(varMap) {
			panic("Not all provided variable are in expression")
		}
		systemVarMap := this.FillExprV2(systemVarNames)
		allVarMap := utils.MergeMaps(userVarMap, systemVarMap)

		// res := utils.ParseExpr(ifElseBranch.GetExpr(), varMap)
		res := utils.ParseExprV2(ifElseBranch.GetExpr(), allVarMap)
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
