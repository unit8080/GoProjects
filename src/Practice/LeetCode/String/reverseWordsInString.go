// 151. Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
    
    //split by spaces
    str := strings.Fields(s)
    l := 0
    r := len(str) -1
    for l < r { // reverse here
        str[l], str[r] = str[r], str[l]
        l++
        r--
    }
    // join the words with a space
    return strings.Join(str, " ")
}
