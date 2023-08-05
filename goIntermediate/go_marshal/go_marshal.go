package go_marshal

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
}

func INIT() {
	emp := Employee{Name: "George Smith", Age: 30, Address: "Newyork, USA"}

	fmt.Println(`Marshalling!`)
	marshalledEmp := marshalling(&emp)

	fmt.Println(`UnMarshalling!`)
	unMarshalling(marshalledEmp)
}

// MARSHALLING Data
func marshalling(emp *Employee) string {
	empData, err := json.Marshal(emp)
	if err != nil {
		panic(err)
	}
	convertEmpData := string(empData)

	fmt.Println(convertEmpData)

	return convertEmpData
}

// UNMARSHALLING Data
func unMarshalling(marshalledEmp string) {
	empBytes := []byte(marshalledEmp)
	var emp Employee

	json.Unmarshal(empBytes, &emp)

	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
	fmt.Println(emp.Address)
}
