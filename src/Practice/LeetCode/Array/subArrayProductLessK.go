// 713. Subarray Product Less Than K
// https://leetcode.com/problems/subarray-product-less-than-k/

func numSubarrayProductLessThanK(nums []int, k int) int {
	
    if k <= 1 {
        return 0
    }
    ans := 0
	left := 0
	right := 0
	product := 1
	
	for right < len(nums) {

        product = product * nums[right]
        for product >= k {
            product = product / nums[left]
            left++
        }

        ans += right - left + 1
		right++
	}
	return ans
}
