package main // https://leetcode.com/problems/integer-to-roman/

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intToRoman(1000))
}

func intToRoman(num int) string {
	if num <= 0 {
		panic("illegality number")
	}

	// - 相同的数字连写，所表示的数等于这些数字相加得到的数，如 Ⅲ=3；
	// - 小的数字在大的数字的右边，所表示的数等于这些数字相加得到的数，如 Ⅷ=8、Ⅻ=12；
	// - 小的数字（限于 Ⅰ、X 和 C）在大的数字的左边，所表示的数等于大数减小数得到的数，如 Ⅳ=4、Ⅸ=9；
	// - 在一个数的上面画一条横线，表示这个数增值 1,000 倍
	// I = 1 //
	// V = 5
	// X = 10 //
	// L = 50
	// C = 100 //
	// D = 500
	// M = 1000

	var I = 1
	var X = 10
	var C = 100
	var M = 1000
	var m = []string{"", "M"}
	var c = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	var x = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	var i = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	var buffer = &bytes.Buffer{}

	var tmpNum = num
	if tmpNum >= M {
		var n = tmpNum / M
		for j := 0; j < n; j++ {
			buffer.WriteString(m[1])
		}
	}

	tmpNum = tmpNum % M
	buffer.WriteString(c[tmpNum/C])

	tmpNum = tmpNum % C
	buffer.WriteString(x[tmpNum/X])

	tmpNum = tmpNum % X
	buffer.WriteString(i[tmpNum/I])

	return buffer.String()
}
