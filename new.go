package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var address1 *Address = new(Address)
	var address2 *Address = address1

	address2.City = "Jakarta"

	fmt.Println(address1) // changed
	fmt.Println(address2) // changed
}
