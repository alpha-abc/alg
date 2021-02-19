package main // https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/

import "fmt"

func debugPrint(s [][]int) {
	var cols = len(s)
	var rows = len(s[0])

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			fmt.Print(s[col][row], " ")
		}
		fmt.Println()
	}
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var rows = len(matrix)
	var cols = len(matrix[0])

	var sum = make([][]int, rows)
	for i := 0; i < rows; i++ {
		sum[i] = make([]int, cols)
	}

	sum[0][0] = matrix[0][0]
	for row := 0; row < rows; row++ {
		var rowSum = 0
		for col := 0; col < cols; col++ {
			rowSum += matrix[row][col]
			if row == 0 {
				sum[row][col] = rowSum
				continue
			}
			sum[row][col] = sum[row-1][col] + rowSum
		}
	}

	var check = func(sum [][]int, x1, y1, x2, y2 int, target int) bool {
		var calc = sum[x2][y2]
		if x1 == 0 && y1 == 0 {
			return calc == target
		}

		if x1 == 0 {
			return (calc - sum[x2][y1-1]) == target
		}

		if y1 == 0 {
			return (calc - sum[x1-1][y2]) == target
		}

		return sum[x2][y2]-sum[x1-1][y2]-sum[x2][y1-1]+sum[x1-1][y1-1] == target
	}

	var cnt = 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for m := 0; m <= i; m++ {
				for n := 0; n <= j; n++ {
					if check(sum, m, n, i, j, target) {
						cnt++
					}
				}
			}
		}
	}

	return cnt
}

func main() {
	var matrix = [][]int{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	}
	fmt.Println(numSubmatrixSumTarget(matrix, 0))
	matrix = [][]int{
		{1, -1},
		{-1, 1},
	}
	fmt.Println(numSubmatrixSumTarget(matrix, 0))
}
