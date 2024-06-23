// 674. Longest Continuous Increasing Subsequence
// https://leetcode.com/problems/longest-continuous-increasing-subsequence/

func findLengthOfLCIS(nums []int) int {
    ans := -1
    l := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            l++
        } else {
            if l > ans {
                ans = l
            } 
            l = 1
        }
    }

    if l > ans {
        ans = l
    }
    return ans
}
