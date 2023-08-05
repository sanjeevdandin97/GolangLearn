# Go Marshalling and Unmarshalling

Data encoding is an important part of any programming language to encode data into `JSON` or decode data into String.

In Golang, struct data is converted into `JSON` and `JSON` data to string with `Marshal()` and `Unmarshal()` method.

The methods returned data in byte format and we need to change returned data into `JSON` or `String`.

## What is Marshalling in Golang

Converting Go objects into `JSON` is called marshalling in Golang.

As the `JSON` is a language-independent data format, the Golang has inbuilt `encoding/json` package to perform `JSON` related operations.

The `encoding/json` package has `json.Marshal()` method. This method takes object as parameter and returned byte code.

**_Example:_**
In below example, we have Employee struct defined.

```go
type Employee struct {
	Name  string
	Age	int
	Address string
}
```

As we will convert an instance of `Employee` struct into JSON, so we will use `encoding/json` Golang package to use `json.Marshal()` method to convert data.

```go
import (
	"encoding/json"
)
```

Now we will create object of `Employee` struct and pass with data and convert Struct Object into Byte data.

Then we will convert Byte data into JSON string to display data.

```go
emp := Employee{Name: "George Smith", Age: 30, Address: "Newyork, USA"}
empData, err := json.Marshal(emp)
fmt.Println(string(empData))
```

Here is complete program to Marshal `Employee` struct into JSON.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name  string
	Age	int
	Address string
}

func main() {
	emp := Employee{Name: "George Smith", Age: 30, Address: "Newyork, USA"}
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(empData))
}
```

Now, when we run above program, we should see the following JSON output:

```go
{"Name":"George Smith","Age":30,"Address":"Newyork, USA"}
```

## What is Unmarshalling in Golang

Unmarshalling just opposite of Marshalling.

The Golang `encoding/json` package has `json.Unmarshal()` method that used to convert json(Byte data) into Struct.

As we have covered Marshalling of struct into JSON, now we will take that JSON string and Unmarshal that JSON into a Struct.

Here in this example, we have employee JSON string.

```go
empJsonData := `{"Name":"George Smith","Age":30,"Address":"Newyork, USA"}
```

We will create `Response` Struct that we will use to match byte code after Unmarshal to display data.

```go
type Response struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Address string `json:"address"`
}
```

Now we will create JSON string into byte code.

The we will Unmarshal byte code using `json.Unmarshal()` method and assign into response variable to display data.

```go
empJsonData := `{"Name":"George Smith","Age":30,"Address":"Newyork, USA"}`
empBytes := []byte(empJsonData)
var resp Response
json.Unmarshal(empBytes, &resp)
fmt.Println(resp.Name)
fmt.Println(resp.Age)
fmt.Println(resp.Address)
```

Here is complete program to Unmarshal JSON byte code.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Address string `json:"address"`
}

func main() {
	empJsonData := `{"Name":"George Smith","Age":30,"Address":"Newyork, USA"}`
	empBytes := []byte(empJsonData)
	var emp Response
	json.Unmarshal(empBytes, &emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
	fmt.Println(emp.Address)
}
```

Now, when we run above program, we should see the following output:

```sh
George Smith
30
Newyork, USA
```

## Resources

- [Marshalling and Unmarshalling in Golang - webdamn.com](https://webdamn.com/marshal-and-unmarshal-in-golang/)
