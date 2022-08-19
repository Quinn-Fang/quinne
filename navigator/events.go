package navigator

//
//import (
//	"github.com/Quinn-Fang/Quinne/procedures"
//)
//
//type EventType int
//
//const (
//	EventTypeFunction EventType = 1
//)
//
//type Event struct {
//	eventType EventType
//	function  *procedures.FFunction
//}
//
//func NewEmptyEvent() *Event {
//	newEvent := &Event{}
//	return newEvent
//}
//
//func NewEvent(eType EventType) *Event {
//	newEvent := &Event{
//		eventType: eType,
//	}
//	return newEvent
//}
//
//func (this *Event) SetEventType(eType EventType) {
//	this.eventType = eType
//}
//
//func (this *Event) GetEventType() EventType {
//	return this.eventType
//}
//
//func (this *Event) SetFunction(fFunction *procedures.FFunction) {
//	this.function = fFunction
//}
//
//func (this *Event) GetFunction() *procedures.FFunction {
//	return this.function
//}
//
//type EventQueue struct {
//	queue []*Event
//	// record traversal time index, not a property of the struct itself
//	curIndex int
//}
//
//func NewEventQueue() *EventQueue {
//	newEventQueue := &EventQueue{
//		queue:    make([]*Event, 0),
//		curIndex: 0,
//	}
//	return newEventQueue
//}
//
//func (this *EventQueue) Enqueue(event *Event) {
//	this.queue = append(this.queue, event)
//}
//
//func (this *EventQueue) IncIndex() {
//	this.curIndex++
//}
//
//type Iterator struct {
//	curFuncIndex        int
//	curCodeSegmentIndex int
//}
//
//func NewIterator() *Iterator {
//	newIterator := &Iterator{
//		curFuncIndex: 0,
//	}
//
//	return newIterator
//}
//
//func (this *Iterator) GetFuncIndex() int {
//	return this.curFuncIndex
//}
//
//func (this *Iterator) IncFuncIndex() {
//	this.curFuncIndex++
//}
//
//func (this *Iterator) SetFuncIndex(funcIndex int) {
//	this.curFuncIndex = funcIndex
//}
//
//func (this *Iterator) GetCodeSegmentIndex() int {
//	return this.curCodeSegmentIndex
//}
//
//func (this *Iterator) IncCodeSegmentIndex() {
//	this.curCodeSegmentIndex++
//}
