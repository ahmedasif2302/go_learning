package main

import "fmt"

func goConditional() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("A is greater than B")
	} else if b > a {
		fmt.Println("B is greater than A")
	} else {
		fmt.Println("Nothing is greater than anything")
	}

	switch "20" {
	case "10":
		fmt.Println("Prints 10")
	case "20":
		fmt.Println("Prints 20")
	default:
		fmt.Println("Prints Default")
	}
}
