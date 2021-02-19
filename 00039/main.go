package main // https://leetcode-cn.com/problems/combination-sum/

import (
	"fmt"
	"sort"
)

/*
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。

说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。

示例 1:
输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]

示例 2:
输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var dict = make(map[int][][]int)

	for _, i := range candidates {
		dict[i] = [][]int{}
	}

	// for _, _ := range candidates {

	// }

	fmt.Println(dict)

	return dict[target]
}

func main() {
	var c = []int{3, 6, 7, 2}
	var target = 7

	fmt.Println(combinationSum(c, target))
}
