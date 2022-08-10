package navigator

import (
	"errors"
	"fmt"
)

func (this *Navigator) GetNextEvent() (*Event, error) {
	//fmt.Println("33333")
	//fmt.Printf("%+v \n", this.eventQueue.queue)

	fmt.Println("---------------------- ", this.iterator.GetFuncIndex(), this.iterator.GetCodeSegmentIndex())
	if this.eventQueue.curIndex == len(this.eventQueue.queue) {
		return NewEmptyEvent(), errors.New("Reached end of queue")
	}

	event := this.eventQueue.queue[this.eventQueue.curIndex]
	if event.GetEventType() == EventTypeFunction {
		curNavigator := GetCurNavigator()
		curCodeSegment := curNavigator.GetCodeSegment()
		codeSegmentQueue := curCodeSegment.GetQueue()
		for _, v := range codeSegmentQueue {
			fmt.Printf("%+v\n", v)
		}
		funcIndex := this.iterator.GetFuncIndex()
		curCursor := codeSegmentQueue[this.iterator.GetCodeSegmentIndex()]

		if funcIndex > curCursor.GetFuncEndIndex() {
			this.iterator.IncCodeSegmentIndex()
			if this.iterator.GetCodeSegmentIndex() >= len(curCodeSegment.GetQueue()) {
				return NewEmptyEvent(), errors.New("No More Element")
			}
			curCursor = codeSegmentQueue[this.iterator.GetCodeSegmentIndex()]
			this.iterator.SetFuncIndex(curCursor.GetFuncStartIndex())
		}

		curSymTable := curCursor.GetSymTable()
		curFunction := curSymTable.GetFunctions()[this.iterator.GetFuncIndex()]
		event.SetFunction(curFunction)

		this.iterator.IncFuncIndex()
		this.eventQueue.curIndex++
		fmt.Println("---------------------- ", this.iterator.GetFuncIndex(), this.iterator.GetCodeSegmentIndex())

		return event, nil

	}

	return NewEmptyEvent(), errors.New("Unknown error")
}
