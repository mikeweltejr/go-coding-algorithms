package main

import "fmt"

func medianOfSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := append(nums1, nums2...)
	mergedArr = mergeSort(mergedArr, 0, len(mergedArr)-1)
	sum := float64(0)

	mid := (len(mergedArr)) / 2

	fmt.Println(mid)

	if mid == 0 {
		sum += float64(mergedArr[0])
		return sum
	} else if len(mergedArr) == 2 {
		sum += (float64(mergedArr[0]) + float64(mergedArr[1])) / 2
		return sum
	} else if len(mergedArr)%2 == 0 {
		sum += float64(mergedArr[mid-1])
		sum += float64(mergedArr[mid])
		return sum / 2
	} else {
		sum += float64(mergedArr[mid])
		return sum
	}
}

func mergeSort(nums []int, left int, right int) []int {
	if left < right {
		mid := (left + right) / 2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, right, mid)
	}

	return nums
}

func merge(nums []int, left int, right int, mid int) {
	b := make([]int, len(nums)+1)
	i := left
	j := mid + 1
	k := left

	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			b[k] = nums[i]
			i++
		} else {
			b[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		b[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		b[k] = nums[j]
		j++
		k++
	}

	for x := left; x < right+1; x++ {
		nums[x] = b[x]
	}
}
