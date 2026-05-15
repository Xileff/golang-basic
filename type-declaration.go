package main

import "fmt"

func main() {
	type NoKTP string

	var ktpSatu NoKTP = "11111111"
	var ktpDua string = "22222222"

	fmt.Println(ktpSatu)
	fmt.Println(NoKTP(ktpDua))
}
