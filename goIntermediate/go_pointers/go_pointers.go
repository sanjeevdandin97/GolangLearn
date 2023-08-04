package go_pointers

import "fmt"

func INIT() {
	pointerVariable()
}

func change(a *int) {
	fmt.Println("Address of a:", a)
	*a = 6
}

// ERROR: argument a is overwritten before first use (SA4009)
func nochange(a int) {
	a = 7
}

func pointerVariable() {
	b := 5
	fmt.Println("Address of b:", &b)
	change(&b)
	fmt.Println("Changed value:", b)
	nochange(b)
	fmt.Println("No change:", b)
}
