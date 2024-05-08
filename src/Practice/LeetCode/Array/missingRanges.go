// 163. Missing Ranges
// https://leetcode.com/problems/missing-ranges/

func findMissingRanges(nums []int, lower int, upper int) [][]int {
    ans := make([][]int, 0)

    n := len(nums)
    if n == 0 {
        return [][]int{{lower, upper}}
    }
    if lower < nums[0]  {
        ans = append(ans, []int{lower, nums[0]-1})
    }

    // middle part
    for i := 0; i <= n -2; i++ {
        if nums[i+1] - nums[i] > 1 {
            ans = append(ans, []int{nums[i] + 1, nums[i+1]-1})
        }
    }

    if nums[n-1] < upper  {
        ans = append(ans, []int{nums[n-1]+1, upper})
    }
    return ans
}
