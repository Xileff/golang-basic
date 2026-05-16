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

func sumAll(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
