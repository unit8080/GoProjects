// recursion with memoization

package main

import "fmt"

func solve2(n int, dp *[]int) int {
	if n == 0 {
		(*dp)[n] = 0
		return 0
	}
	if n == 1 {
		(*dp)[n] = 1
		return 1
	}

	if (*dp)[n] != 0 {
		return (*dp)[n]
	}

	(*dp)[n] = solve2(n-1, dp) + solve2(n-2, dp)

	return (*dp)[n]
}

func main() {

	var dp []int

	n := 5
	dp = make([]int, n+1)
	fmt.Printf("fib(%d): %d\n", n, solve2(n, &dp))
	n = 8
	dp = make([]int, n+1)
	fmt.Printf("fib(%d): %d\n", n, solve2(n, &dp))

}
