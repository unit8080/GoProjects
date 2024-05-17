// 231. Power of Two
// https://leetcode.com/problems/power-of-two/

func isPowerOfTwo(n int) bool {
    if n == 0 { // "0" is NOT power of 2
        return false
    }
    return n & (n-1) == 0
}
