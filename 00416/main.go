package main // https://leetcode-cn.com/problems/partition-equal-subset-sum/ 416. 分割等和子集

import "fmt"

func main() {
	fmt.Println(canPartition([]int{2, 2, 1, 1}))
}

/*
tips:
	正整数不包括零

dp[i][j] = b; 前i个物品用容量j的容器装在, b为true时, 表示刚好装满, b为false, 表示不能刚好装满

dp[i][j] =
	| true ; j == 0, 背包没有空间就相当于装满
	| false; i == 0, 没有物品装时肯定装不满背包
	| dp[i-1][j] ; j < nums[i-1], 容量不足
	| dp[i-1][j] || dp[i-1][j-nums[i-1]]
*/
func canPartition(nums []int) bool {
	var sum = 0
	for _, n := range nums {
		sum = sum + n
	}
	// 需要划分为两个相等子集, sum必须为偶数
	if sum%2 != 0 {
		return false
	}

	// 转换为背包问题
	sum = sum / 2
	var n = len(nums)

	var dp = make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if j < nums[i-1] {
				// 容量不足
				dp[i][j] = dp[i-1][j]
				continue
			}

			dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
		}
	}

	return dp[n][sum]
}
