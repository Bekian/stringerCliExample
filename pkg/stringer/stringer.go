package stringer

import (
	"fmt"
	"strconv"
)

func Reverse(input string) (output string) {
	for _, c := range input {
		output += string(c)
	}
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
