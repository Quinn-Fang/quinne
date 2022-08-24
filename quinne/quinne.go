package quinne

import (
	"github.com/Quinn-Fang/quinne/listeners"
	"github.com/Quinn-Fang/quinne/navigator"
	"github.com/Quinn-Fang/quinne/parser"
	"github.com/Quinn-Fang/quinne/procedures/buildin"
	"github.com/Quinn-Fang/quinne/sym_tables"
	"github.com/Quinn-Fang/quinne/uspace"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type EventHandler struct {
	uNavigator *navigator.Navigator
}

func newEventHandler() *EventHandler {
	eventHandler := &EventHandler{}
	return eventHandler
}

func (this *EventHandler) setUNavigator(curNavigator *navigator.Navigator) {
	this.uNavigator = curNavigator
}

func (this *EventHandler) setCurNavigator() {
	curNavigator := navigator.GetCurNavigator()
	this.setUNavigator(curNavigator)
}

func NewEventHandler(fileName string) *EventHandler {
	runListener(fileName)

	handler := newEventHandler()
	handler.setUNavigator(navigator.GetCurNavigator())
	return handler
}

func (this *EventHandler) GetNextEvent() (*uspace.Event, error) {
	if this.uNavigator == nil {
		panic("Navigator uninitialized")
	}
	event, err := this.uNavigator.GetNextEvent()

	if err == nil {
		if _, ok := event.GetEventContext().(*sym_tables.IfElseBranch); ok {
			event.GetSymTable().IsExecutable()
			return event, nil
		}
		executable := event.GetSymTable().IsExecutable()
		event.GetSymTable().SetExecutable(executable)
		if !executable {
			return this.GetNextEvent()
		} else {
			return event, nil
		}

	} else {
		return nil, err
	}
}

func runListener(fileName string) {
	input, _ := antlr.NewFileStream(fileName)
	// Create First SymTable
	sym_tables.NewEntryTable()
	// Create Buildin Function Table
	buildin.InitFuncTable()

	// Create Cursor
	navigator.InitCursor()
	curNavigator := navigator.NewNavigator()
	navigator.SetCurNavigator(curNavigator)

	// Create the Lexer
	lexer := parser.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGoParser(stream)

	//listen := GoListener{}
	tree := p.SourceFile()
	antlr.ParseTreeWalkerDefault.Walk(listeners.NewGoListener(p, tree), tree)
	// utils.PrintAllSymTale()
	//curNavigator.PrintStack()

	//curNavigator.PrintCodeSegments()
}

func TNewListener(fileName string) {
	input, _ := antlr.NewFileStream(fileName)
	// Create First SymTable
	sym_tables.NewEntryTable()
	// Create Cursor
	navigator.InitCursor()
	curNavigator := navigator.NewNavigator()
	navigator.SetCurNavigator(curNavigator)

	// Create the Lexer
	lexer := parser.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGoParser(stream)

	//listen := GoListener{}
	tree := p.SourceFile()
	antlr.ParseTreeWalkerDefault.Walk(listeners.NewGoListener(p, tree), tree)
	// utils.PrintAllSymTale()
	//curNavigator.PrintStack()

	//curNavigator.PrintCodeSegments()

}
