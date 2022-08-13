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
	ContextTypeIfExpr                   = 4
	ContextTypeIfBlock                  = 5
	ContextTypeElseIfBlock              = 6
	ContextTypeElseBlock                = 7
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

type Cursor struct {
	curSymTable  *sym_tables.SymTable
	curStatement *Statement
	curIndex     int
	curContext   ContextType
	curJudge     bool
	// curExpr      *Expr
	curExpr        string
	curIfElseStack []ContextType
}

type Statement struct {
	leftValues  []string
	rightValues []*variables.Variable
}

//type ExprState int
//
//const (
//	ExprStateEmpty    ExprState = 1
//	ExprStateLeftSet            = 2
//	ExprStateOpSet              = 3
//	ExprStateRightSet           = 4
//)
//
//type Expr struct {
//	leftValue  interface{}
//	operator   OperatorType
//	rightValue interface{}
//
//	exprState ExprState
//}
//
//func NewExpr() *Expr {
//	newExpr := &Expr{
//		exprState: ExprStateEmpty,
//	}
//	return newExpr
//}
//
//func (this *Cursor) SetExpr(expr *Expr) {
//	this.curExpr = expr
//}
//
//func (this *Cursor) GetExpr() *Expr {
//	return this.curExpr
//}
//
//func (this *Cursor) GetExprRes() (bool, error) {
//	exprRes, err := this.curExpr.Parse()
//	if err != nil {
//		panic(err)
//	}
//
//	return exprRes, nil
//}
//
//func (this *Expr) ToExprString() string {
//	// TODO: exclude uncomparable types like map, slice etc...
//	var stringBuilder strings.Builder
//	leftV, okLeft := this.leftValue.(*variables.Variable)
//	rightV, okRight := this.leftValue.(*variables.Variable)
//	if okLeft && okRight {
//		leftString, err := leftV.ToString()
//		if err != nil {
//			panic(err)
//		}
//
//		rightString, err := rightV.ToString()
//		if err != nil {
//			panic(err)
//		}
//		stringBuilder.WriteString(leftString)
//		stringBuilder.WriteString(string(this.operator))
//		stringBuilder.WriteString(rightString)
//	}
//	return stringBuilder.String()
//}
//
//func (this *Expr) Parse() (bool, error) {
//	if this.exprState != ExprStateRightSet {
//		return false, errors.New("Expr not full set")
//	}
//
//	exprString := this.ToExprString()
//
//	return false, errors.New("Unknown error")
//}
//
//func (this *Expr) PushValue(value interface{}) error {
//	if this.exprState == ExprStateRightSet {
//		return errors.New("Expr already full set")
//	}
//
//	if this.exprState == ExprStateEmpty {
//		this.SetLeft(value)
//	} else if this.exprState == ExprStateOpSet {
//		this.SetRight(value)
//	}
//
//	return nil
//}

//func (this *Expr) SetLeft(value interface{}) {
//	this.leftValue = value
//	this.exprState = ExprStateLeftSet
//}
//
//func (this *Expr) SetRight(value interface{}) {
//	this.rightValue = value
//	this.exprState = ExprStateRightSet
//}
//
//func (this *Expr) SetOperator(opType OperatorType) {
//	this.operator = opType
//	this.exprState = ExprStateOpSet
//}

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

func (this *Cursor) GetExpr() string {
	return this.curExpr
}

func (this *Cursor) PushExpr(exprComponent string) {
	this.curExpr += exprComponent
}

func (this *Cursor) ClearExpr() {
	this.curExpr = ""
}

func (this *Cursor) SetCursorContext(contextType ContextType) {
	this.curContext = contextType
}

func (this *Cursor) GetCursorContext() ContextType {
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
