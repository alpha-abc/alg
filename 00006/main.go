package main // https://leetcode.com/problems/zigzag-conversion/

import "fmt"

/*
P   A   H   N
A P L S I I G
Y   I   R
*/
func main() {
	fmt.Println(convert("A", 1))
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	var chars = []rune(s)
	var sLen = len(chars)

	var n = 2*numRows - 2
	var newChars = []rune{}

	for i := 0; i < numRows; i++ {
		var step = 0
		for {
			var index = i + step*n
			var index1 = index + n - i*2

			if index > sLen-1 {
				break
			}

			if i > 0 && i < numRows-1 && index1 < sLen {
				newChars = append(newChars, chars[index])
				newChars = append(newChars, chars[index1])
			} else {
				newChars = append(newChars, chars[index])
			}
			step++
		}
	}
	return string(newChars)
}
