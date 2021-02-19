package main // https://leetcode.com/problems/string-to-integer-atoi/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("-98"))
}

func myAtoi(str string) int {
	var chars = []rune(str)
	var resp = format(chars)

	if len(resp) <= 0 {
		return 0
	}

	if len(resp) == 1 && resp[0] == '-' {
		return 0
	}

	var s int64 = 1
	if resp[0] == '-' {
		s = -1
		resp = resp[1:]
	}

	if resp[0] == '+' {
		resp = resp[1:]
	}

	var n int64

	for i := len(resp) - 1; i >= 0; i-- {
		n = n*10 + int64(resp[len(resp)-1-i]-'0')
		if s == -1 && n > math.MaxInt32 {
			return math.MinInt32
		}
		if s == 1 && n > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	n = s * n
	return int(n)
}

func format(chars []rune) []rune {
	var resp []rune

	var l = 0
	var space = 0
	var zero = 0
	for i, char := range chars {
		// 0 - 48 , ... , 9 - 57
		switch {
		case char == ' ':
			if l == i && zero == 0 {
				l++
				space++
			} else {
				return resp
			}
		case char == '-' || char == '+':
			if l == i && zero == 0 {
				resp = append(resp, char)
			} else {
				return resp
			}
		case char == '0':
			if l == i && space == 0 {
				l++
				zero++
			} else {
				resp = append(resp, char)
			}
		case char > 48 && char <= 57:
			resp = append(resp, char)
		default:
			return resp
		}
	}

	return resp
}
