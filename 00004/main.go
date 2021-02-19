package main // https://leetcode.com/problems/median-of-two-sorted-arrays/

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1 = len(nums1)
	var l2 = len(nums2)

	if (l1+l2)%2 == 1 {
		return findKth(nums1, nums2, (l1+l2)/2+1)
	}
	return (findKth(nums1, nums2, (l1+l2)/2) + findKth(nums1, nums2, (l1+l2)/2+1)) / 2.0
}

func findKth(n1, n2 []int, k int) float64 {
	var l1 = len(n1)
	var l2 = len(n2)

	if l1 > l2 {
		return findKth(n2, n1, k)
	}

	if l1 == 0 {
		return (float64)(n2[k-1])
	}

	if k == 1 {
		return (float64)(min(n1[0], n2[0]))
	}

	var i = min(k/2, l1)
	var j = k - i

	if n1[i-1] <= n2[j-1] {
		return findKth(n1[i:], n2, j)
	}
	return findKth(n1, n2[j:], i)
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}
