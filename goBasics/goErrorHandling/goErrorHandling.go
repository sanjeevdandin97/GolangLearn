package goErrorHandling

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var fileName = "filename.ext"
var ErrDivideByZero = errors.New("divide by zero")

func GoErrorHandling(systemArgument string) {
	if systemArgument == "ERRH1" {
		basicErrorHandling()
	} else if systemArgument == "ERRH2" {
		fmt.Println(constructErrors())
	} else if systemArgument == "ERRH3" {
		fmt.Println(divide(1, 0))
	} else if systemArgument == "ERRH4" {
		sentinelError(10, 0)
	}
}

// ERRH1: BASIC ERROR HANDLING
func basicErrorHandling() *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the open *File f
	return f
}

// ERRH2: CONSTRUCTING ERRORS 1
func constructErrors() error {
	return errors.New("something didn't work")
}

// ERRH3: CONSTRUCTING ERRORS 2: DIVIDING EXCEPTION
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return a / b, nil
}

// ERRH4: DEFINING SENTINEL ERRORS
func sentinelDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func sentinelError(a, b int) {
	result, err := sentinelDivide(a, b)
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
