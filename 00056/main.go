package main // https://leetcode-cn.com/problems/merge-intervals/

import "fmt" // 56. 合并区间
//
// 给出一个区间的集合，请合并所有重叠的区间。
// 示例 1:
// 	输入: [[1,3],[2,6],[8,10],[15,18]]
// 	输出: [[1,6],[8,10],[15,18]]
// 	解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
// 示例 2:
// 	输入: [[1,4],[4,5]]
// 	输出: [[1,5]]
// 	解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

func quickSort(intervals [][]int, i, j int) {
	if i >= j {
		return
	}

	var l, r, p = i, j, i

	for l < r {
		for l < r && intervals[r][0] >= intervals[p][0] {
			r--
		}

		for l < r && intervals[l][0] <= intervals[p][0] {
			l++
		}

		swap(intervals, l, r)
	}

	swap(intervals, p, l)

	quickSort(intervals, i, l-1)
	quickSort(intervals, r+1, j)
}

func swap(intervals [][]int, l, r int) {
	var tmp = intervals[l]
	intervals[l] = intervals[r]
	intervals[r] = tmp
}

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}

	quickSort(intervals, 0, len(intervals)-1)

	var resp [][]int
	var l, r = intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if r >= intervals[i][0] {
			if r < intervals[i][1] {
				r = intervals[i][1]
			}
		} else {
			resp = append(resp, []int{l, r})
			l, r = intervals[i][0], intervals[i][1]
		}
	}

	resp = append(resp, []int{l, r})
	return resp
}

func main() {
	var arr = [][]int{}
	fmt.Println(merge(arr))
}
