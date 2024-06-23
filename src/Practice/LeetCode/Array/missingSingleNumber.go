// 136. Single Number
// https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
    missing := 0
    for _, num := range nums {
        missing = missing ^ num
    }
    return missing
}
