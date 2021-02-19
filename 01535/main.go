package main // https://leetcode-cn.com/problems/find-the-winner-of-an-array-game/

import "fmt"

func getWinner(arr []int, k int) int {
	var win, cnt = arr[0], 0

	for i := 1; i < len(arr) && cnt < k; i++ {
		if win > arr[i] {
			cnt++
			continue
		}

		win = arr[i]
		cnt = 1
	}

	return win
}

func main() {
	var arr = []int{2, 1, 3, 5, 4, 6, 7}
	var k = 2
	fmt.Println(getWinner(arr, k))
}
