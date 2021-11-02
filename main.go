package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2}

	// { 1, 2, 3, 4, 5, 6, 7 }

	result := removeDuplicates(nums1)

	fmt.Println(result)
}
