package navigator

import (
	"errors"
	"fmt"

	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/variables"
)

var (
	cursor *Cursor
)

type OperatorType string

const (
	OperatorTypeAnd         OperatorType = "&&"
	OperatorTypeOR                       = "||"
	OperatorTypeEqual                    = "=="
	OperatorTypeLarger                   = ">"
	OperatorTypeLargerEqual              = ">="
	OperatorTypeLess                     = "<"
	OperatorTypeLessEqual                = "<="
)

// record useful current context information
type Cursor struct {
	curSymTable  *sym_tables.SymTable
	curStatement *Statement
	curIndex     int
	curContext   sym_tables.ContextType
	curJudge     bool
	// curExpr      *Expr
	curExpr         string
	curExprVarNames []string
	curIfElseClause *sym_tables.IfElseClause
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
		//////////////////////////////////

	}

	return newCursor
}

func (this *Cursor) SetIfElseClause(ifElseClause *sym_tables.IfElseClause) {
	this.curIfElseClause = ifElseClause
}

func (this *Cursor) GetIfElseClause() *sym_tables.IfElseClause {
	return this.curIfElseClause
}

func (this *Cursor) InitIfElseClause() {
	this.curIfElseClause = sym_tables.NewIfElseClause()
}

func (this *Cursor) GetExprVarNames() []string {
	return this.curExprVarNames
}

func (this *Cursor) AddExprVarNames(varName string) {
	this.curExprVarNames = append(this.curExprVarNames, varName)
}

func (this *Cursor) InitExprVarNames() {
	this.curExprVarNames = make([]string, 0)
}

func (this *Cursor) GetExpr() string {
	return this.curExpr
}

func (this *Cursor) PushExpr(exprComponent string) {
	this.curExpr += exprComponent
}

func (this *Cursor) ClearExpr() {
	this.curExpr = ""
}

func (this *Cursor) SetCursorContext(contextType sym_tables.ContextType) {
	this.curContext = contextType
}

func (this *Cursor) GetCursorContext() sym_tables.ContextType {
	return this.curContext
}

func (this *Cursor) SetCurJudge(judgeRes bool) {
	this.curJudge = judgeRes
}

func (this *Cursor) GetCurJudge() bool {
	return this.curJudge
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
