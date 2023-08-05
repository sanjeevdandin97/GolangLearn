# [![Go Back Image][1]](../) Go Channels

In Golang, or Go, channels are a means through which different `goroutines` communicate.

Think of them as pipes through which you can connect with different concurrent `goroutines`. The communication is bidirectional by default, meaning that you can send and receive values from the same channel.

Moreover, by default, channels send and receive until the other side is ready. This allows `goroutines` to synchronize without explicit locks or condition variables.

![Channels Image][2]

<p style="text-align: center;">
    Sending data from one goroutine to another
</p>

**_Syntax:_**

- _Make a channel:_ `make(chan [value-type])`, where `[value-type]` is the data type of the values to send and receive, e.g., `int`.

- _Send and receive values:_ `channel <-` and `<- channel`,
  where `<-` is the channel operator.

- _Close a channel:_ `close(channel)`.
  After closing, no value will be sent to the channel.

> Both sending and receiving are blocking operations by default.

**_Example:_**

- We make a channel with the type string and start function greet as a goroutine.
- The function greet is blocked when it encounters <- c and waits to receive a value.
- The main goroutine sends a value to the channel, which is then printed by the greet function.

```go
package main
import "fmt"

// Prints a greeting message using values received in
// the channel
func greet(c chan string) {

	name := <- c	// receiving value from channel
	fmt.Println("Hello", name)
}

func main() {

	// Making a channel of value type string
	c := make(chan string)

	// Starting a concurrent goroutine
	go greet(c)

	// Sending values to the channel c
	c <- "World"

	// Closing channel
	close(c)
}
```

## Resources

[What are channels in Golang? - educative.io](https://www.educative.io/answers/what-are-channels-in-golang)

[comment]: <> (Images Paths specified below)
[1]: ./../../assets/icons/go-back-arrow-svgrepo-com.svg
[2]: ./../../assets/icons/4671705944424448.svg
