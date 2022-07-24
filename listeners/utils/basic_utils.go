package utils

import (
	"errors"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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
