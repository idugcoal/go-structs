package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	var peter person
	peter.firstName = "Peter"
	peter.lastName = "Piper"

	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("%+v", peter)
	fmt.Printf("%+v", alex)
}
