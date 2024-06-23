// 78. Subsets
// https://leetcode.com/problems/subsets/

import "math"
func subsets(nums []int) [][]int {

    ans := [][]int{}

    n := len(nums)
    for i := math.Pow(2, float64(n)); i < math.Pow(2, float64(n+1)); i++ {
        bitmask := fmt.Sprintf("%b", int(i))[1:]

        output := []int{}
        for j := 0; j < n; j++ {
            if bitmask[j] == '1' {
                output = append(output, nums[j])
            }
        }
        ans = append(ans, output)
    }
    return ans
}
