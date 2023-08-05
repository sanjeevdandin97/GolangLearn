# Go Buffers

The `buffer` belongs to the `byte` package of the Go language and we can use these packages to manipulate the `byte` of the string.

For example, suppose we have a string. We can read the length of the string with the `len` function, which will return the numeric length.

But what if the strings are too large? We want to calculate in the form of chunks of data.

So in such situations, we can use the `buffer`.

The `buffer` allows us to handle any size of the dynamic length, making it flexible.

**_Syntax:_**

Given below is a simple syntax for the go language byte buffer.

We can see in the example, here after the creation of the variable `strBuffer` from the `bytes` package.

We can call it's many functions for further calculations.

```go
str Buffer bytes.Buffer
```

## How does a Buffer work in the Go language?

The buffer name itself clarifies its purposes, It allows us to give buffer storage where we can store some data and access the same data when needed.

In the case of the string being a large size, there are too many strings that we will store into a variable. Then we will create a buffer variable and keep storing the data onto that variable.

We will see the working of the buffer in go language with the help of the below points.

To use the buffer in the go language, we need to import the bytes package of the go language.

- Once we have imported the `bytes` package, we can create a variable with the byte package like `var x = bytes`. Buffer, and on the variable `x`, we can perform all the operations related to the buffering of `string`.
- We can store data of `string` onto the buffer variable `x` like `x.WriteString(“string of message ”)`, and the data which we stored on the string can be accessed like `x.String()`.
- We can check the length of the string stored on the buffer variable which we have created.
- Even we can store the bytes of data like `x.Write([]byte(“Hello “))`, and we can get the length of the stored value in the form of a number like `x.len()`.

## Examples

### **Example 1:**

**_Code:_**

```go
package main
import (
	"fmt"
)
//importing the bytes package so that buffer can be used
import (
	"bytes"
)
func main() {
//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Ranjan")
	strBuffer.WriteString("Kumar")
	fmt.Println("The string buffer output is",strBuffer.String())
}
```

**_Output:_**

```sh
The string buffer output is RanjanKumar
```

### **Example 2:**

**_Code:_**

```go
package main
import (
	"fmt"
)
//importing the bytes package so that buffer can be used
import (
	"bytes"
)

func main() {
	//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")
	strBuffer.WriteString("[DONE]")
	fmt.Println("The string buffer output is",strBuffer.String())
}
```

**_Output:_**

```sh
The string buffer output is It is number: 10, This is a string: Bridge
[DONE]
```

### **Example 3:**

**_Code:_**

```go
package main
import (
	"fmt"
)
//importing the bytes and os package so that buffer can be used on the os based
import (
	"bytes"
    "os"
)

func main() {
//Creating buffer variable to hold and manage the string data
   var byteString bytes.Buffer
    byteString.Write([]byte("Hello "))
    fmt.Fprintf(&byteString, "Hello friends how are you")
    byteString.WriteTo(os.Stdout)
}
```

**_Output:_**

```sh
Hello Hello friends how are you
```

### **Example 4:**

**_Code:_**

```go
package main
import (
	"fmt"
)
//importing the bytes package so that buffer can be used
import (
	"bytes"
)

func main() {
//Creating buffer variable to hold and manage the string data
   var strByyte bytes.Buffer
    strByyte.Grow(64)
    strByytestrByyte := strByyte.Bytes()
    strByyte.Write([]byte("It is a 64 byte"))
    fmt.Printf("%b", strByytestrByyte[:strByyte.Len()])
}
```

**_Output:_**

```sh
[1001001 1110100 100000 1101001 1110011 100000 1100001 100000 110110 110100 100000 1100010 1111001 1110100 1100101]
```

### **Example 5:**

**_Code:_**

```go
package main
import (
	"fmt"
)
//importing the bytes package so that buffer can be used
import (
	"bytes"
)
func main() {
//Creating buffer variable to hold and manage the string data
   var strByyte bytes.Buffer
    strByyte.Grow(64)
    strByyte.Write([]byte("kumar"))
    fmt.Printf("%b", strByyte.Len())
}
```

**_Output:_**

```sh
101
```

### **Example 6:**

**_Code:_**

```go
package main
import (
	"fmt"
)
//importing the bytes package so that buffer can be used
import (
	"bytes"
)
func main() {
//Creating buffer variable to hold and manage the string data
   var strByyte = []byte("Ranjan, Kumar")
    strByyte = bytes.TrimPrefix(strByyte, []byte("Hello, "))
    strByyte = bytes.TrimPrefix(strByyte, []byte("Good we will meet again,"))
    fmt.Printf("Hello%s", strByyte)
}
```

**_Output:_**

```sh
HelloRanjan, Kunar
```

## Resources

- [GoLang Buffer - educba.com](https://www.educba.com/golang-buffer/)
- [Bytes Package | Go Packages - pkg.go.dev](https://pkg.go.dev/bytes)
