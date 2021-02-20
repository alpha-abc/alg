package main // https://leetcode-cn.com/problems/coin-change/

import "fmt"

func main() {
	var coins = []int{2, 5, 10, 1}
	var amount = 27
	fmt.Println(coinChange(coins, amount))
}

func coinChange1(coins []int, amount int) int {
	var memo = map[int]int{}
	var n = dp(amount, coins, memo)
	return n
}

/*
dp(amount) = n

dp(amount) =
	|-> -1; amount < 0
	|-> 0; amount = 0
	|-> min(dp(amount - coin) + 1, dp(amount)); amount > 0
*/
func dp(amount int, coins []int, memo map[int]int) int {
	if resp, ok := memo[amount]; ok {
		return resp
	}

	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	var res = -2 // -2当正无穷处理
	for _, c := range coins {
		var n = dp(amount-c, coins, memo)

		if n == -1 {
			continue
		}

		if res == -2 {
			res = n + 1
			continue
		}

		if res > n+1 {
			res = n + 1
		}
	}

	if res == -2 {
		res = -1
	}

	memo[amount] = res
	return res
}

func coinChange(coins []int, amount int) int {
	// 初始化
	// 当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出
	var dp = make([]int, amount+1)
	for i := 1; /*dp[0]=0*/ i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < len(dp); i++ {
		for _, c := range coins {
			if i-c < 0 {
				continue
			}

			var n = 1 + dp[i-c]
			if dp[i] > n {
				dp[i] = n
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
