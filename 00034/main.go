package main // https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

import (
	"fmt"
)

/*
34. 在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func searchRange(nums []int, target int) []int {
	var res = []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return res
	}

	var size = len(nums)
	// find first
	var l, r = 0, size - 1
	for l <= r {
		var m = l + (r-l)>>1
		if nums[m] >= target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if l < size && nums[l] == target {
		res[0] = l
	}

	// find last
	l, r = 0, size-1
	for l <= r {
		var m = l + (r-l)>>1
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if r >= 0 && nums[r] == target {
		res[1] = r
	}

	return res
}

func main() {
	var nums = []int{1}
	var target = 1
	var res = searchRange(nums, target)
	fmt.Println(res)
}
