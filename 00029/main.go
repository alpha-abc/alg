package main // https://leetcode.com/problems/divide-two-integers/

/*
Given two integers dividend and divisor,
divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.
The integer division should truncate toward zero.

Example 1:
	Input: dividend = 10, divisor = 3
	Output: 3

Example 2:
	Input: dividend = 7, divisor = -3
	Output: -2

Note:
	Both dividend and divisor will be 32-bit signed integers.
	The divisor will never be 0.
	Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
	For the purpose of this problem, assume that your function returns 2^31 − 1 when the division result overflows.
*/

import "fmt"

import "math"

func main() {
	var a, b = -31, 4
	fmt.Println(" >", divide(a, b))
	fmt.Println(">>", a/b)

}

//  m   n
// -17 - 4   1 2^0
// -13 - 8   2 2^1
// - 5 -16
// - 5 - 4   1 2^0

//   m   n
// -32 - 4   1 2^0
// -28 - 8   2 2^1
// -20 -16   4 2^2
// - 4 -32
// - 4 - 4   1 2^0

func divide(dividend int, divisor int) int {
	if divisor == 0 || (divisor == -1 && dividend == math.MinInt32) {
		return math.MaxInt32
	}

	var b = (dividend ^ divisor) >= 0

	var m = dividend
	if dividend >= 0 {
		m = ^dividend - -1
	}

	var n, k = divisor, divisor
	if divisor >= 0 {
		n = ^divisor - -1
		k = n
	}

	var l = 0
	var r = 0

	for m <= n {
		r -= -(1 << l)
		l -= -1

		m -= n
		n <<= 1

		if m > n && n < k {
			n = k
			l = 0
		}
	}

	if b {
		return r
	}

	return -r
}
