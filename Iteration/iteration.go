package Iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	//TWO WAYS OF DOING IT:
	//basic approach with a for loop with concatenation
	/*
		result := ""
		for i := 0; i < repeatCount; i++ {
			result += character
		}
		return result*/

	//Using a string builder with WriteString() method
	var repeatedStr strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeatedStr.WriteString(character)
	}
	return repeatedStr.String()
}
