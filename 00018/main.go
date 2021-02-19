package main //https://leetcode.com/problems/4sum/

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	var n = len(nums)

	if n < 4 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}

			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}

			var l = j + 1
			var r = n - 1

			for l < r {
				var sum = nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}

					r--
					for l < r && nums[r+1] == nums[r] {
						r--
					}
				}
			}
		}
	}

	return res
}

func main() {
	var a = []int{-4, -1, -1, -1, 0, 1, 2}
	a = []int{0, 0, 0, 1, 1}
	//a = []int{1, 0, -1, 0, -2, 2}
	var res = fourSum(a, 2)
	fmt.Println(res)
}
