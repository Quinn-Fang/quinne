package utils

import (
	"errors"
	"fmt"

	"github.com/Quinn-Fang/quinne/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func IsFunction(children []antlr.Tree) bool {
	if len(children) == 2 {
		if _, ok := children[0].(*parser.PrimaryExprContext); ok {
			if _, ok := children[1].(*parser.ArgumentsContext); ok {
				return true
			}
		}
	}
	return false
}

func GetTerminalNodeText(antlrTree antlr.Tree) (string, error) {
	subTree := antlrTree
	for true {
		if subTreeNode, ok := subTree.(*antlr.TerminalNodeImpl); ok {
			return subTreeNode.GetText(), nil
		} else {
			subTree = subTree.GetChildren()[0]
		}

	}

	return "", errors.New("Unknown error ... ")
}

func PrintChildren(children []antlr.Tree) {
	for _, v := range children {
		switch parserContext := v.(type) {
		case *antlr.TerminalNodeImpl:
			{
				terminalString, _ := GetTerminalNodeText(parserContext)
				fmt.Printf("Terminal text : %+v\n", terminalString)
			}
		default:
			{
				fmt.Printf("--- %T ---\n", v)
				PrintChildren(v.GetChildren())
			}
		}
	}
}
