package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefixTwo(strs []string) string {
	retStr := ""

	if len(strs) == 1 {
		return strs[0]
	}

	initCharArr := []rune(strs[0])
	curIndex := 0

	for curIndex < len(initCharArr) {
		curChar := initCharArr[curIndex]

		for i := 1; i < len(strs); i++ {

			if curIndex >= len([]rune(strs[i])) || curChar != []rune(strs[i])[curIndex] {
				return retStr
			}
		}

		retStr += string(curChar)
		curIndex++
	}

	return retStr
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) == 0 {
			prefix := prefix[0 : len(prefix)-1]
			fmt.Println(prefix)
			if prefix == "" {
				return ""
			}
			fmt.Println(strings.Index(strs[i], prefix))
		}
	}

	return prefix
}
