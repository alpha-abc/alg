package main // https://leetcode-cn.com/problems/edit-distance/ 72. 编辑距离

import (
	"fmt"
)

func main() {
	fmt.Println(minDistance1("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

/*
// word1, word2 均为小写字母


*/

// 暴力解法
func minDistance1(word1 string, word2 string) int {
	return dp(len(word1)-1, len(word2)-1, word1, word2)
}

func dp(i int, j int, word1, word2 string) int {
	// base case
	if i == -1 {
		return j + 1
	}

	if j == -1 {
		return i + 1
	}

	if word1[i] == word2[j] {
		return dp(i-1, j-1, word1, word2)
	}

	return min(
		dp(i, j-1, word1, word2)+1,   // 插入
		dp(i-1, j, word1, word2)+1,   // 删除
		dp(i-1, j-1, word1, word2)+1, // 替换
	)
}

func min(a, b, c int) int {
	var min = a
	if min > b {
		min = b
	}

	if min > c {
		min = c
	}

	return min
}

// 动态规划, DP TABLE解法
func minDistance(word1 string, word2 string) int {
	var n1, n2 = len(word1), len(word2)

	// dp[i][j] 代表 word1[0 ... i] 到 word2[0 ... j]的最小编辑距离
	var dp = make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
	}

	// base case
	dp[0][0] = 0
	for i := 1; i < n1+1; i++ {
		dp[i][0] = i
	}
	for j := 1; j < n2+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			dp[i][j] = min(
				// 插入
				dp[i][j-1]+1,
				// 删除
				dp[i-1][j]+1,
				// 替换
				dp[i-1][j-1]+1,
			)
		}
	}

	return dp[n1][n2]
}
