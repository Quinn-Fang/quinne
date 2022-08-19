package utils

import (
	"errors"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/Quinn-Fang/quinne/parser"
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
