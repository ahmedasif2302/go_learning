package main

import (
	"fmt"
)

func goLoops() {
	nums := []int{10, 20, 30, 40, 50}
	names := []string{"ahmed asif", "sharmi", "rayyan"}

	fmt.Println(nums)

	for i := 0; i < len(nums); i += 1 {
		if nums[i] == 30 {
			break
		}
		fmt.Println(nums[i] * 10)
	}

	for _, name := range names {
		for _, num := range nums {
			fmt.Printf("The name is %s and num is %d \n", name, num)
		}
	}

	c := 0
	for c < 10 {
		fmt.Println(c)
		c++
	}

}
