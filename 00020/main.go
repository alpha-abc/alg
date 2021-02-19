package main //https://leetcode.com/problems/valid-parentheses/

import (
	"fmt"
)

func isValid(s string) bool {
	var ab = []byte(s)
	var size = len(ab)

	if size%2 != 0 {
		return false
	}

	var arr = make([]byte, size)
	var i = 0
	for _, b := range ab {
		if b == '(' || b == '[' || b == '{' {
			arr[i] = b
			i++
			continue
		}

		i--
		if i < 0 {
			return false
		}

		if b == '}' && arr[i] == '{' {
			continue
		}

		if b == ')' && arr[i] == '(' {
			continue
		}

		if b == ']' && arr[i] == '[' {
			continue
		}

		return false
	}

	return i == 0
}

func main() {
	fmt.Println(isValid("(("))
}
