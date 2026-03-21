package main

import "fmt"

func arraysExample() {

	var numbers = [5]int{1, 2, 3, 4, 5}
	var arr1 = [...]int{}

	fmt.Println(arr1)
	var nums = [...]int{1, 2, 3, 4, 5}
	numbers[1] = 20

	var cars = [...]string{"abc", "abc"}
	fmt.Println(cars)
	fmt.Println(len(nums))

	arr2 := [...]int{2: 20, 4: 40}
	fmt.Println(arr2)

}
