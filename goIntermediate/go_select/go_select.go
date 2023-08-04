package go_select

import (
	"fmt"
	"time"
)

func INIT() {
	goSelect()
	goSelectDefault()
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func goSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func goSelectDefault() {
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()

	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick.C:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
