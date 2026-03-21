package main

import "fmt"

func goSlicesExamples() {
	nums := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(cap(nums))

	// creating slice using make
	a := make([]int, 10)

	a = append(a, 10, 20, 40)

	// a = append(a, nums...)

	fmt.Println(len(a))
}
