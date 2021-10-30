package main

import "fmt"

func main() {
	intArr := []int{3, 2, 4}
	intVals := twoSum(intArr, 6)

	fmt.Println(intVals[0], intVals[1])
}
