// 154. Find Minimum in Rotated Sorted Array II
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

func findMin(nums []int) int {
    n := len(nums)

    l := 0
    r := n -1

    for l < r {
        pivot := l + (r -l)/2
        if nums[pivot] > nums[r] {
            l = pivot +1
        } else if nums[pivot] == nums[r] { // !!!! Important
            r = r -1 // nums = [3,3,1,3]
        } else {
            r = pivot
        }
    }
    return nums[r]
}
