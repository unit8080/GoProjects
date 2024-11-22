// 2444. Count Subarrays With Fixed Bounds
// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/

func countSubarrays(nums []int, minK int, maxK int) int64 {
    
    n := len(nums)

    minKPosition := -1
    maxKPosition := -1
    culpritIdx := -1

    ans := 0
    for i := 0; i < n; i++ {
        if nums[i] < minK || nums[i] > maxK {
            culpritIdx = i
        }

        if nums[i] == minK {
            minKPosition = i
        }

        if nums[i] == maxK {
            maxKPosition = i
        }
        smaller := min(minKPosition, maxKPosition)
        temp := smaller - culpritIdx

        if temp > 0 {
            ans += temp
        }
    }
    return int64(ans)
}
