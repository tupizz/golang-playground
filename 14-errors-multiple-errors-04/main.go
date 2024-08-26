package main

import (
	"errors"
	"fmt"
)

var (
	SomeError  = errors.New("error not found - 10123")
	SomeError2 = errors.New("error not found - 10124")
)

func main() {
	error := foo()
	fmt.Println(error)
	fmt.Println(errors.Is(error, SomeError))  // true
	fmt.Println(errors.Is(error, SomeError2)) // true
}

func a() error { return SomeError }
func b() error { return SomeError2 }

// Return multiple errors
func foo() error {
	var errorResult error

	if err := a(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	if err := b(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	return errorResult
}
