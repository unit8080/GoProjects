// 209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) int {
    
    n := len(nums)
    minLen := math.MaxInt

    i := 0
    j := 0
    sum := 0
    for j < n {
        sum = sum + nums[j]
        for sum >= target {
            minLen = min(minLen, j -i +1)

            sum -= nums[i]
            i++
        }
        j++
    }
    if minLen == math.MaxInt {
        return 0
    }
    return minLen
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

