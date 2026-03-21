package main

import "fmt"

func goFunctionsWithMultipleReturn(x int, y int) (sum int, sub int, mul int, div int) {
	sum = x + y
	sub = x - y
	mul = x * y
	div = x / y
	return
}

func testRecursion(x int) int {
	if x == 10 {
		return 0
	}
	fmt.Println(x)
	return testRecursion(x + 1)
}

func goMainFunction() {
	sum, sub, mul, div := goFunctionsWithMultipleReturn(20, 10)
	fmt.Println(sum, sub, mul, div)

	testRecursion(1)
}
