package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var address1 Address = Address{"Jakarta Barat", "Jakarta", "Indonesia"}
	var address2 Address = address1 // pass by value, not reference -> all data from address 1 copied to address 2

	address2.City = "Jakarta Timur"

	fmt.Println(address1) // no change
	fmt.Println(address2) // changed

	// pass by reference using pointer
	var address3 *Address = &address1
	address3.City = "Bekasi"
	address3.Province = "Jawa Barat"

	fmt.Println(address1) // changed
	fmt.Println(address3) // changed
}
