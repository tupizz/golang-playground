package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[2:4:4] // index 0 (inclusive) ~ index 2 (exclusive) and capacity of 2, full slice expression
	fmt.Println(slice, cap(slice))

	slice = []int{1, 2, 3, 4, 5}
	foo(slice)

	//	Array is passed by value, slice is passed by reference
}

func foo(slice []int) {
	_ = slice[4] // bounds check -> check if I have a slice of 5 elements which are being used inside the function, PERFORMANCE
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
	fmt.Println(slice[4])
}
