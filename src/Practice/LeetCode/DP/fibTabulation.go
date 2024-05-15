package main

import "fmt"

func solve3(n int) int {

	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[0] = 0

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {

	n := 5
	fmt.Printf("fib(%d): %d\n", n, solve3(n))

	n = 8
	fmt.Printf("fib(%d): %d\n", n, solve3(n))

}
