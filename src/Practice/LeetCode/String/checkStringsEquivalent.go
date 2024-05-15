// 1662. Check If Two String Arrays are Equivalent
// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    
    str1 := strings.Join(word1, "")
    str2 := strings.Join(word2, "")
    return str1 == str2 
}
