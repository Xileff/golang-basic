package main

import (
	"fmt"
)

func main() {
	const int32Value = 32768
	const int64Value = int64(int32Value)

	fmt.Println(int32Value)
	fmt.Println(int64Value)

	var name = "Felix"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
