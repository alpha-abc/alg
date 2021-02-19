package main // https://leetcode-cn.com/problems/search-insert-position/

import "fmt"

// 35. 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 你可以假设数组中无重复元素。
//
// 示例 1:
// 输入: [1,3,5,6], 5
// 输出: 2

// 示例 2:
// 输入: [1,3,5,6], 2
// 输出: 1

// 示例 3:
// 输入: [1,3,5,6], 7
// 输出: 4

// 示例 4:
// 输入: [1,3,5,6], 0
// 输出: 0

func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	var size = len(nums)
	var l, r = 0, size - 1

	for l <= r {
		var m = l + (r-l)/2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if l < size && nums[l] < target {
		return l + 1
	}

	return l
}

func main() {
	var nums = []int{1, 3, 5, 6}
	var target = 2
	var res = searchInsert(nums, target)
	fmt.Println(res)
}
