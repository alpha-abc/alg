package main // https://leetcode-cn.com/problems/fibonacci-number/

func main() {
	println(fib(100))
}

/* .   /
       |-> 0; (n = 0)
f(n) = |-> 1; (n = 1)
       |-> f(n-1) + f(n-2); (n > 1)
       \
*/

func fib(n int) int {
	var n1, n2 = 0, 1
	if n == 0 {
		return n1
	}

	if n == 1 {
		return n2
	}

	var curr = 0
	for i := 2; i < n+1; i++ {
		curr = n1 + n2

		n1 = n2
		n2 = curr
	}

	return curr

}
