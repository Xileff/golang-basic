package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address Address = Address{"Jakarta Barat", "DKI Jakarta", "Wakanda"}

	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
