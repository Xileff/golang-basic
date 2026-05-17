package main

import "fmt"

// interface as contract
type HasName interface {
	GetName() string
}

// function for interface implementations
func sayHello(value HasName) {
	fmt.Println("Hello from : ", value.GetName())
}

// struct implementing interface
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// struct implementing interface
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Felix"}
	sayHello(person)

	animal := Animal{Name: "Dog"}
	sayHello(animal)
}
