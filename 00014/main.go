package main // https://leetcode.com/problems/longest-common-prefix/

import "fmt"

func main() {
	var strs = []string{"flower", "a", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	var str = strs[0]
	var index = 0

OUT:
	for {
		for _, v := range strs {
			if index > len(v)-1 {
				break OUT
			}

			if str[index] != v[index] {
				break OUT
			}
		}

		index++
	}

	return str[0:index]
}
