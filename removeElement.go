package main

func removeElement(nums []int, val int) int {
	count := 0
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			count++
			nums = remove(nums, i)
		} else {
			i++
		}
	}
	return len(nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
