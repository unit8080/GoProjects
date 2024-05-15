// Space optimization SC - O(1)

package main

import "fmt"

func solve4(n int) int {
	if n < 2 {
		return n
	}
	prev1 := 0
	prev2 := 1
	curr := 0
	for i := 2; i <= n; i++ {
		curr = prev1 + prev2
		prev1 = prev2
		prev2 = curr
	}
	return curr
}
func main() {

	n := 5
	fmt.Printf("fib(%d): %d\n", n, solve4(n))
	n = 8
	fmt.Printf("fib(%d): %d\n", n, solve4(n))
}
