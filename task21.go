package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func task21() {
	animals := []Animal{
		Dog{Name: "Lucky"},
		Cat{Name: "Tama"},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
