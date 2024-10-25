// 1498. Number of Subsequences That Satisfy the Given Sum Condition
// https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
func numSubseq(nums []int, target int) int {

    sort.Ints(nums)
    
    n := len(nums)
    l := 0 
    r := n -1
    M := 1e9+7
    result := 0

    pow := make([]int, n)
    pow[0] = 1
    for i := 1; i < n; i++ {
        pow[i] = pow[i-1]*2 % int(M)
    }
    for l <= r {
        if nums[l] + nums[r] <= target {
            diff := r -l
            result = int (result % int(M) + pow[diff] % int(M))
            l++
        } else {
            r--
        }
    }
    return result
}
