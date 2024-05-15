// 70. Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/

// TLE : 31/45
// func solveRecur1( n int) int {

//     if n < 3 {
//         return n
//     }
    
//     return solveRecur1(n-1) + solveRecur1(n-2)
// }
// func climbStairs(n int) int {
//     return solveRecur1(n)
// }

// TLE: 21/45
// func solveRecur2(numStairs int, currStair int) int {
//     if currStair == numStairs {
//         return 1
//     }
//     if currStair > numStairs {
//         return 0
//     }

//     return solveRecur2(numStairs, currStair +1) + solveRecur2(numStairs, currStair + 2)
// }
// func climbStairs(n int) int {
//     return solveRecur2(n, 0)
// }

// Approach 3: Dynamic Programming
// Algorithm
// As we can see this problem can be broken into subproblems, and 
// it contains the optimal substructure property i.e. its optimal 
// solution can be constructed efficiently from optimal solutions 
// of its subproblems, we can use dynamic programming to solve this problem.

// One can reach i^{th} step in one of the two ways:

// Taking a single step from (i−1)th step.

// Taking a step of 222 from (i−2)th step.

// So, the total number of ways to reach i^{th} is equal
// to sum of ways of reaching (i−1)th step and ways of 
// reaching (i−2)th step.

// Let dp[i] denotes the number of ways to reach on i^{th} 
// step: 

// dp[i]=dp[i−1]+dp[i−2]

func climbStairs(n int) int {
    if n < 3 {
        return n
    }
    ways := make([]int, n+1)
    ways[1] = 1
    ways[2] = 2
    for i := 3; i <= n; i++ {
        ways[i] = ways[i-1] + ways[i-2]
    }
    return ways[n]
}

