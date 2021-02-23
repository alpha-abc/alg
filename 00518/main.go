package main // https://leetcode-cn.com/problems/coin-change-2/ 518. 零钱兑换 II

import "fmt"

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}

/*
dp[i][j] = v, 若只使用coins中的前i个硬币的面值, 若想凑出金额j, 有v种凑法

dp[i][j] =
	| 0 ; i == 0
	| 1 ; j == 0
	| dp[i-1][j] ; j < coins[i-1] | 不使用coins[i]凑出j
	| dp[i][j-conins[i-1]] + dp[i-1][j] ; j >= coins[i-1] | 使用coins[i]凑出j
*/
func change(amount int, coins []int) int {
	var n = len(coins)

	var dp = make([][]int, n+1)
	// base case
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < amount+1; j++ {
			if j < coins[i-1] {
				// 容量不足
				dp[i][j] = dp[i-1][j]
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
		}
	}

	return dp[n][amount]
}
