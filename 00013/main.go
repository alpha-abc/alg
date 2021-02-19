package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}

func romanToInt(s string) int {
	var r = []rune(s)
	var n = len(r)

	var number = 0

	var m = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < n; {
		var first = r[i]
		var second = 'N'
		if i < n-1 {
			second = r[i+1]
		}

		var f, okF = m[first]
		var s, okS = m[second]

		if !okF || (!okS && second != 'N') {
			panic("illegality")
		}

		if f >= s {
			number += f
			i++
		} else {
			number += (s - f)
			i += 2
		}
	}
	return number
}
