package utils

import (
	"errors"

	"github.com/Quinn-Fang/quinne/sym_tables"
)

type SymTableCursor struct {
	symTable       *sym_tables.SymTable
	funcStartIndex int
	funcEndIndex   int
}

func NewSymTableCursor() *SymTableCursor {
	newSymTableCursor := &SymTableCursor{
		funcStartIndex: 0,
		funcEndIndex:   0,
	}
	return newSymTableCursor
}

func (this *SymTableCursor) SetSymTable(symTable *sym_tables.SymTable) {
	this.symTable = symTable
}

func (this *SymTableCursor) GetSymTable() *sym_tables.SymTable {
	return this.symTable
}

func (this *SymTableCursor) SetFuncStartIndex(funcStartIndex int) {
	this.funcStartIndex = funcStartIndex
}

func (this *SymTableCursor) GetFuncStartIndex() int {
	return this.funcStartIndex
}

func (this *SymTableCursor) SetFuncEndIndex(funcEndIndex int) {
	this.funcEndIndex = funcEndIndex
}

func (this *SymTableCursor) GetFuncEndIndex() int {
	return this.funcEndIndex
}

type SymTableCursorStack struct {
	stack []*SymTableCursor
}

func NewStack() *SymTableCursorStack {
	newStack := &SymTableCursorStack{
		stack: make([]*SymTableCursor, 0),
	}

	return newStack
}

func (this *SymTableCursorStack) GetStack() []*SymTableCursor {
	return this.stack
}

func (this *SymTableCursorStack) IsEmpty() bool {
	return len(this.stack) == 0
}

func (this *SymTableCursorStack) Push(symTable *SymTableCursor) {
	this.stack = append(this.stack, symTable)
}

func (this *SymTableCursorStack) Pop() (*SymTableCursor, error) {
	if this.IsEmpty() {
		return &SymTableCursor{}, errors.New("No Symbol Table on stack")
	}

	lastSymTable := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	return lastSymTable, nil
}

func (this *SymTableCursorStack) Peek() (*SymTableCursor, error) {
	if this.IsEmpty() {
		return NewSymTableCursor(), errors.New("No Symbol Table on stack")
	}
	return this.stack[len(this.stack)-1], nil
}

type CodeSegment struct {
	queue []*SymTableCursor
}

func NewCodeSegment() *CodeSegment {
	newCodeSegment := &CodeSegment{
		queue: make([]*SymTableCursor, 0),
	}
	return newCodeSegment
}

func (this *CodeSegment) InsertFront(symTableCursor *SymTableCursor) {
	this.queue = append([]*SymTableCursor{symTableCursor}, this.queue...)
}

func (this *CodeSegment) InsertBack(symTableCursor *SymTableCursor) {
	this.queue = append(this.queue, symTableCursor)
}

func (this *CodeSegment) GetQueue() []*SymTableCursor {
	return this.queue
}
