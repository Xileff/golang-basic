package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()

	switch result.(type) {
	case string:
		fmt.Println("String", result)
	case int:
		fmt.Println("Int", result)
	default:
		fmt.Println("Unknown", result)
	}
}
