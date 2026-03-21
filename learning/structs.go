package main

import "fmt"

type Address struct {
	street,
	zip,
	state,
	country string
}

type Person struct {
	name    string
	age     int
	address Address
}

func goStructs() {
	var person Person

	person.age = 10
	person.name = "Ahmed Asif"

	person.address.country = "India"
	person.address.zip = "600006"

	fmt.Println(person.name)
}
