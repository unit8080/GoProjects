// 153. Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func findMin(nums []int) int {
    
    n := len(nums)

    l := 0 
    r := n - 1
    for l < r {
        pivot := l + (r-l)/2
        if nums[pivot] > nums[r] {
            l = pivot +1
        } else {
            r = pivot
        }
    }
    return nums[l]
}
