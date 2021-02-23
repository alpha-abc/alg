package main // https://leetcode-cn.com/problems/jump-game/ 55. 跳跃游戏

func main() {}

/*
dp[i] = n ; 当前位置i能跳的最远距离

dp[i] = i + nums[i]
*/
func canJump(nums []int) bool {
	var farthest = 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest <= i {
			return false
		}
	}
	return farthest >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
