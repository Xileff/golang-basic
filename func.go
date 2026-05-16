package main

import "fmt"

func main() {
	helloFelix := getHello("Felix")
	fmt.Println(helloFelix)

	_, lastName := getFullName()
	fmt.Println(lastName)

	firstName, lastName := getNamedFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Full", "Name"
}

func getNamedFullName() (firstName string, lastName string) {
	firstName = "Felix"
	lastName = "Savero"

	return firstName, lastName
}
