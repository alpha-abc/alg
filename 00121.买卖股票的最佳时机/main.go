package main // https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

/*

dp[i][k][0 or 1] ; 0 表示持有, 1表示卖出
0 <= i <= N; 0 <= k <= K ; N为天数, K为最大交易次数

dp[0][k][0] = 0
dp[0][k][1] = 不存在
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k-1][0] - prices[i], dp[i-1][k][1])

*/
func maxProfit(prices []int) int {
	var n = len(prices)
	var k = 1

	// 不可能达到的值
	var no = -1
	for _, e := range prices {
		no -= e
	}

	var dp = make([][][]int, n+1)
	for i := 0; i < n+1; i++ {
		var kArr = make([][]int, k+1)
		dp[i] = kArr

		for j := 0; j < k+1; j++ {
			var arr = make([]int, 2)
			dp[i][j] = arr

			if i == 0 || j == 0 {
				dp[i][j][1] = no // 不存在
				continue
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < k+1; j++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i-1])
			dp[i][k][1] = max(dp[i-1][k-1][0]-prices[i-1], dp[i-1][k][1])
		}
	}

	return dp[n][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func print(dp [][][]int) {
	for i, e1 := range dp {
		for j, e2 := range e1 {
			fmt.Println("持股[否]", i, "天", j, "次交易", "最大收益", e2[0])
			fmt.Println("持股[是]", i, "天", j, "次交易", "最大收益", e2[1])
			fmt.Println()
		}
		fmt.Println("-----")
		fmt.Println()
	}
}
