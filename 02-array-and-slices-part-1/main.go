package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	// slice is a pointer to the array
	slice := arr[0:5] // index 0 (inclusive) ~ index 5 (exclusive)
	fmt.Println(slice)

	// since slice is a pointer to the array if I change the array the slice will change as well
	fmt.Println("Changing the array")
	arr[0] = 100
	arr[1] = 200
	fmt.Println(arr)
	fmt.Println(slice)

	// Declaring a slice literal
	sliceLiteral := []int{1, 2, 3} // without the length [10]int{} and without the index [0:5]
	fmt.Println(sliceLiteral)
	fmt.Println("\n\nLength of sliceLiteral", len(sliceLiteral))
	fmt.Println("Capacity of sliceLiteral", cap(sliceLiteral))

	sliceLiteral = append(sliceLiteral, 4)
	fmt.Println("\n\nLength of sliceLiteral", len(sliceLiteral))
	fmt.Println("Capacity of sliceLiteral", cap(sliceLiteral))
	fmt.Println(sliceLiteral)

	sliceLiteral = append(sliceLiteral, 5)
	sliceLiteral = append(sliceLiteral, 6)
	sliceLiteral = append(sliceLiteral, 7)
	sliceLiteral = append(sliceLiteral, 8)
	fmt.Println("\n\nLength of sliceLiteral", len(sliceLiteral))
	fmt.Println("Capacity of sliceLiteral", cap(sliceLiteral))
	fmt.Println(sliceLiteral)
	//	IMP: the capacity of the slice is doubled every time it reaches the limit

	var newSlice []int // this is different then sliceLiteral := []int{} because this is nil
	fmt.Println("\n\nLength of newSlice", len(newSlice))
	fmt.Println("Capacity of newSlice", cap(newSlice))
	fmt.Println(newSlice == nil)

}
