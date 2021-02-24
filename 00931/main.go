package main // https://leetcode-cn.com/problems/minimum-falling-path-sum/ 931. 下降路径最小和

import (
	"fmt"
)

func main() {
	fmt.Println(
		minFallingPathSum([][]int{
			{},
		}),
	)
}

/*
e.g.

2 1 3
6 5 4
7 8 9

dp[i][j] = n ; 到达坐标i,j时, 最短路径为n

dp[i][j] =
	| matrix[i][j] ; i == 0
	| matrix[i][j] + min(matrix[i-1][j], matrix[i-1][j+1]); i > 0 and j == 0
	| matrix[i][j] + min(matrix[i-1][j], matrix[i-1][j-1]); i > 0 and j == len(matrix[i]) - 1
	| matrix[i][j] + min(matrix[i-1][j], matrix[i-1][j-1], matrix[i-1][j+1]); i > 0 and j > 0 and j < len(matrix[i]) - 1
*/
func minFallingPathSum(matrix [][]int) int {
	var m = len(matrix)
	if m == 0 {
		return 0
	}
	var n = len(matrix[0])
	if n == 0 {
		return 0
	}

	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// base case
	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
				continue
			}

			if j == n-1 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
				continue
			}
			// else if (j > 0 and j < n - 1)
			dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-1]), dp[i-1][j+1]) + matrix[i][j]
		}
	}

	var resp = dp[m-1][0]
	for j := 1; j < n; j++ {
		resp = min(resp, dp[m-1][j])
	}
	return resp
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func dpPrint(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}
}
