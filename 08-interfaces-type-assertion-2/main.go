package main

import "fmt"

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

// function that uses the interface
func WhatDoesTheAnimalSaying(a Animal) string {
	return a.Sound()
}

func main() {
	slice := []Animal{
		Dog{},
		Cat{},
	}

	// this is called OOP polymorphism
	for i, v := range slice {
		fmt.Println(i, WhatDoesTheAnimalSaying(v))
	}

}
