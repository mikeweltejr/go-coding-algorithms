package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{}

	// { 1, 2, 3, 4, 5, 6, 7 }

	result := medianOfSortedArrays(nums1, nums2)

	fmt.Println(result)
}
