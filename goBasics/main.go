package main

import (
	"fmt"
	"goBasics/goErrorHandling"
	"goBasics/goLoops"
	"goBasics/goPrint"
	"goBasics/goReturnType"
	"goBasics/goStdIO"
	"goBasics/goStructs"
	"goBasics/goTypeConversions"
	"os"
)

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
	} else if systemArgument == "TC" {
		// Example for GO Type Conversions
		goTypeConversions.GoTypesConversions()
	} else if systemArgument == "STRCT_IMPLICIT" {
		// Example for GO Struct
		goStructs.ImplicitCallPrintPerson()
	} else if systemArgument == "STRCT_EXPLICIT" {
		// Example for GO Struct
		goStructs.ExpilcitCallPrintPerson()
	} else if systemArgument == "LOOP_BASE" || systemArgument == "LOOP_COND" || systemArgument == "LOOP_RNGE" {
		// Example for GO Loops
		goLoops.RunGoLoops(systemArgument)
	} else if systemArgument == "ERRH1" || systemArgument == "ERRH2" || systemArgument == "ERRH3" || systemArgument == "ERRH4" {
		// Example for Go Error Handling
		goErrorHandling.GoErrorHandling(systemArgument)
	}
}
