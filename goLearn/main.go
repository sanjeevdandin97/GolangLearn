package main

import (
	"fmt"
	"goLearn/goPrint"
	"goLearn/goReturnType"
	goStdIO "goLearn/goStdInputOutput"
	"os"
)

// SINGLE LINE COMMENT
/* MULTI LINE COMMENTS */

/*
* command `go run . RT PRNT`
* `RT` = Return Type
* `PRNT` = Terminal Print/Console log Statements
 */
func main() {
	runLoop()
}

func runLoop() {
	for i := 1; i < len(os.Args); i++ {
		systemArgument := os.Args[i]

		fmt.Println(i, ".CURRENT METHOD:", systemArgument)

		checkSysArg(systemArgument)

		fmt.Println()
		fmt.Println()
	}
}

func checkSysArg(systemArgument string) {
	if systemArgument == "RT" {
		// Print Return types
		goReturnType.GoReturnTypePrint()
	} else if systemArgument == "PRNT" {
		// Example of PRINT Statements
		goPrint.GoPrintValues()
	} else if systemArgument == "STDIO" {
		// Example for Standard Input output statements
		goStdIO.Gostdinputoutput()
	}
}
