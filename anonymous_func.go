package main

import "fmt"

type BlacklistFunc func(string) bool

func registerUser(name string, blacklistFunc BlacklistFunc) {
	if blacklistFunc(name) {
		fmt.Println("Blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklistFunc := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Felix", blacklistFunc)
	registerUser("anjing", blacklistFunc)
}
