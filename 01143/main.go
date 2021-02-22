package main // https://leetcode-cn.com/problems/longest-common-subsequence/ 1143. 最长公共子序列

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}

/*
假设有两个字符串 s1, s2
则, dp[i][j]表示 s1[0 ... i]和s2[0 ... j]最长的公共字符串长度

dp[i][j] =
	| 0; i == 0 or j == 0
	| dp[i-1][j-1]+1; i, j > 0 and s[i] == s[j]
	| max(dp[i-1][j], dp[i][j-1]); i, j > 0 and s[i] != s[j]
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	var n1, n2 = len(text1), len(text2)

	var dp = make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
	}

	// base case
	for i := 1; i < n1+1; i++ {
		dp[i][0] = 0
	}
	for j := 1; j < n2+1; j++ {
		dp[0][j] = 0
	}

	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}

			dp[i][j] = max(
				dp[i-1][j],
				dp[i][j-1],
			)
		}
	}

	return dp[n1][n2]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
