package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var felix Customer
	felix.Name = "Felix"
	felix.Address = "Jakarta"
	felix.Age = 24

	fmt.Println(felix)

	var a [6]int = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)

	var joko = Customer{
		Name:    "Joko",
		Age:     20,
		Address: "Jakarta",
	}

	budi := Customer{"Budi", "Jakarta", 20}

	fmt.Println(joko)
	fmt.Println(budi)
}
