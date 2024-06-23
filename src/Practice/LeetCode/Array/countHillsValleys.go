// 2210. Count Hills and Valleys in an Array
// https://leetcode.com/problems/count-hills-and-valleys-in-an-array/

func countHillValley(nums []int) int {
    n := len(nums)
    prev := nums[0]
    ans := 0

    for i := 1; i < n; i++ {
        for i < n && nums[i] == nums[i-1] {
            i++
        }
        if i < n && (nums[i-1] > nums[i] && nums[i-1] > prev ||
            nums[i-1] < nums[i] && nums[i-1] < prev) {
                ans++
        }
        prev = nums[i-1]
    } 
    return ans
}
