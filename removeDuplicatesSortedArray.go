package main

func removeDuplicates(nums []int) int {
	existingNums := make(map[int]int)
	curIndex := 0

	for i := 0; i < len(nums); i++ {
		if _, ok := existingNums[nums[i]]; !ok {
			existingNums[nums[i]] = nums[i]
			nums[curIndex] = nums[i]
			curIndex++
		}
	}

	nums = nums[0:curIndex]

	return len(nums)
}
