package goPrint

import "fmt"

/*
Go has three functions to output text:

1. Print():
   prints its arguments with their default format.
2. Println():
   The Println() function is similar to Print() with the difference that a whitespace is added between the arguments,
   and a newline is added at the end
3. Printf():
   The Printf() function first formats its argument based on the given formatting verb and then prints them.

   Here we will use two formatting verbs:

   %v is used to print the value of the arguments
   %T is used to print the type of the arguments
   For more foramtting verbs checkout this link:
   https://www.w3schools.com/go/go_formatting_verbs.php
*/

func goPrint() {
	var i, j string = "Hello", "World"
	fmt.Print("PRINT METHOD:", '\n')

	// Without new Line Arguments
	fmt.Print(i)
	fmt.Print(j)

	// With new Line Arguments
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	// Multiple Argument in single print
	fmt.Print(i, "\n", j, "\n")

	// Inserts a space between each argument if neither are strings
	fmt.Print(10, 20) // 10 20
	fmt.Print("\n")   // 10 20
}

func goPrintln() {
	var i, j string = "Hello", "World"

	fmt.Println("PRINTLN METHOD:")
	fmt.Println(i, j)
}

func goPrintF() {
	var i string = "Hello"
	var j int = 15

	fmt.Println("PRINTF METHOD")
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}

func GoPrintValues() {
	goPrint()
	goPrintln()
	goPrintF()
}
