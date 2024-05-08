// 921. Minimum Add to Make Parentheses Valid
// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/

func minAddToMakeValid(s string) int {
    var left, right int

    if len(s) == 0 {return 0}

    for pos := 0; pos < len(s); pos++ {
        if s[pos] == '(' {
            left++
        } else if s[pos] == ')' {
            if left > 0 {
                left--
            } else {
                right++
            }
        }
    }
    return left+right
}
