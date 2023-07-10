# GO Loops

A loop is a repetition control structure.

It allows you to write a loop that needs to execute a specific number of times.

**_Syntax_**:

```
for [condition |  ( init; condition; increment ) | Range] {
   statement(s);
}
```

**_Example_**:

Declare values

```
var b int = 15
var a int
numbers := [6]int{1, 2, 3, 5}
```

For the above declared values we can run for loop in three different ways:

- `for` loop `[( init; condition; increment )]`

  ```
  for a := 0; a < 10; a++ {
      fmt.Printf("value of a: %d\n", a)
  }
  ```

- `for` loop `[condition]`

  ```
  for a < b {
      a++
      fmt.Printf("value of a: %d\n", a)
  }
  ```

- `for` loop `[range]`

  ```
  for i,x:= range numbers {
      fmt.Printf("value of x = %d at %d\n", x,i)
  }
  ```

## Resources

[Go - for Loop | TutorialsPoint](https://www.tutorialspoint.com/go/go_for_loop.htm)
