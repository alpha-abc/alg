package main // https://leetcode-cn.com/problems/find-majority-element-lcci/

import "fmt"

/*
核心思想:
不同元素相互抵消
*/
func majorityElement(nums []int) int {
	var major, cnt = 0, 0
	for _, n := range nums {
		if cnt == 0 {
			major = n
			cnt++
			continue
		}

		if major == n {
			cnt++
		} else {
			cnt--
		}
	}

	if cnt == 0 {
		return -1
	}

	var c = 0
	for _, n := range nums {
		if n == major {
			c++
		}
	}

	if c <= len(nums)/2 {
		return -1
	}

	return major
}

func main() {
	var nums = []int{1, 3, 3, 1, 3}
	fmt.Println(majorityElement(nums))
}
