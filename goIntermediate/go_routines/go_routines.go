package go_routines

import (
	"fmt"
	"time"
)

func INIT() {
	exec()
}

// Prints numbers from 1-3 along with the passed string
func goRoutines(s string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, ": ", i)
	}
}

func exec() {

	// Starting two goroutines
	go goRoutines("1st goroutine")
	go goRoutines("2nd goroutine")

	// Wait for goroutines to finish before main goroutine ends
	time.Sleep(time.Second)
	fmt.Println("Main goroutine finished")
}
