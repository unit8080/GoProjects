// 1160. Find Words That Can Be Formed by Characters
// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/

func countCharacters(words []string, chars string) int {
    
    var charCount [26]int

    counts := 0
    for i := 0; i < len(chars); i++ {
        charCount[int(chars[i]- byte('a'))]++
    }
    for _, str := range words {
        var strCount [26]int
        for i := 0; i < len(str); i++ {
            strCount[int(str[i] - byte('a'))]++
        }
        ok := true
        for i := 0; i < 26; i++ {
            if strCount[i] > charCount[i] {
                ok = false
            }
        }
        if ok {
            counts += len(str)
        }
    }
    return counts
}

