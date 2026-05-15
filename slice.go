package main

import "fmt"

// len : qty of data inside the slice
// cap : maximum capacity of data inside the slice

// cap is counted from start index to last index from initial array

// mutating slice = mutating array, because slice is actually pointing to the original array

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[4:6]

	fmt.Println(slice1)
	fmt.Println(len(slice1)) // 2
	fmt.Println(cap(slice1)) // 8

	slice1 = append(slice1, "AAA")
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // 3
	fmt.Println(cap(slice1)) // 8

	fmt.Println(months)
}
