package main

import "fmt"

// nil can only be used in : interface, function, map, slice, pointer, channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Empty data")
	} else {
		fmt.Println(data)
	}
}
