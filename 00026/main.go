package main // https://leetcode.com/problems/remove-duplicates-from-sorted-array/

import "fmt"

/**
Given a sorted array nums,
remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
	Given nums = [1,1,2],
	Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
	It doesn't matter what you leave beyond the returned length.

Example 2:
	Given nums = [0,0,1,1,1,2,2,3,3,4],
	Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.
	It doesn't matter what values are set beyond the returned length.


*/

func main() {
	var arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	var l = removeDuplicates(arr)
	fmt.Println(l, arr)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	var i = 0
	for _, v := range nums[1:] {
		if v == nums[i] {
			continue
		}
		i++
		nums[i] = v
	}

	return i + 1
}
