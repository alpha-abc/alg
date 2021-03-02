package main // https://leetcode-cn.com/problems/n-queens/

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

/*

 */

func solveNQueens(n int) [][]string {
	var resp [][]string

	var track1 = make([][]byte, n)
	for i := 0; i < n; i++ {
		var b = make([]byte, n)
		for i := 0; i < n; i++ {
			b[i] = '.'
		}
		track1[i] = b
	}

	var dfs func([][]byte, int)
	dfs = func(track [][]byte, row int) {
		if n == row /*超过最后一行*/ {
			var r = make([][]byte, n)
			copy(r, track)

			var r1 []string
			for _, bs := range r {
				r1 = append(r1, string(bs))
			}
			resp = append(resp, r1)
		}

		for j := 0; j < n; j++ {
			if !check(track, row, j) {
				// 忽略无效放置
				continue
			}

			track[row][j] = 'Q'
			dfs(track, row+1)
			track[row][j] = '.'

		}
	}

	dfs(track1, 0)
	return resp
}

func check(track [][]byte, row, col int) bool {
	var n = len(track)

	// 检测正上方
	for i := 0; i < row; i++ {
		if track[i][col] == 'Q' {
			return false
		}
	}

	// 检测左上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if track[i][j] == 'Q' {
			return false
		}

		i--
		j--
	}

	// 检测右上方
	for i, j := row-1, col+1; i >= 0 && j < n; {
		if track[i][j] == 'Q' {
			return false
		}

		i--
		j++
	}
	return true
}

func printTrack(t [][]byte) {
	for _, bs := range t {
		fmt.Println(string(bs))
	}

	fmt.Println()
}
