package main // https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

import "fmt"

func main() {
	var nums = []int{1, 3}
	fmt.Println(search(nums, 3))
}

func search1(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}

	return -1
}

func search(nums []int, target int) int {
	var n = len(nums)
	if n == 0 {
		return -1
	}

	var l, r = 0, n - 1

	for l <= r {
		var m = (r-l)>>1 + l
		if nums[m] == target {
			return m
		}

		if nums[m] > nums[l] { // 左侧升序.
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[r] { // 右侧升序.
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[l] == nums[m] {
				l++
			}

			if nums[r] == nums[m] {
				r--
			}
		}
	}

	return -1

}
