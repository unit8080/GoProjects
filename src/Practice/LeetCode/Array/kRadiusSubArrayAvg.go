// 2090. K Radius Subarray Averages
// https://leetcode.com/problems/k-radius-subarray-averages/

func getAverages(nums []int, k int) []int {
    
    n := len(nums)
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i] = -1
    }
    left := 0
    right := 2 *k
    i := k

    if k >= n || right >= n {
        return ans
    }
    windowSum := 0
    for x := left; x <= right; x++ {
        windowSum += nums[x] 
    }
    avg := windowSum/(2*k + 1)
    ans[i] = avg
    right++
    i++
    for right < n {
        windowSum += nums[right]
        windowSum -= nums[left]
        avg := windowSum/(2*k + 1)
        ans[i] = avg
        left++
        i++
        right++
    }
    return ans
}
