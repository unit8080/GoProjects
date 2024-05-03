// 2330. Valid Palindrome IV
// https://leetcode.com/problems/valid-palindrome-iv/

func makePalindrome(s string) bool {
    length := len(s)

    if length < 2 {
        return true
    }
    left := 0
    right := length -1

    used := 0
    for left < right {
        lchar := s[left]
        rchar := s[right]

        if lchar != rchar {
            used++
            if used > 2 {
                return false
            }
        }
        left++
        right--
    }
    return true
}
