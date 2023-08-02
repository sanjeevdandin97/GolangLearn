package go_interfaces

import "fmt"

func INIT() {
	go_interfaces()
}

// interface
type Shape interface {
	area() float32
	getName() string
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
	name            string
}

type Triangle struct {
	base, height float32
	name         string
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func (t Triangle) area() float32 {
	return (t.base * t.height) / 2
}

func (r Rectangle) getName() string {
	return r.name
}

func (t Triangle) getName() string {
	return t.name
}

// access method of the interface
func calculate(s Shape) {
	fmt.Println("Area of", s.getName(), ":", s.area())
}

// main function
func go_interfaces() {

	// assigns value to struct members
	rect := Rectangle{7, 4, "Rectangle"}
	tri := Triangle{10, 5, "Triangle"}

	// call calculate() with struct variable rect
	calculate(rect)
	calculate(tri)
}
