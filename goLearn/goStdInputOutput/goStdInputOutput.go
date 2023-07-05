package goStdIO

import "fmt"

var first string
var second string

func Gostdinputoutput() {
	enterUserInput()
	printOutput()
}

func enterUserInput() {
	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your First Name: ")

	// Taking input from user
	fmt.Scanln(&first)

	fmt.Println("Enter Second Last Name: ")

	fmt.Scanln(&second)
}

func printOutput() {
	// Print function is used to
	// display output in the same line
	fmt.Print("Your Full Name is: ")

	// Addition of two string
	fmt.Print(first + " " + second)
}
