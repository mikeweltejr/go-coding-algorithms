package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)

	fmt.Println(result)
}
