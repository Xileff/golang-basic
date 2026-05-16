package main

import "fmt"

func main() {
	name := "Felixx"

	switch name {
	case "Felix":
		fmt.Println("Name is felix")
	case "Albert":
		fmt.Println("name is albert")
	default:
		fmt.Println("unknown")
	}
}
