package utils

import (
	"fmt"

	"quinn007.com/sym_tables"
)

func PrintAllSymTale() {
	rootTable := sym_tables.GetRootSymTale()
	PrintSymbolTable(rootTable, 1)
}

func PrintSymbolTable(symTable *sym_tables.SymTable, index int) {
	if symTable == nil {
		return
	}
	fmt.Printf("************** SymTale: %d *********************\n", index)
	fmt.Println("*************** Variables : ******************")
	for k, v := range symTable.GetVariables() {
		fmt.Println(k)
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("*************** Functions : ******************")
	curFunctions := symTable.GetFunctions()
	for _, v := range curFunctions {
		fmt.Printf("%+v ", v)
		for _, v1 := range v.GetParams() {
			fmt.Printf("%+v ", v1)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for _, child := range symTable.GetChildren() {
		index++
		PrintSymbolTable(child, index)
	}
}
