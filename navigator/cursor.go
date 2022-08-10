package navigator

import (
	"errors"
	"fmt"

	"quinn007.com/sym_tables"
	"quinn007.com/variables"
)

var (
	cursor *Cursor
)

type ContextType int

const (
	ContextTypeDefault      ContextType = 1
	ContextTypeFunctionName             = 2
	ContextTypeFunctionArgs             = 3
)

type Cursor struct {
	curSymTable  *sym_tables.SymTable
	curStatement *Statement
	curIndex     int
	curContext   ContextType
}

type Statement struct {
	leftValues  []string
	rightValues []*variables.Variable
}

func InitCursor() {
	newCursor := NewCursor()
	cursor = newCursor
}

func GetCursor() (*Cursor, error) {
	if cursor != nil {
		return cursor, nil
	} else {
		return nil, errors.New("Cursor not initialized ... ")
	}
}

func NewCursor() *Cursor {
	newCursor := &Cursor{
		curIndex: 1,
	}

	return newCursor
}

func (this *Cursor) SetCursorContext(contextType ContextType) {
	this.curContext = contextType
}

func (this *Cursor) GetCursorContext() ContextType {
	return this.curContext
}

func (this *Cursor) PrintStatement() {
	fmt.Println()
	fmt.Println("---------- Value  ----------")
	for i, v := range this.curStatement.leftValues {
		fmt.Printf("variable name: %s ", v)
		rightValue := this.curStatement.rightValues[i]
		fmt.Printf("variable type: %d ", rightValue.GetVariableType())
		fmt.Printf("variable value: %s ", rightValue.GetVariableValue())
		fmt.Println()
	}
	fmt.Println()
}

func (this *Cursor) GetIndex() int {
	return this.curIndex
}

func (this *Cursor) IncreaseIndex() {
	this.curIndex++
}

func (this *Cursor) NewStatement() {
	this.curStatement = NewStatement()
}

func (this *Cursor) GetStatement() *Statement {
	return this.curStatement
}

func NewStatement() *Statement {
	newStatement := &Statement{
		leftValues:  make([]string, 0),
		rightValues: make([]*variables.Variable, 0),
	}
	return newStatement
}

func (this *Statement) AddLeftValue(leftValue string) {
	this.leftValues = append(this.leftValues, leftValue)
}

func (this *Statement) GetLeftValues() []string {
	return this.leftValues
}

func (this *Statement) AddRightValue(rightValue *variables.Variable) {
	this.rightValues = append(this.rightValues, rightValue)
}

func (this *Statement) GetRightValues() []*variables.Variable {
	return this.rightValues
}

func (this *Statement) Assign() error {
	if len(this.leftValues) != len(this.rightValues) {
		return errors.New("number of left values and right values does not Match")
	}

	for i := 0; i < len(this.leftValues); i++ {
		leftValue := this.leftValues[i]
		rightValue := this.rightValues[i]
		rightValue.SetVariableName(leftValue)
	}
	return nil

}
