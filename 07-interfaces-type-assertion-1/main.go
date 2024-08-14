package main

func foo(a interface{}) {
	_ = a.(int) // type assertion
}

// any is an alias for interface{}
func bar(a any) {
	_ = a.(string) // type assertion
}

func main() {
	bar("string")
	foo(1)
}
