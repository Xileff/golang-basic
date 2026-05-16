package main

import "fmt"

func main() {
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Felix"
	newSlice[1] = "Savero"

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice = append(newSlice, "Xilef")
	fmt.Println(newSlice)

	// copy slice
	fromSlice := newSlice[:]
	toSlice := make([]string, len(newSlice), cap(newSlice))

	copy(fromSlice, toSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
