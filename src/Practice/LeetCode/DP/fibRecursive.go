// Fibnacci - recursive with overlap subproblems

package main

import "fmt"

func solve(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return solve(n-1) + solve(n-2)
}
func main() {
	n := 5
	fmt.Printf("fib(%d):%d\n", n, solve(n))
	n = 8
	fmt.Printf("fib(%d):%d\n", n, solve(n))
}
