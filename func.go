package main

import "fmt"

func main() {
	helloFelix := getHello("Felix")
	fmt.Println(helloFelix)

	// multiple return value
	_, lastName := getFullName()
	fmt.Println(lastName)

	firstName, lastName := getNamedFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	// variadic function (varargs)
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6))

	numSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sumAll(numSlice...)) // like spred operator in TS

	// function as param
	sayHelloWithFilter("Anjing", spamFilter)
}

// normal func
func getHello(name string) string {
	return "Hello " + name
}

// multiple return value
func getFullName() (string, string) {
	return "Full", "Name"
}

func getNamedFullName() (firstName string, lastName string) {
	firstName = "Felix"
	lastName = "Savero"

	return firstName, lastName
}

// variadic function (varargs)
func sumAll(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// function as param
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
