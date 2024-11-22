// 1456. Maximum Number of Vowels in a Substring of Given Length
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

func isVowel( ch byte) bool {
    if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
        return true
    }
    return false
}
func maxVowels(s string, k int) int {
    
    n := len(s)

    i := 0
    j := 0
    count := 0
    ans := 0

    for j < n {
        if isVowel(s[j]) {
            count++   
        }
        if j -i +1 == k {
                ans = max(ans, count)
                if isVowel(s[i]) {
                    count--
                }
                i++
            }
        j++
    }
    return ans
}
