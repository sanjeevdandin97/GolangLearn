package go_channels

import (
	"fmt"
	"time"
)

func INIT() {
	goChannels()
	goChannels1()
}

// Prints a greeting message using values received in
// the channel
func greet(c chan string) {
	name := <-c // receiving value from channel
	fmt.Println("Hello", name)
}

func goChannels() {

	// Making a channel of value type string
	c := make(chan string)

	// Starting a concurrent goroutine
	go greet(c)

	// Sending values to the channel c
	c <- "World"

	// Closing channel
	close(c)
}

func goChannels1() {
	numValues := []int{1, 2, 3, 4, 5}
	// newNumValues := []int{}

	for i := 0; i < len(numValues); i++ {
		c := make(chan int)

		go printNum(c)

		c <- numValues[i]

		close(c)
	}
}

func printNum(c chan int) {
	time.Sleep(100 * time.Millisecond)
	value := <-c
	fmt.Println(value)
}
