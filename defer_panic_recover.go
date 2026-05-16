package main

import "fmt"

func endApp() {
	fmt.Println("END APP")
	message := recover()
	fmt.Println("Error occured : ", message)
}

func runApp(error bool) {
	defer endApp()

	fmt.Println("RUN APP")

	if error {
		panic("Oops error")
	}
}

func main() {
	runApp(true)
	fmt.Println("CONTINUE APP")
}
