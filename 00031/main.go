package main // https://leetcode.com/problems/next-permutation/

import "fmt" /**
Implement next permutation,
which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible,
it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples.
Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

此题就是给出当前排列P, 求出下一个大于它的最小排列P+1
正序, 是一组排列中最小的排列, 逆序, 是一组排列中最大的排列

我们先根据结论反推逻辑
假设        p1 = 1,2,3,4,5
则   p1+1 = p2 = 1,2,3,5,4
     p2+1 = p3 = 1,2,4,3,5

我们先看p2到p3的变化, p2最后两位是逆序, 不可能有更大的排列, 只能向前考虑跟多的位,
p2后三位是3,5,4, 不是逆序, 显然存在一个比它大的排列, 我们把p2参考点的下标n定为2, 即p2[2] = 3

我们在3,5,4中, 选择比3大的最小数字, 保证了最小的首位元素, 然后交换, 得到 4,5,3, 显然4,5,3不是我们
要的下一个排列, 不难看出, 首位元素一定是4, 但是5,3是一个逆序排列, 为什么会是逆序排列? 因为我们
寻找的时候就是以第一个非逆序元素为分割点, 而4作为与3交换的元素, 又比3大, 因此交换后还是一个逆序的排列,
但是我们要的是比当前排列大的最小排列, 所以将最后的逆序排列反向得到正向排列(最小排列)
*/
func main() {
	var nums = []int{1, 2, 3, 5, 4}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	var idx = -1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			idx = i
			break
		}
	}

	for i := idx + 1; i <= len(nums); i++ {
		if idx == -1 {
			break
		}

		if i == len(nums) || nums[idx] >= nums[i] {
			swap(nums, idx, i-1)
			break
		}
	}

	for i, j := idx+1, len(nums)-1; i < j; {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, m, n int) {
	var tmp = nums[m]
	nums[m] = nums[n]
	nums[n] = tmp
}
