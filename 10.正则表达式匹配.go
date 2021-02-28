package main // https://leetcode-cn.com/problems/regular-expression-matching/

import "fmt"

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start

/*
e.g.
1
	b	a	a
idx	i-3	i-2	i-1
dp	i-2	i-1	i

	b	a	*
idx	j-3	j-2	j-1
dp	j-2	j-1	j

s只能包含a_z的字符串
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素

dp[i][j] = b; s前i “[0,i)” 个字符与p前j “[0,j)” 个字符是否匹配, b-true匹配, b-false不匹配

dp[i][j] =
	| dp[i-1][j-1] 	// s[i-1] == p[j-1]
	| false 		// s[i-1] != p[j-1] and p[j-1] == “普通字符”
	| dp[i-1][j-1]	// s[i-1] != p[j-1] and p[j-1] == '.'
	|				// s[i-1] != p[j-1] and p[j-1] == '*'
		| dp[i][j-2]	// s[i-1] != p[j-2]
		| dp[i-1][j]	// s[i-1] == p[j-2]


*/
func isMatch(s string, p string) bool {

	var ns, np = len(s), len(p)
	var dp = make([][]bool, ns+1)
	for i := 0; i < ns+1; i++ {
		dp[i] = make([]bool, np+1)
	}

	// base case

	// 两个空字符串是匹配的
	dp[0][0] = true

	// s, p中只有一个为空, 数组初始化值默认false, 此处不初始化
	// do nothing

	// s, p都不为空
	for i := 0; i < ns+1; i++ {
		for j := 1; j < np+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
				continue
			}

			if i > 0 && (p[j-1] == '.' || s[i-1] == p[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[ns][np]
}

// @lc code=end

func printDP(dp [][]bool) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Print(fmt.Sprintf("%8v", dp[i][j]), " ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println(isMatch("aa", "a"))                   // false
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // false
	fmt.Println(isMatch("aa", "a*"))                  // true
	fmt.Println(isMatch("ab", ".*"))                  // true
	fmt.Println(isMatch("aab", "c*a*b*"))             // true
}
