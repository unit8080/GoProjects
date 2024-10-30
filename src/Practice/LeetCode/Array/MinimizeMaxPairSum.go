// 1877. Minimize Maximum Pair Sum in Array
// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/

func minPairSum(nums []int) int {
    
    ans := 0 
    sort.Ints(nums)
    n := len(nums)

    l := 0
    r := n -1
    for l < r {
        sum := nums[l] + nums[r]
        ans = max(ans, sum)
        l++
        r--
    }
    return ans
}
