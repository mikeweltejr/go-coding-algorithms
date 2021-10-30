package main

func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int)
	var retNums []int
	for i := 0; i < len(nums); i++ {
		if _, ok := intMap[target-nums[i]]; ok {
			retNums = append(retNums, intMap[target-nums[i]])
			retNums = append(retNums, i)
			break
		}

		intMap[nums[i]] = i
	}

	return retNums
}
