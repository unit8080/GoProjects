// 266. Palindrome Permutation
// https://leetcode.com/problems/palindrome-permutation/

func canPermutePalindrome(s string) bool {
    charMap := make(map[string]struct{})
    for _, str := range s {
        ch := string(str)
        if _, ok := charMap[ch]; ok {
            delete(charMap, ch)
        } else {
            charMap[ch] = struct{}{}
        }
    }
    return len(charMap) == 0 || (len(charMap) == 1 && len(s) % 2 == 1)
}

func canPermutePalindrome(s string) bool {
    charMap := make(map[rune]struct{})
    for _, ch := range s {

        if _, ok := charMap[ch]; ok {
            delete(charMap, ch)
        } else {
            charMap[ch] = struct{}{}
        }
    }
    return len(charMap) == 0 || (len(charMap) == 1 && len(s) % 2 == 1)
}
