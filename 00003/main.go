package main // https://leetcode.com/problems/longest-substring-without-repeating-characters/

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba")) //abcbbefghg
}

func lengthOfLongestSubstring(s string) int {
	var chars = []rune(s)

	var m = make(map[rune]int)
	for i, char := range chars {
		var _, found = m[char]
		if !found {
			m[char] = i
		}
	}

	var longest = 0
	var l = 0
	for i, char := range chars {
		var v, found = m[char]

		if found && v < i {
			longest = max(longest, i-l)
			l = max(l, v+1)
			m[char] = i
		}
	}
	longest = max(longest, len(chars)-l)

	return longest
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
