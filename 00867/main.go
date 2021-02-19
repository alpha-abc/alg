package main // https://leetcode-cn.com/problems/transpose-matrix/

import (
	"fmt"
)

func main() {
	var arr = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var resp = transpose(arr)
	fmt.Println(resp)
}

func transpose(A [][]int) [][]int {
	if A == nil {
		return nil
	}

	var row, col = len(A), len(A[0])
	var arr = make([][]int, col)

	for i := 0; i < col; i++ {
		arr[i] = make([]int, row)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			arr[j][i] = A[i][j]
		}
	}

	return arr
}

