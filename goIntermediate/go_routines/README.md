# `goroutine`

A `goroutine` is a lightweight execution thread in the Go programming language and a function that executes concurrently with the rest of the program.

Goroutines are incredibly cheap when compared to traditional threads as the overhead of creating a `goroutine` is very low.

Therefore, they are widely used in Go for concurrent programming.

To invoke a function as a `goroutine`, use the `go` keyword.

**_syntax:_**
We write `go` before the function foo to invoke it as a goroutine. The function `foo()` will run concurrently or asynchronously with the calling function

```
go foo()
```

**_example:_**
Let’s look at an example. We define a function foo that prints numbers from 1 to 3 along with the passed string. We add a delay using time.Sleep() inside the for loop. Without the delay, the first goroutine will finish executing even before the second one starts. The delay ensures that the goroutines are running concurrently before the results are shown.

```go
package main

import (
    "fmt"
    "time"
)

// Prints numbers from 1-3 along with the passed string
func foo(s string) {
    for i := 1; i <= 3; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s, ": ", i)
    }
}

func main() {

    // Starting two goroutines
    go foo("1st goroutine")
    go foo("2nd goroutine")

    // Wait for goroutines to finish before main goroutine ends
    time.Sleep(time.Second)
    fmt.Println("Main goroutine finished")
}
```

We also add a delay before the main goroutine ends so that the two started goroutines have time to finish.

> The main goroutine must be running for any other goroutines to run.
> If the main goroutine terminates, then the program will exit and no other goroutine will run.

As you can see from the output, the printing of both goroutines are interleaved, which reflects that the goroutines are running concurrently by the Go runtime.

> What if two goroutines want to communicate with each other? For that, Go has channels. Read more [here](https://www.educative.io/answers/what-are-channels-in-golang).

## Resources

[What is a goroutine? - educative.io](https://www.educative.io/answers/what-is-a-goroutine)
