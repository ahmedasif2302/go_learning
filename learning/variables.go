package main

import "fmt"

const NAME float32 = 10

type User struct {
	Name string
	Age  int
}

const (
	SUNDAY = iota
	MONDAY
	TUESDAY
)

func variableExample() {
	var msg string
	var x, y float64
	fmt.Println(x, y, NAME)
	msg = "Ismail"
	fmt.Println(msg)

	fmt.Printf("%T \n", NAME)
	fmt.Printf("%v%% \n", NAME)

	u := User{Name: "Ahmed Asif", Age: 20}

	fmt.Printf("%#v \n", u)
	x = 1.7e+308
}
