// 65. Valid Number
// https://leetcode.com/problems/valid-number/

func isNumber(s string) bool {
    
    seenDigit := false
    seenDot := false
    seenExponent := false

    for i := 0; i < len(s); i++ {
        ch := s[i]
        if ch == '-' || ch == '+' {
            if i == 0  {
                // return false
                continue
            }
            if i > 0 && (s[i-1] == 'e' || s[i-1] == 'E') {
                //return false
                continue
            }
            return false

        } else if ch == '.'  {
            if seenDot || seenExponent {
                return false
            }
            seenDot = true
        } else if ch == 'e' || ch == 'E' {
            if seenExponent {
                return false
            }

            if !seenDigit {
                return false
            }
            seenExponent = true
            seenDigit = false
        } else if ch < '0' || ch > '9' {
            return false
        } else {
            seenDigit = true
        }
    }

    return seenDigit

}
