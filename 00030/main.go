package main // https://leetcode.com/problems/substring-with-concatenation-of-all-words/

import "fmt"

/**
You are given a string, s, and a list of words, words, that are all of the same length.
Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

Example 1:

Input:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
  Output: [0,9]
  Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
  The output order does not matter, returning [9,0] is fine too.

Example 2:

Input:
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
  Output: []
*/

func main() {
	var s = "barfoofoobarthefoobarman"
	var words = []string{"bar", "foo", "the"}
	var arr = findSubstring(s, words)
	fmt.Println(arr)

	s = "wordgoodgoodgoodbestword"
	words = []string{"word", "good", "best", "good"}
	arr = findSubstring(s, words)
	fmt.Println(arr)

	s = "aaaaaaaa"
	words = []string{"aa", "aa", "aa"}
	arr = findSubstring(s, words)
	fmt.Println(arr)
}

// tips 滑动窗口
func findSubstring(s string, words []string) []int {
	var lens, lenwords = len(s), len(words)
	var res = make([]int, 0, lenwords)

	if lens <= 0 || lenwords <= 0 {
		return res
	}

	var lenword = len(words[0])
	// 记录每个word出现的次数
	var m = make(map[string]int, lenwords)

	for i, word := range words {
		// 条件检测
		if len(word) != lenword {
			return res
		}

		// 如果words总长度大于s, 则必定不存在结果
		if (i+1)*lenword > lens {
			return res
		}

		m[word]++
	}

	var window = make(map[string]int, lenwords)

	for i := 0; i < lenword; i++ {
		/*window left 左指针*/ /*window right 右指针*/
		var wl, wr = i, i

		// 重置window数据, 归零则符合, 复用map, 减少空间使用
		for k, v := range m {
			window[k] = v
		}

		for lens-wl >= lenword*lenwords {
			var word = s[wr : wr+lenword]
			var n, ok = window[word]

			switch {
			case !ok:
				// s中出现的word不符合words列表, 直接跳到下一个
				wl, wr = wr+lenword, wr+lenword

				for k, v := range m {
					window[k] = v
				}

			case n == 0:
				// 将要入队的word不符合words列表中的限制, 此时右指针不动, 移动左指针
				window[s[wl:wl+lenword]]++
				wl += lenword
			default:
				// word符合条件, 归入window, 重新计算, 移动右指针
				window[word]--
				wr += lenword

				// 由于右指针移动的前提都满足列表条件,
				// 如果左右指针差值刚好是列表的总长度, 则符合结果
				if wr-wl == lenwords*lenword {
					res = append(res, wl)

					window[s[wl:wl+lenword]]++
					wl += lenword
				}

			}
		}
	}

	return res
}
