// 791. Custom Sort String
// https://leetcode.com/problems/custom-sort-string/

func customSortString(order string, s string) string {
    
    ans := ""
    // frequency map
    cMap := make(map[rune]int)
    for _, ch := range s {
        cMap[ch] += 1
    }

    for _, ch := range order {
        for cMap[ch] > 0 {
            ans += string(ch)
            cMap[ch] -= 1
        }
    }
    for ch, _ := range cMap {
        for cMap[ch] > 0 {
            ans += string(ch)
            cMap[ch] -= 1
        }
    }
    return ans
}

