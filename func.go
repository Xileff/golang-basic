package main

import "fmt"

func main() {
	helloFelix := getHello("Felix")
	fmt.Println(helloFelix)
}

func getHello(name string) string {
	return "Hello " + name
}
