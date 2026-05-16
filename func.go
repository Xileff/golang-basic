package main

import "fmt"

func main() {
	helloFelix := getHello("Felix")
	fmt.Println(helloFelix)

	_, lastName := getFullName()
	fmt.Println(lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Full", "Name"
}
