package main // https://leetcode.com/problems/remove-element/

import "fmt"

func main() {
	var arr = []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var i = 0
	for _, n := range nums {
		if n != val {
			nums[i] = n
			i++
		}
	}

	return i
}
