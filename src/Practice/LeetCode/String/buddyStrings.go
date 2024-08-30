// 859. Buddy Strings
// https://leetcode.com/problems/buddy-strings/

// buddyStrings checks if two strings can become equal by swapping two letters
func  buddyStrings(a string, b string) bool {
    
    if len(a) != len(b) {
        return false
    }

    chMap := make(map[rune] bool)
    if a == b {
        for _, ch := range a {
            if chMap[ch] == true {
                return true
            }
            chMap[ch] = true
        }
        return false
    }
    idx := []int{}
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            idx = append(idx, i)
        }
        if len(idx) > 2 {
            return false
        }
    }

    if len(idx) == 2 && a[idx[0]] == b[idx[1]] && a[idx[1]] == b[idx[0]] {
        return true
    } else {
        return false 
    }
}


