package main

import (
	"fmt"
	"time"
)

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	var revertedNumber = 0

	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("-----")
	}
}
