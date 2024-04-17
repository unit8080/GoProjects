func validPalindrome(s string) bool {
    
    length := len(s)
    if length < 2 {
        return true
    }

    left := 0
    right := length -1
    for left < right {
        if s[left] != s[right] {
            return isPalindrome(s, left, right -1) || isPalindrome(s, left+1, right) 
        }
        left++
        right--
    }
    return true
}

func isPalindrome(s string, left, right int ) bool {
    if left > right { return false }
    if s == "" { return true }

    for left < right {
    
        if s[left] != s[right] {
            return false
        } 
        left++
        right--
    }
    return true
}
