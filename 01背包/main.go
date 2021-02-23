package main //

import (
	"fmt"
)

/*
0-1 背包问题
描述:
	给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。
	其中第 i 个物品的重量为 wt[i]，价值为 val[i]，现在让你用这个背包装物品，最多能装的价值是多少？
*/
func main() {
	var w = 4

	var n = 3
	var wt = []int{2, 1, 3}
	var val = []int{4, 2, 3}
	fmt.Println(knapsack(w, n, wt, val))
}

/*
dp[i][w] = v; 容量为w的背包, 装载前i个物品时, 最大价值为v

dp[i][w] =
	| 0 ; i == 0 or w == 0
	| max(dp[i-1][w-wt[i-1]] + val[i-1], dp[i-1][w]) ; wt[i-1] <= w
	| dp[i-1][w] ; wt[i] > w
*/
func knapsack(w int, n int, wt []int, val []int) int {

	// base case, golang中0值已初始化
	var dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < w+1; j++ {
			if wt[i-1] <= j {
				dp[i][j] = max(
					dp[i-1][j-wt[i-1]]+val[i-1],
					dp[i-1][j],
				)
				continue
			}

			dp[i][j] = dp[i-1][j]
		}
	}

	return dp[n][w]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
