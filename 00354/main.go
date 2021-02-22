package main // https://leetcode-cn.com/problems/russian-doll-envelopes/ 354. 俄罗斯套娃信封问题

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(
		maxEnvelopes([][]int{
			{5, 4}, {6, 4}, {6, 7}, {2, 3}, // 3
			// {1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}, // 3
			// {1, 2}, {2, 3}, {3, 4}, {3, 5}, {4, 5}, {5, 5}, {5, 6}, {6, 7}, {7, 8}, // 7
			// {5, 4}, {6, 4}, {6, 7}, {2, 3},
		}),
	)
}

/*
[[5,4],[6,4],[6,7],[2,3]]

假设有一序列s
dp[i] 表示以s[i]结尾(前面的数字比s[i]的值小)的最长上升子序列的长度
O(N^2)

dp[i] =
	| 1 ; i == 0
	| max(dp[j]) + 1; s[i] > s[j] && i > j >= 0
*/
func maxEnvelopes(envelopes [][]int) int {
	var n = len(envelopes)
	if n == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		// 按照w升序排列, 若w相等则按照h降序排列
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	var dp = make([]int, n)
	dp[0] = 1

	var result = 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
