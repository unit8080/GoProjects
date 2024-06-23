// 896. Monotonic Array
// https://leetcode.com/problems/monotonic-array/

func isMonotonic(nums []int) bool {
    increasing := true
    decreasing := true

    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i+1] {
            increasing = false
        } else if nums[i] < nums[i+1] {
            decreasing = false

        }
    }
    return increasing || decreasing
}
