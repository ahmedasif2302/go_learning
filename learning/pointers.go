package main

import "fmt"

func pointerExample() {
	var i = 20
	p := &i

	fmt.Println(i)
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)
	fmt.Println(*p)

	var a = new(int)

	fmt.Println(*a)
}
