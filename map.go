package main

import "fmt"

func main() {
	personMap := map[string]string{
		"ID-001": "Felix",
		"ID-002": "Albert",
	}

	fmt.Println(personMap)
	fmt.Println(personMap["ID-001"])

	animalMap := make(map[string]string)
	animalMap["ID-001"] = "Dog"

	fmt.Println(animalMap)
}
