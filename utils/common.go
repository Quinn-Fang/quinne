package utils

import (
	"container/list"
	"errors"

	"github.com/Knetic/govaluate"
	exprV2 "github.com/antonmedv/expr"
)

type Queue struct {
	queue *list.List
}

func NewQueue() *Queue {
	newQueue := &Queue{
		queue: list.New(),
	}
	return newQueue
}

func (this *Queue) IsEmpty() bool {
	return this.queue.Len() == 0
}

func (this *Queue) PushBack(item interface{}) {
	this.queue.PushBack(item)
}

func (this *Queue) PushFront(item interface{}) {
	this.queue.PushFront(item)
}

func (this *Queue) GetFront() *list.Element {
	return this.queue.Front()
}

func (this *Queue) GetBack() *list.Element {
	return this.queue.Back()
}

func (this *Queue) PopFront() (interface{}, error) {
	item := this.queue.Front()
	var ret interface{}
	if item != nil {
		ret = item.Value
		this.queue.Remove(item)
	} else {
		return nil, errors.New("queue empty")
	}
	return ret, nil
}

func (this *Queue) PopBack() (interface{}, error) {
	item := this.queue.Back()
	var ret interface{}
	if item != nil {
		ret = item.Value
		this.queue.Remove(item)
	} else {
		return nil, errors.New("queue empty")
	}
	return ret, nil
}

func (this *Queue) Clear() {
	this.queue.Init()
}

func ParseExpr(expr string, parameters map[string]interface{}) bool {
	expression, err := govaluate.NewEvaluableExpression(expr)

	// parameters := make(map[string]interface{}, 8)

	result, err := expression.Evaluate(parameters)
	if err != nil {
		panic(err)
	}

	if ret, ok := result.(bool); ok {
		return ret
	} else {
		panic("result is not bool")
	}
}

func ParseExprV2(expr string, parameters map[string]interface{}) bool {
	out, err := exprV2.Eval(expr, parameters)

	if err != nil {
		panic(err)
	}

	if ret, ok := out.(bool); ok {
		return ret
	} else {
		panic("result is not bool")
	}
}

func MergeMaps(mapA map[string]interface{}, mapB map[string]interface{}) map[string]interface{} {
	for k, v := range mapB {
		mapA[k] = v
	}
	return mapA
}
