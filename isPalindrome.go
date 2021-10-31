package main

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	chars := []rune(s)

	for i, j := 0, len(chars)-1; i < len(chars); i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}

	return true
}
