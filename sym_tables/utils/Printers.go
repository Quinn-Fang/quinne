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
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for _, v := range symTable.GetFunctions() {
		fmt.Printf("%+v\n", v)
	}
	for _, child := range symTable.GetChildren() {
		index++
		PrintSymbolTable(child, index)
	}
}
