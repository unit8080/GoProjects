// 1464. Maximum Product of Two Elements in an Array
// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/

func maxProduct(nums []int) int {
    
    largest := 0
    sec_largest := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] > largest {
            sec_largest = largest
            largest = nums[i]
        } else {
            sec_largest = max(sec_largest, nums[i])
        }
    }
    return (largest -1) * (sec_largest -1)
}
