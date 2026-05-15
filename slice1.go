package main

import "fmt"

func main() {
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	// Create slice with length 2 and capacity 2
	slice1 := days[5:] // Sabtu, Minggu
	fmt.Println(slice1)

	slice1[0] = "Sabtu Baru"
	slice1[1] = "Minggu Baru"

	// mutating data in slice == mutating data in original array
	fmt.Println(slice1)
	fmt.Println(days)

	// internally creates a new array, because original array can't change capacity
	slice2 := append(slice1, "Libur Baru")
	fmt.Println(slice2)

	fmt.Println(days)
}
