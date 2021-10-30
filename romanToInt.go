package main

import "fmt"

func romanToInt(s string) int {
	// "MCMXCIV"
	chars := []rune(s)
	intVal := 0
	for i := 0; i < len(chars); i++ {
		strChar := string(chars[i])
		curInt := getCharVal(strChar)

		if i > 0 {
			prevInt := getCharVal(string(chars[i-1]))

			if curInt > prevInt {
				intVal += (curInt - prevInt - prevInt)
			} else {
				intVal += curInt
			}
		} else {
			intVal = curInt
		}

		fmt.Println(intVal)
	}

	return intVal
}

func getCharVal(c string) int {
	intVal := 0

	if c == "I" {
		intVal = 1
	} else if c == "V" {
		intVal = 5
	} else if c == "X" {
		intVal = 10
	} else if c == "L" {
		intVal = 50
	} else if c == "C" {
		intVal = 100
	} else if c == "D" {
		intVal = 500
	} else if c == "M" {
		intVal = 1000
	}
	return intVal
}
