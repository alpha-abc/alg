package main // https://leetcode.com/problems/letter-combinations-of-a-phone-number/

import (
	"fmt"
	"strings"
)

func main() {
	var r = letterCombinations("23743874636826")
	fmt.Println(r)

}

func letterCombinations(digits string) []string {
	var m = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return nil
	} else if len(digits) == 1 {
		return m[digits]
	}

	var n = len(digits)
	var prev = letterCombinations(digits[0 : n-1])
	var last = m[digits[n-1:n]]

	var res []string
	for _, v := range prev {
		for _, v1 := range last {
			var builder = new(strings.Builder)
			builder.WriteString(v)
			builder.WriteString(v1)
			res = append(res, builder.String())
		}
	}
	return res
}
