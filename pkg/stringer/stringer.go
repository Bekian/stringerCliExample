package stringer

import (
	"fmt"
	"strconv"
)

func Reverse(input string) (output string) {
	// Convert the string to a rune slice to work with individual characters
	inputRunes := []rune(input)
	length := len(inputRunes)
	// Initialize an empty rune slice to store the reversed characters
	reversed := make([]rune, length)
	// Reverse the characters
	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[i] = inputRunes[j]
	}
	// Convert the reversed rune slice back to a string
	output = string(reversed)
	return

}

func Inspect(input string, digits bool) (count int, kind string) {
	if !digits {
		kind = "char"
		count = len(input)
		return
	}
	kind = "digit"
	count = inspectNumbers(input)
	return
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("string to number conversion error: ", err)
		}
		count++
	}
	return
}
