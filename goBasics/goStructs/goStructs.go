package goStructs

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary int
}

var person = Person{
	Name:   "John Doe",
	Age:    100,
	Job:    "retired",
	Salary: 0,
}

func ImplicitCallPrintPerson() {
	PrintPerson(&person)
}

func ExpilcitCallPrintPerson() {
	fmt.Println("Enter Person Name: ")
	fmt.Scanln(&person.Name)

	fmt.Println("Enter Person Age: ")
	fmt.Scanln(&person.Age)

	fmt.Println("Enter Person Job: ")
	fmt.Scanln(&person.Job)

	fmt.Println("Enter Person Salary: ")
	fmt.Scanln(&person.Salary)

	PrintPerson(&person)
}

func PrintPerson(person *Person) {
	fmt.Println("Name: ", person.Name)
	fmt.Println("Age: ", person.Age)
	fmt.Println("Job: ", person.Job)
	fmt.Println("Salary: ", person.Salary)
}
