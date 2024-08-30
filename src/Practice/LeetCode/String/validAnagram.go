// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/

// func isAnagram(s string, t string) bool {
    
//     smap := make(map[rune]int)
//     for _, ch := range s {
//         smap[ch] += 1
//     }
//     for _, ch := range t {
//         smap[ch] -= 1
//         if smap[ch] == 0 {
//             delete(smap, ch)
//         }
//     }
//     return len(smap) == 0
    
// }

func isAnagram(s string, t string) bool {
    strMap := make(map[rune]int)

    for _, str := range s {
        strMap[str] += 1
    }
    for _, str := range t {
        _, ok := strMap[str]
        if !ok {
            return false
        } else {
            strMap[str] -= 1
            if strMap[str] == 0 {
                delete(strMap, str)
            }
        }
    }
    if len(strMap) != 0 {
        return false
    } else {
        return true
    }
}

