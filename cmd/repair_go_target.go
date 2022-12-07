package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RepairGoCode() {
	// lines := []string{"4106", "17100", "17150", "17201"}
	lines := []string{"4138", "17154", "17957", "18008"}

	for _, line := range lines {
		// cmdString := "sed -i \"%ss/closingBracket/p.closingBracket/\" ./parser/go_parser.go"
		opt1 := "-i"
		opt2 := "%ss/closingBracket/p.closingBracket/"
		opt2 = fmt.Sprintf(opt2, line)
		opt3 := "parser/go_parser.go"
		fmt.Println("sed", opt1, opt2, opt3)
		cmd := exec.Command("sed", opt1, opt2, opt3)
		cmd.Dir = "/home/quinn/workspace/PL-Framework/formal_1/quinne"
		// cmd.Dir = "../."
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			panic(err)
		} else {
			fmt.Println("Successfully repaired go target code ...")
		}
	}
}
