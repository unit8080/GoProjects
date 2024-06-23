// 198. House Robber
// https://leetcode.com/problems/house-robber/

// Top down - Recursion + Memoization
func rob(nums []int) int {
    n := len(nums)
    memo := make([]int, n)
    for i, _ := range memo { memo[i]=-1 }
    return solve(nums,  0, n, memo)
    
}

func solve(nums []int, i, n int, memo []int) int {
    if i >= n {
        return 0
    }
    if memo[i] != -1 {
        return memo[i]
    }

    steal := nums[i] + solve(nums, i+2, n, memo)
    skip := solve(nums, i+1, n, memo)
    memo[i] = max(steal, skip)
    return memo[i]
}
func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

// Bottom Up 
// func rob(nums []int) int {
//     n := len(nums)
//     t := make([]int, n+1)
//     // t[i] : maximum stolen till house "i"
    
//     // no house: i = 0
//     t[0] = 0

//     // 1 house: i = 1
//     t[1] = nums[0]

//     for i := 2; i <= n; i++ {
//         steal :=  nums[i-1] + t[i-2]
//         skip := t[i-1]
//         t[i] = max(steal, skip)
//     }
//     return t[n]
// }
// func max(a, b int) int {
//     if a > b {
//         return a
//     } else {
//         return b
//     }
// }

