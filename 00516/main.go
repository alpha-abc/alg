package main // https://leetcode-cn.com/problems/longest-palindromic-subsequence/
// 516. 最长回文子序列
import "fmt"

func main() {
	longestPalindromeSubseq("bbbab")
}

func dpPrint(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}
}

/*
// 例子
s = "bbbab"

dp[i][j] = n | i <= j; n表示下表i到j之间的回文子序列长度

// 初始化值
if j - i == 0
	if s[i] == s[j]
		dp[i][j] = 2
	if s[i] != s[j]
		dp[i][j] = 1

if i == j
	dp[i][j] = 1


if s[i] == s[j]
	dp[i][j] = dp[i+1][j-1] + 2

if s[i] != s[j]
	dp[i][j] = max(dp[i+1][j], dp[i][j-1])
*/

func longestPalindromeSubseq(s string) int {
	// s 只包含小写英文字母
	var ln = len(s)
	if ln <= 0 {
		return 0
	}

	var dp = make([][]int, ln)
	// 初始化 (长度1)
	for i := 0; i < ln; i++ {
		dp[i] = make([]int, ln)
		// 单个字符也是回文, 长度 1
		dp[i][i] = 1
	}

	// 初始化 (长度2)
	for i := 0; i < ln-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
			continue
		}
		dp[i][i+1] = 1
	}

	// b b b a b
	// 遍历, 从长度3开始, 直到字符串总长度为止
	for l := 2; l < ln; l++ {
		for j := 0; j < ln-l; j++ {
			if s[j] == s[j+l] {
				dp[j][j+l] = dp[j+1][j+l-1] + 2
				continue
			}

			var n1 = dp[j+1][j+l]
			var n2 = dp[j][j+l-1]

			if n1 > n2 {
				dp[j][j+l] = n1
				continue
			}
			dp[j][j+l] = n2
		}
	}

	return dp[0][ln-1]
}
