package main

func findDisappearingNumbers(nums []int) []int {
	var retNums = make([]int, len(nums))

	for i := 1; i <= len(nums); i++ {
		retNums[i-1] = i
	}

	for i := 0; i < len(nums); i++ {
		curNum := nums[i]
		if curNum <= len(nums) {
			retNums[curNum-1] = -1
		}
	}

	tempSlice := []int{}

	for i := 0; i < len(retNums); i++ {
		if retNums[i] != -1 {
			tempSlice = append(tempSlice, retNums[i])
		}
	}

	missingNums := make([]int, len(tempSlice))

	copy(missingNums, tempSlice)

	return missingNums
}
