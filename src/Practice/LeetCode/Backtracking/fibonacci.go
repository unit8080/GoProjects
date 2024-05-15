// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/

// func fib(n int) int {
//     if n < 0 {return n}
//     if n < 2 { return n}
//     f0 := 0
//     f1 := 1
//     fn := f0 + f1
//     for i := 2; i < n; i++ {
        
//         f0 = f1
//         f1 = fn
//         fn = f0 + f1
//     }
//     return fn
// }

// Recursion - plain
// func fib(n int) int {

//     if n == 0 {
//         return 0
//     }
//     if n == 1 {
//         return 1
//     }
//     return fib(n-1) + fib(n-2)
// }

// Recusion + Memoization
// func fibMem(n int, dp []int) int {
//     if n < 2 {
//         return dp[n]
//     }
//     if dp[n] != 0 {
//         return dp[n]
//     }

//     dp[n] = fibMem(n-1, dp) + fibMem(n-2, dp)
//     return dp[n]
// }
// func fib(n int) int {
//     dp := make([]int, n+1)
//     if n < 2 {
//         return n
//     }
//     dp[0] = 0
//     dp[1] = 1
//     ans := fibMem(n, dp)
//     return ans
// }

// With Tabulation
// func fib(n int) int {
//     // base case
//     if n < 2 {
//         return n
//     }
//     dp := make([]int, n+1)
//     dp[0] = 0
//     dp[1] = 1
    
//     for i := 2; i <=n; i++ {
//         dp[i] = dp[i-1] + dp[i-2]
//     }
//     return dp[n]
// }

// Space optimization
func fib(n int) int {
    if n < 2 {
        return n
    }
    prev1 := 1
    prev2 := 0

    for i := 2; i <= n; i++ {
        curr := prev1 + prev2
        prev2 = prev1
        prev1 = curr
    }
    return prev1
}

