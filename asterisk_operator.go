package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var address1 Address = Address{"Jakarta Barat", "Jakarta", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	// Updating field of a pointer variable = updating the original variable
	fmt.Println("Updating field of pointer variable")
	address2.City = "Jakarta Timur"

	fmt.Println(address1)
	fmt.Println(address2)

	// by using *, reassigning new value to a pointer variable means also reassingning to the real variable
	fmt.Println("Reassigning to pointer variable using * operator")
	*address2 = Address{"Singkawang", "Kalimantan Barat", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
