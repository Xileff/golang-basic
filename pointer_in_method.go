package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	var john Man = Man{"John"}

	john.Married()

	fmt.Println(john)
}
