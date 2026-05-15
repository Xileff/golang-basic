package main

import "fmt"

// array in golang have immutable qty
// -> can't delete array elements, can only set to default values (empty string, 0)

func main() {
	var names [3]string

	names[0] = "felix0"
	names[1] = "felix1"
	names[2] = "felix2"

	fmt.Println(names)

	var values1 = [3]int{
		0,
		1,
		2,
	}

	fmt.Println(values1)

	var values2 = [...]string{
		"Felix",
		"Steven",
		"Vincent",
		"Renard",
		"Ferry",
	}

	fmt.Println(values2)
}
