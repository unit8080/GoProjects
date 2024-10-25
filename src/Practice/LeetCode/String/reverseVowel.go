// 345. Reverse Vowels of a String
// https://leetcode.com/problems/reverse-vowels-of-a-string/

func reverseVowels(s string) string {
    b := []byte (s)
    len := len(b)

    i := 0
    j := len -1

    for i < j {
        if !isVowel(b[i]) {
            i++
        } else if !isVowel(b[j]) {
            j--
        } else {
            b[i], b[j] = b[j], b[i]
            i++
            j--
        }
    }
    return string(b)
}
func isVowel(ch byte) bool {

    if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
     ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
            return true
    }
    return false
}
