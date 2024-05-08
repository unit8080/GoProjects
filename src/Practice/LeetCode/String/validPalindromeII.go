// 680. Valid Palindrome II
// https://leetcode.com/problems/valid-palindrome-ii/

func validPalindrome(s string) bool {
    
    length := len(s)
    if length < 2 {
        return true
    }

    left := 0
    right := length -1
    for left < right {
        if s[left] != s[right] { // first mismatch
            return isPalindrome(s, left, right -1) || isPalindrome(s, left+1, right) 
        }
        left++
        right--
    }
    return true
}

func isPalindrome(s string, left, right int ) bool {

    if left > right { return false } // important !!!

    for left < right {
    
        if s[left] != s[right] { // second mismatch !!!
            return false
        } 
        left++
        right--
    }
    return true
}
