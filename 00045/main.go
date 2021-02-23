package main // https://leetcode-cn.com/problems/jump-game-ii/ 45. 跳跃游戏 II

import "fmt"

func main() {
	fmt.Println(
		jump([]int{2, 3, 1, 1, 4}),
	)
}

/*
dp[i] = n ; 到达i所用的最小步数为n

dp[i] =
	| 0 ; i == 0
	| min(dp[j] + 1, dp[i]) ; j < i and nums[j] + j >= i
*/
func jump(nums []int) int {
	var n = len(nums)

	var dp = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = 0
			continue
		}
		// 设置一个超出的步数
		dp[i] = n
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
