package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greeting() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p Person) ChangeName(newName string) *Person {
	p.Name = newName
	return &p
}

func task16() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.Greeting())
	fmt.Println(p.ChangeName("Bob").Greeting())
}
