package main // https://leetcode-cn.com/problems/maximum-subarray/ 53. 最大子序和

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// 以 nums[i] 为结尾的「最大子数组和」为 dp[i]。
func maxSubArray(nums []int) int {
	var n = len(nums)

	// 压缩状态, 降低空间复杂度
	var prev = nums[0]
	var curr = 0
	var m = prev

	for i := 1; i < n; i++ {
		curr = max(prev+nums[i], nums[i])
		prev = curr

		m = max(m, curr)
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
