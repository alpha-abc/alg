package main // https://leetcode.com/problems/longest-palindromic-substring/

import "fmt"

/*
Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.


Example 2:

Input: "cbbd"
Output: "bb"
*/
func main() {
	var r = longestPalindrome("babadada")
	fmt.Println(r)
}

func longestPalindrome(s string) string {
	var chars = []rune(s)
	var n = len(chars)
	if n <= 1 {
		return s
	}

	var left, right = 0, 0
	for i := 0; i < n; i++ {
		var l, r = i - 1, i + 1
		for l >= 0 && r <= n-1 && chars[l] == chars[r] {
			if r-l >= right-left {
				left, right = l, r
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r <= n-1 && chars[l] == chars[r] {
			if r-l >= right-left {
				left, right = l, r
			}
			l--
			r++
		}
	}

	return string(chars[left : right+1])
}
