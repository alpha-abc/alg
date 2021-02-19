package main // https://leetcode.com/problems/container-with-most-water/

import "fmt"

func main() {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	var l = 0
	var r = len(height) - 1
	var maxArea = 0

	for {
		if l >= r {
			break
		}

		var b = true
		var min = height[r]
		if height[l] <= height[r] {
			b = false
			min = height[l]
		}

		var currArea = (r - l) * min
		if currArea > maxArea {
			maxArea = currArea
		}

		if b {
			r--
		} else {
			l++
		}
	}
	return maxArea
}
