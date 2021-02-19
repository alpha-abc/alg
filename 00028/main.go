package main // https://leetcode.com/problems/implement-strstr/

import "fmt"

func main() {
	fmt.Println(strStr("hello", "el"))
}

func strStr(haystack string, needle string) int {
	var s = []rune(haystack)
	var p = []rune(needle)
	var d = -1
	if len(p) == 0 {
		return 0
	}

	if len(p) > len(s) {
		return d
	}

	var next = nextArr(p)

	var i, j = 0, 0
	for i < len(s) {
		if j == -1 || s[i] == p[j] {
			i++
			j++

			if j == len(p) {
				return i - j
			}

		} else {
			j = next[j]
		}
	}

	return d
}

func nextArr(s []rune) []int {
	if len(s) <= 0 {
		return []int{}
	}

	var arr = []int{-1}
	var i, j = 0, -1

	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			arr = append(arr, j)
		} else {
			j = arr[j]
		}
	}

	return arr
}
