package main

import "strconv"

func decodeString(s string) string {
	chars := []rune(s)
	retStr := ""

	for i := 0; i < len(chars); i++ {
		curChar := string(chars[i])
		if !isNumeric(curChar) {
			retStr += curChar
		} else if isNumeric(curChar) {

		}
	}
	return retStr

}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
