package utils

import (
	"errors"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/parser"
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
		//if len(subTree) != 1 {
		//	return "", errors.New("More than 1 children ... ")
		//}

		if subTreeNode, ok := subTree.(*antlr.TerminalNodeImpl); ok {
			// fmt.Println("1111111", subTreeNode.GetText(), len(subTreeNode.GetText()))
			return subTreeNode.GetText(), nil
		} else {
			// fmt.Println("2222222")
			subTree = subTree.GetChildren()[0]
		}

	}
	// fmt.Println("33333333")

	return "", errors.New("Unknown error ... ")
}
