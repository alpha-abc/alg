package main // https://leetcode-cn.com/problems/subarray-sum-equals-k/

import "fmt"

func subarraySum1(nums []int, k int) int {
	var resp = 0

	for i := 0; i < len(nums); i++ {
		var cnt = 0
		for j := i; j < len(nums); j++ {
			cnt += nums[j]
			if cnt == k {
				resp++
				continue
			}
		}
	}

	return resp
}

/*
核心思想
idx:  0  1  2  3  4
arr:  2  3  5  9  1
sum:  2  5 10 19 20

sum[i] = sum[i-1] + arr[i]

sum[i] = sum[j] + k
sum[j] = sum[i] - k
*/
func subarraySum(nums []int, k int) int {
	var cnt, sum = 0, 0
	var mp = make(map[int]int)

	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := mp[sum-k]; ok {
			cnt += mp[sum-k]
		}
		mp[sum]++
	}

	return cnt
}

func main() {
	var arr = []int{1, 1, 1}
	var k = 2
	fmt.Println(subarraySum(arr, k))
}
