package main

func longestCommonPrefix(strs []string) string {
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
