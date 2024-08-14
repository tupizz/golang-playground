package main

import "fmt"

// pointer -> address in memory
// *pointer -> value at address in memory
// &variable -> address in memory of variable
// pointers point to the stack memory
func main() {
	x := 10
	p := &x

	fmt.Println("Value of x")
	fmt.Println(x) // e.g. 10

	fmt.Println("Address memory value from pointer")
	fmt.Println(p) // e.g. 0xc0000b0008

	fmt.Println("Dereferencing the pointer")
	fmt.Println(*p) // e.g. 10

	take(&x) // address in memory of variable x
	fmt.Println("Value of x after take")
	fmt.Println(x) // e.g. 20

	y := create() // this is returning a pointer to the address in memory of y
	fmt.Println("Value of y")
	fmt.Println(*y) // Dereferencing the pointer to get the value at the address in memory

	fmt.Println("Nil value")
	var nilValue *int
	take(nilValue)
	fmt.Println(nilValue) // nil

}

// When we see *int or * being used with a type, this means that we're dealing with a pointer of a given type.
// But when we see the same * used with a variable like x for example, this means we are dereferencing the value of the pointer, and then accessing the value.
func take(x *int) {
	// This is very important to do, because golang is a language doesn't provide Nil pointer safety
	// this only happens in runtime, not build :(
	if x == nil {
		fmt.Println("x is nil")
		return
	}

	*x = *x + 10
}

func create() *int {
	x := 100
	return &x
}

// IMPORTANT:
// panic: runtime error: invalid memory address or nil pointer dereference
// in some place you made a dereferencing of a nil pointer in golang
