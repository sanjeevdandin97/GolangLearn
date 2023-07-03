package main

import (
	"fmt"
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

	for i := 1; i < len(os.Args); i++ {
		systemArgument := os.Args[i]
		fmt.Println(i, ".CURRENT METHOD:", systemArgument)

		if systemArgument == "RT" {
			// Print Return types
			GoReturnTypePrint()
		} else if systemArgument == "PRNT" {
			// Example of PRINT Statements
			GoPrintValues()
		}
		fmt.Println()
		fmt.Println()
	}
}
