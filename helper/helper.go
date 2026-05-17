package helper

import "fmt"

// private
var version = "1.0.0"

func sayGoodbye() {
	fmt.Println("Goodbye")
}

// public
func SayHello(name string) {
	fmt.Println("Hello ", name)
}
