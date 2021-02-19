package main // https://leetcode.com/problems/generate-parentheses/

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var res = make([]string, 0)
	var arr = make([]rune, n*2)
	do(&res, arr, 0, n, n)
	return res
}

func do(res *[]string, arr []rune, i int, l, r int) {
	if l == 0 && r == 0 {
		*res = append(*res, string(arr))
		return
	}

	if l > 0 && l <= r {
		arr[i] = '('
		do(res, arr, i+1, l-1, r)
	}

	if r > 0 {
		arr[i] = ')'
		do(res, arr, i+1, l, r-1)
	}
}
