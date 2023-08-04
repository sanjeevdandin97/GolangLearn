# Go pointers

Pointers are simply values which holds the memory address of another variables.

**_Example:_**

Let’s say `a = 5`.

This means the value of a is located in some random memory space, say `0x416038`.

This memory address can be seen printing out the value with ampersand(`&`) prefix.

```go
func main() {
    a := 5
    fmt.Println(&a) // prints 0x416038 in console
}
```

We can keep this memory space value to another variable `b` which literally means `b = 0x416038`.

Now, we can get value of a by dereferencing.

```go
func main() {
    a := 5
    fmt.Println(&a) // 0x416038
    b := &a
    fmt.Println(*b) // prints 5
}
```

## Why use pointers instead of value?

Its obvious we can assign value of `a` directly to `b` and make it easy.

```go
a := 5
b := a
fmt.Println(b) // 5
```

But, let’s see real scenario.

```go
func change(a *int) {
    fmt.Println("Address of a:",a)
    *a = 6
}

func nochange(a int){
    a = 7
}

func main(){
    b := 5
    fmt.Println("Address of b:",&b)
    change(&b)
    fmt.Println("Changed value:",b)
    nochange(b)
    fmt.Println("No change:",b)
}
```

Output would be

```
Address of b: 0x416020
Address of a: 0x416020
Changed value: 6
No change: 6
```

Here, instead of sending copy of `b`, we sent the address of `b` to change function and `a` changed the real value of `b` because both `a` and `b` have same address of their values.

Now, after calling the function `b`, we can later use the changed value in the `main` function leaving its original value `5` .

In the `nochange` function, we sent the value of `b` which creates copy on another memory location so it runs independently without affecting original value.

## Relating Array/Slice with Pointers

In C/C++, array variable is pointer to the first element of array.

```c
int arr[3] = {0,1,2};
print("%d", arr) // 1024
```

If the memory address of `arr[0]` is `1024`.

Then `arr[1]` is `1028` and so on but `arr` itself has value of the address of first element of array i.e., `1024`.

But in GoLang, the case is bit different. Go’s arrays are values rather than memory address.

```go
var myarr = [...]int{1,2,3}
fmt.Println(myarr) // [1 2 3]
fmt.Println(&myarr) // &[1 2 3]
```

But Go has introduced concept of **slice** which is more flexible and advanced than array. We can use built-in function `make` to create slice which takes `type(T — byte), length(len) and optional capacity(cap)`.

**_Syntax:_**

```go
func make([]T, len, cap) []T
```

The initialization structure is

```go
var sl []byte
sl = make([]byte, 5, 5)
fmt.Println(sl) // [0 0 0 0 0] //output
```

What we are going to focus on slice it how it allocate memory. Let’s go with an example:

```go
a := []byte{'w', 'x', 'y', 'z'}
b := a[2:] // b == []byte{'y', 'z'}
b[0] = 's' // b == []byte{'s', 'x'}
// a == []byte{'s', 'x', 'y', 'z'}
```

Here, slice didn’t create a new copy of `a` but created new slice value that points to the original array which modified parent array.

So, for the manipulation of array element, we don’t need to perform extra step.

Similarly, the size can be expanded easily by builtin `copy` function.

You can read more interesting things about slice and its internals from this [official blog](https://go.dev/blog/slices-intro) also.

## Resources

- [Deep Dive into Pointers, Arrays & Slice - Medium.com](<https://dwdraju.medium.com/deep-dive-into-pointers-arrays-slice-309a843c63ad#:~:text=Pointers%20are%20simply%20values%20which,with%20ampersand(%26)%20prefix.>)
- [Go Slices: usage and internals | The Go Blog - go.dev](https://go.dev/blog/slices-intro)
