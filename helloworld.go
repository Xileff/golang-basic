package main

import "fmt"

func main() {
	var name string

	name = "test"

	age := 10

	fmt.Print("Hello world" + name)
	fmt.Print(age)

	var (
		firstName = "First Name"
		lastName  = "Last Name"
	)

	fmt.Print(firstName + " " + lastName)
}
