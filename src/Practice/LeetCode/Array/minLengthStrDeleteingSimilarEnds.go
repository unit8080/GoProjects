// 1750. Minimum Length of String After Deleting Similar Ends
// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/

func minimumLength(s string) int {
    n := len(s)

    i := 0
    j := n -1

    for i < j && s[i] == s[j] {
        ch := s[i]
        for i <= j && s[i] == ch {
            i++
        }
        for j >= i && s[j] == ch {
            j--
        }
    }
    return j -i + 1
}
