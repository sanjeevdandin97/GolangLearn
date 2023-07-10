package goLoops

import "fmt"

var b int = 15
var a int
var numbers = [6]int{1, 2, 3, 5}

func RunGoLoops(value string) {
	checkLoopType(value)
}

func checkLoopType(value string) {
	if value == "LOOP_BASE" {
		callLoopBase()
	} else if value == "LOOP_COND" {
		callLoopCond()
	} else if value == "LOOP_RNGE" {
		callLoopRange()
	}
}

func callLoopBase() {
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
}

func callLoopCond() {
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
}

func callLoopRange() {
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
}
