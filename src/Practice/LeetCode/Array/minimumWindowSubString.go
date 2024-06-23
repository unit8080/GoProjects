// 76. Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
    n := len(s)
    if len(t) > n {
        return ""
    }

    minWindowSize := math.MaxInt
    mp := make(map[byte]int)

    for i := 0; i < len(t); i++ {
        ch := t[i]
        mp[ch]++
    }
    i := 0
    j := 0
    start_i := 0
    countRequired := len(t)
    for j < n {
        ch := s[j]

        if mp[ch] > 0 {
            countRequired--
        }

        mp[ch]--
        for countRequired == 0 {
            windowSize := j - i + 1
            if windowSize < minWindowSize {
                minWindowSize = windowSize
                start_i = i
            }
            
            mp[s[i]]++  
            if mp[s[i]] > 0 {
                countRequired++
            }
            i++  
        }

        j++
    }
    if minWindowSize == math.MaxInt {
        return ""
    } else {
        return s[start_i : start_i + minWindowSize]
    }
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
