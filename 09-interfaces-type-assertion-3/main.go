package main

func takeEmptyInterface(a any) {
	str, ok := a.(string) // check in runtime if a is from string type
	if !ok {
		println("not a string")
		return
	}

	println(str)
}

type Animal interface {
	Sound() string
}
type Dog struct{}

func (d Dog) Sound() string {
	return "woof"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "meow"
}

func takeAnimal(a Animal) {
	switch t := a.(type) {
	case Dog:
		println("Dog says", t.Sound())
	case Cat:
		println("Cat says", t.Sound())
	}

	// Another way:
	//switch a.(type) {
	//case Dog:
	//	println("Dog say", a.Sound())
	//case Cat:
	//	println("Cat say", a.Sound())
	//}
}

func main() {
	takeEmptyInterface("string")
	takeEmptyInterface("10")
	takeEmptyInterface(10)
	// passing a struct
	takeEmptyInterface(struct{ Name string }{
		Name: "Tadeu",
	})

	bigCat := Cat{}
	takeAnimal(bigCat)

	bigDog := Dog{}
	takeAnimal(bigDog)
}
