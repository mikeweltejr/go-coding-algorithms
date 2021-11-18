package main

import "fmt"

func main() {
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}

	retNums := findDisappearingNumbers(nums1)

	for _, num := range retNums {
		fmt.Println(num)
	}
}
