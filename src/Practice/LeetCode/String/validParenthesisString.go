// 678. Valid Parenthesis String
// https://leetcode.com/problems/valid-parenthesis-string/

// Algorithm:
// - Use two pointer strategy.
// - traverse from both end and then count as below
// - increment "openCount" for '(' and '*' else decrement it
// - increment "closeCount" for ')' and '*' else decrement it
// - if any count falls negative any time, string is invalid

  
func checkValidString(s string) bool {
    
    openCount := 0
    closeCount := 0
    n := len(s)
    for i := 0; i < n; i++ {
        
        ch := s[i]
        if ch == '(' || ch == '*' {
            openCount++
        } else {
            openCount--
        }

        ch = s[n-i-1]
        if ch == ')' || ch == '*' {
            closeCount++
        } else {
            closeCount--
        }
        if openCount < 0 || closeCount < 0 {
            return false
        }
    }
    return true
}
