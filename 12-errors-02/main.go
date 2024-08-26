package main

import "errors"

/*
*
Custom error
*/
type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string {
	return s.msg
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, SqrtError{"square root of negative number"}
	}
	return n, nil
}

func calculate() {
	x := -10
	result, err := squareRoot(float64(x))
	if err != nil {
		println("Error:", err.Error())
		return
	}

	println("Result:", result)
}

var ErrNotFound = errors.New("not found")

func foo() error {
	return ErrNotFound
}

func assertErrors() {
	err := foo()
	if err != nil && errors.Is(err, ErrNotFound) {
		println("Error:", err.Error())
		return
	}
	println("No error")
}

func main() {
	assertErrors()
}
