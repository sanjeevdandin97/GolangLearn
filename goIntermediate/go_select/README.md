# Go Select

The `select` statement is similar to switch cases but lets a `goroutine` wait on multiple communication operations.

A `select` blocks until one of its cases can run, then it executes that case.

It chooses one at random if multiple are ready.

**_Syntax:_**

```go
select {
    case 1:
        // code here
    ........
    case n:
        // code here
}
```

## Default Selection

The `default` case in a select is run if no other case is ready.

Use a `default` case to try a send or receive without blocking:

```go
select {
    case i := <-c:
        // use i
    default:
        // receiving from c would block
}
```

## Resources

- [Select | A Tour of Go -go.dev](https://go.dev/tour/concurrency/5)
- [Default Selection | A Tour of Go - go.dev](https://go.dev/tour/concurrency/6)
- [Will time.Tick cause memory leak when I never need to stop it? - Stackoverflow](https://stackoverflow.com/questions/68289916/will-time-tick-cause-memory-leak-when-i-never-need-to-stop-it)
