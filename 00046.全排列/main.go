package main // https://leetcode-cn.com/problems/permutations/

import "fmt" 

func permute(nums []int) [][]int {
	var resp [][]int
	var track []int

	var dfs func([]int)

	dfs = func(track []int) {
		if len(nums) == len(track) {
			var r = make([]int, len(nums))
			copy(r, track)
			resp = append(resp, r)
			return
		}

		for _, n := range nums {
			if contains(track, n) {
				continue
			}

			track = append(track, n)
			dfs(track)
			track = track[:len(track)-1]
		}
	}

	dfs(track)
	return resp
}

func contains(track []int, n int) bool {
	for _, e := range track {
		if n == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(
		permute([]int{1, 2, 3}),
	)
}

