package main // https://leetcode.com/problems/3sum-closest/

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{0, 0, 0}
	var n = threeSumClosest(arr, 1)
	fmt.Println(n)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var target1 = target
	var diffNum = abs(nums[0] + nums[1] + nums[len(nums)-1] - target)

	for i := 0; i < len(nums)-2; i++ {
		var curr = nums[i]

		if i > 0 && curr == nums[i-1] {
			continue
		}

		var l, r = i + 1, len(nums) - 1
		for l < r {
			var sum = curr + nums[l] + nums[r]
			var tmpDiffNum = abs(target - sum)

			if diffNum >= tmpDiffNum {
				diffNum = tmpDiffNum
				target1 = sum
			}

			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				return sum
			}
		}

	}
	return target1
}
