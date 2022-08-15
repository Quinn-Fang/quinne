package utils

import (
	"container/list"
	"errors"
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

func (this *Queue) PushBack(item interface{}) {
	this.queue.PushBack(item)
}

func (this *Queue) PushFront(item interface{}) {
	this.queue.PushFront(item)
}

func (this *Queue) GetFront() interface{} {
	return this.queue.Front()
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

func (this *Queue) Clear() {
	this.queue.Init()
}
