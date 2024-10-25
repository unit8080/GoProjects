// 26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {

    len := len(nums)
    i := 0
    j := 1

    for j < len {

        if nums[i] != nums[j] {
            i++
            nums[i] = nums[j]
        }
        j++
    }
    return i+1
}
