# Go Structures

A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.

**_Syntax:_** To declare a structure in Go, use the `type` and `struct` keywords

```
type _struct_name_ struct {
_member1_  _datatype_;
_member2_  _datatype_;
_member3_  _datatype_;
...
}
```

**_Example:_** Here we declare a struct type `Person` with the following members: `name`, `age`, `job` and `salary`:

```
type Person struct {
name string
age  int
job string
salary  int
}
```

**Tip:** Notice that the struct members above have different data type

## Resources

- [Go Struct - w3schoools](https://www.w3schools.com/go/go_struct.php)
- [Structures in Golang - GeeksforGeeks](https://www.geeksforgeeks.org/structures-in-golang/)
