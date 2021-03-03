package main // https://leetcode-cn.com/problems/open-the-lock/

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(
		openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"),
	)
}

func plus(s string, i int) string {
	var b = make([]byte, 4)
	copy(b, s)

	if b[i] == '9' {
		b[i] = '0'
		return string(b)
	}
	b[i]++
	return string(b)
}

func minus(s string, i int) string {
	var b = make([]byte, 4)
	copy(b, s)

	if b[i] == '0' {
		b[i] = '9'
		return string(b)
	}
	b[i]--

	return string(b)
}

func openLock(deadends []string, target string) int {
	var lst = list.New()
	var deadMp = make(map[string]bool)
	for _, e := range deadends {
		deadMp[e] = true
	}

	var visited = make(map[string]bool)

	lst.PushBack("0000")
	visited["0000"] = true
	var depth = 0

	for lst.Len() > 0 {
		var sz = lst.Len()
		for i := 0; i < sz; i++ {
			var head = lst.Front()
			// 出队列
			lst.Remove(head)

			var v = head.Value.(string)

			if _, ok := deadMp[v]; ok {
				continue
			}

			if v == target {
				return depth
			}

			/*加入新节点*/
			for j := 0; j < 4; /*四个转盘*/ j++ {
				var s1 = plus(v, j)
				var s2 = minus(v, j)

				if _, ok := visited[s1]; !ok {
					lst.PushBack(s1)
					visited[s1] = true
				}
				if _, ok := visited[s2]; !ok {
					lst.PushBack(s2)
					visited[s2] = true
				}
			}
		}
		depth++
	}
	return -1
}
