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
  - one for traversing and reading zeros
  - another for writing the non-zero number
  - at the end, write remaing write index with zero
*/
package main

func moveZeroes(nums []int) {

	size := len(nums)
	if size < 2 {
		return
	}

	rIdx := 0
	wIdx := 0

	for rIdx < size {
		if nums[rIdx] == 0 {
			//skip

		} else {
			if rIdx != wIdx {
				nums[wIdx] = nums[rIdx]
			}
			wIdx++
		}
		rIdx++
	}
	for wIdx < size {
		nums[wIdx] = 0
		wIdx++
	}
}
