// 338. Counting Bits
// https://leetcode.com/problems/counting-bits/

func countBits(n int) []int {
    var ans []int
    ans = make([]int, n+1)
    for i := 0; i <= n; i++ {
        count := bitCount(i)
        ans[i] = count
    }
    return ans
}
func bitCount (n int) int {
    ans := 0
    for n > 0 {
        ans++
        n = n & (n-1)
    }
    return ans
}
