// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/

/*
Move Zeroes

Given an integer array nums, move all 0's to the end of it while
maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]

Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

Follow up: Could you minimize the total number of operations done?
*/

/*
Algorithm:
    - use two pointers.
    - one for traversing and looking for zeros - rdIdx
    - another for writing the non-zero number  - wrIdx
    - at the end, write remaing write indexes upto size-1 with zero
*/

func moveZeroes(nums []int)  {
    rdIdx := 0
    wrIdx := 0

    for rdIdx < len(nums) {

        if nums[rdIdx] != 0 {
            nums[wrIdx] = nums[rdIdx]
            wrIdx++
            rdIdx++
        } else {
            rdIdx++
        }
    }
    for wrIdx < len(nums) {
        nums[wrIdx] = 0
        wrIdx++
    }
}


func moveZeroes(nums []int)  {
    size := len(nums)
    if size < 2 {return}

    rdIdx := 0
    wrIdx := 0

    for rdIdx < size {
        if nums[rdIdx] == 0 {
            //skip - do nothing
            rdIdx++
        } else {
            if rdIdx != wrIdx { // shift non-zero num
                nums[wrIdx] = nums[rdIdx]
            }
            wrIdx++
            rdIdx++
        }
    }
    for wrIdx < size { // some left over to fill zeros
        nums[wrIdx] = 0
        wrIdx++
    }
}
