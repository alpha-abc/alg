package main // https://leetcode.com/problems/reverse-integer/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-120))
}

func reverse(x int) int {
	var num = 0

	for x != 0 {
		num = num*10 + x%10
		x /= 10
	}

	if num < math.MinInt32 || num > math.MaxInt32 {
		return 0
	}

	return num
}
