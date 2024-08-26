package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	if err != nil && errors.Is(err, SomeError) {
		println("main() Some Error happened:", err.Error())
		return
	}

	println("No error")
}

// No error printed
//func foo() error {
//	err := bar()
//	if err != nil {
//		return errors.New("foo: " + err.Error())
//	}
//	return nil
//}

// it is a good practice to wrap the error, never return a new error
func foo() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("foo() Some Error happened: %w", err) // %w is used to wrap the error
	}
	return nil
}

var SomeError = errors.New("error not found - 10123")

func bar() error {
	return SomeError
}
