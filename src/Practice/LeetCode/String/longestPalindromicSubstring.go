// 5. Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/

func checkPalindrome(start, end int, s string) bool {
    i := start
    j := end - 1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}
func longestPalindrome(s string) string {
    for length := len(s); length > 0; length-- {
        for start := 0; start <= len(s) - length; start++ {
            if checkPalindrome(start, start+length, s) {
                return s[start:start+length]
            }
        }
    }
    return ""
}
