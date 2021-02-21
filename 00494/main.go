package main // https://leetcode-cn.com/problems/target-sum/
// 494. 目标和
import "fmt"

func main() {
	var nums = []int{2, 4, 1, 5, 15}
	var S = 1
	fmt.Println(findTargetSumWays(nums, S))
}

/*
将nums分为A, B两个子集

sum(A) - sum(B) = S
sum(A) = sum(B) + S
sum(A) + sum(A)
	= S + sum(A) + sum(B)
	= S + sum(nums)

最终问题变为背包问题:
sum(A) = (target + sum(nums)) / 2

dp[N][sum] = r
N个物品, 装满容量为sum的背包, 有r种方法可以装满

dp[N][sum]
	| = 1; 容量sum为0时, 只有不装这一种情况
	| = dp[N-1][sum] + dp[N-1][sum-nums[N-1]]
*/

func findTargetSumWays(nums []int, S int) int {
	var sum = 0
	for _, n := range nums {
		sum = sum + n
	}

	if sum < S || (sum+S)%2 == 1 {
		// 不存在合法子集划分
		return 0
	}

	return subsets(nums, (sum+S)/2)
}

// 计算 nums 中有几个子集的和为 sum
func subsets(nums []int, sum int) int {
	var n = len(nums)
	var dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, sum+1)
		// 只有不装这一种情况
		dp[i][0] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j < sum+1; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
				continue
			}

			dp[i][j] = dp[i-1][j]
		}
	}

	return dp[n][sum]
}

func dpPrint(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}
}
