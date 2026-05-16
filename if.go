package main

import "fmt"

func main() {
	number := 10

	if number >= 10 {
		fmt.Println("greater than 10")
	} else if number >= 5 {
		fmt.Println("greater than 5")
	} else {
		fmt.Println("less than 5")
	}

	// short statement
	name := "Felix"

	if length := len(name); length > 5 {
		fmt.Println("name too long")
	} else {
		fmt.Println("name ok")
	}
}
