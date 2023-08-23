# Go Error Handling

Error handling in Go is a little different than other mainstream programming languages like Java, JavaScript, or Python.

Go’s built-in errors don’t contain stack traces, nor do they support conventional `try`/`catch` methods to handle them.

Instead, errors in Go are just values returned by functions and they can be treated in much the same way as any other datatype leading to a surprisingly lightweight and simple design.

**_Example:_**
The `os.Open` function returns a non-nil `error` value when it fails to open a file.

```go
func Open(name string) (file *File, err error)
```

The following code uses `os.Open` to open a file.

If an error occurs it calls `log.Fatal` to print the error message and stop.

```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

You can get a lot done in Go knowing just this about the error type, but in this document we’ll take a closer look at error and discuss some good practices for error handling in Go.

## The Error Type

The `error` type in Go is implemented as the following `interface`:

```go
type error interface {
    Error() string
}
```

So basically, an error is anything that implements the `Error()` method, which returns an error message as a `string`. It’s that simple!

## Constructing Errors

Errors can be constructed on the fly using Go’s built-in errors or fmt packages.

**_Example:_**
The following function uses the errors package to return a new error with a static error message.

```go
package main

import "errors"

func DoSomething() error {
    return errors.New("something didn't work")
}
```

**_Example:_**
Similarly, the `fmt` package can be used to add dynamic data to the `error`, such as an `int`, `string`, or another `error`.

```go
package main

import "fmt"

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("can't divide '%d' by zero", a)
    }
    return a / b, nil
}
```

Note that `fmt.Errorf` will prove extremely useful when used to wrap another error with the `%w` format verb - but I’ll get into more detail on that further down in the article.

There are a few other important things to note in the example above.

- Errors can be returned as `nil`, and in fact, it’s the default, or “zero”, value of on error in Go. This is important since checking if `err != nil` is the idiomatic way to determine if an error was encountered (replacing the `try`/`catch` statements you may be familiar with in other programming languages).

- Errors are typically returned as the last argument in a function. Hence in our example above, we return an `int` and an `error`, in that order.

- When we do return an `error`, the other arguments returned by the function are typically returned as their default “zero” value. A user of a function may expect that if a non-nil `error` is returned, then the other arguments returned are not relevant.

- Lastly, `error` messages are usually written in lower-case and don’t end in punctuation. Exceptions can be made though, for example when including a proper noun, a function name that begins with a capital letter, etc.

## Defining Expected Errors

Another important technique in Go is defining expected Errors so they can be checked for explicitly in other parts of the code.

This becomes useful when you need to execute a different branch of code if a certain kind of error is encountered.

## Defining Sentinel Errors

Building on the Divide function from earlier, we can improve the error signaling by pre-defining a “Sentinel” error.

**_Example:_**
Calling functions can explicitly check for this error using `errors.Is`.

```go
package main

import (
    "errors"
    "fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}

func main() {
    a, b := 10, 0
    result, err := Divide(a, b)
    if err != nil {
        switch {
        case errors.Is(err, ErrDivideByZero):
            fmt.Println("divide by zero error")
        default:
            fmt.Printf("unexpected division error: %s\n", err)
        }
        return
    }

    fmt.Printf("%d / %d = %d\n", a, b, result)
}
```

## Defining Custom Error Types

Many error-handling use cases can be covered using the strategy above.

However, there can be times when you might want a little more functionality.

Perhaps you want an error to carry additional data fields, or maybe the `error`’s message should populate itself with dynamic values when it’s printed.

You can do that in Go by implementing custom errors type.

Below is a slight rework of the previous example.

Notice the new type `DivisionError`, which implements the `Error` `interface`.

We can make use of `errors.As` to `check` and `convert` from a standard error to our more specific `DivisionError`.

```go
package main

import (
    "errors"
    "fmt"
)

type DivisionError struct {
    IntA int
    IntB int
    Msg  string
}

func (e *DivisionError) Error() string {
    return e.Msg
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, &DivisionError{
            Msg: fmt.Sprintf("cannot divide '%d' by zero", a),
            IntA: a, IntB: b,
        }
    }
    return a / b, nil
}

func main() {
    a, b := 10, 0
    result, err := Divide(a, b)
    if err != nil {
        var divErr *DivisionError
        switch {
        case errors.As(err, &divErr):
            fmt.Printf("%d / %d is not mathematically valid: %s\n",
              divErr.IntA, divErr.IntB, divErr.Error())
        default:
            fmt.Printf("unexpected division error: %s\n", err)
        }
        return
    }

    fmt.Printf("%d / %d = %d\n", a, b, result)
}
```

> **Note:** when necessary, you can also customize the behavior of the `errors.Is` and `errors.As`. See [this Go.dev blog][1] for an example.

> **Another note:** `errors.Is` was added in Go 1.13 and is preferable over checking `err == ....` More on that below.

## Wrapping Errors

In these examples so far, the errors have been created, returned, and handled with a single function call.

In other words, the stack of functions involved in “bubbling” up the error is only a single level deep.

Often in real-world programs, there can be many more functions involved - from the function where the error is produced, to where it is eventually handled and any number of additional functions in-between.

In Go 1.13, several new error APIs were introduced, including `errors.Wrap` and `errors.Unwrap`, which are useful in applying additional context to an error as it “bubbles up”, as well as checking for particular error types, regardless of how many times the error has been wrapped.

> **_A bit of history:_**
>
> - Before Go 1.13 was released in 2019, the standard library didn’t contain many APIs for working with errors - it was basically just `errors.New` and `fmt.Errorf`.
> - As such, you may encounter legacy Go programs in the wild that do not implement some of the newer error APIs.
> - Many legacy programs also used 3rd-party error libraries such as pkg/errors.
> - Eventually, a formal proposal was documented in 2018, which suggested many of the features we see today in Go 1.13+.

## Resources

- [Error handling and Go | The Go Blog - go.dev](https://go.dev/blog/error-handling-and-go)
- [Effective Error Handling in Golang - earthly.dev](https://earthly.dev/blog/golang-errors/)

[1]: https://go.dev/blog/go1.13-errors
