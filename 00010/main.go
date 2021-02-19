package main // https://leetcode.com/problems/regular-expression-matching/

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}

func isMatch(s string, p string) bool {
	if len(s) == 0 {
		return len(p) > 0
	}

	if len(p) == 0 {
		return len(s) > 0
	}

	if s[0] == p[0] {
		return true
	}

	if p[0] == '.' {
		return isMatch(s[1:], p[1:])
	}

	if p[0] == '*' {

	}

	return false
}
