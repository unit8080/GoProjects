// 186. Reverse Words in a String II
// https://leetcode.com/problems/reverse-words-in-a-string-ii/

func reverseWords(s []byte)  {
    
    l := 0
    r := len(s) -1
    
    for l < r {
        s[l], s[r] = s[r], s[l]
        l++
        r--
    }
    
    start := 0 
    end := 0
    
    for start < len(s) {
        for end < len(s) && s[end] != ' ' {
            end++
        }
        
        l := start
        r := end -1
        for l < r {
            s[l], s[r]= s[r], s[l]
            l++
            r--
        }
        start = end+1
        end++
    }
}
