// 344. Reverse String
// https://leetcode.com/problems/reverse-string/description/


// func reverseString(s []byte)  {
//     left := 0
//     right := len(s) - 1
//     for left < right {
//         s[left], s[right] = s[right], s[left]
//         left++
//         right--
//     }
// }

func reverse(s *[]byte, left, right int) {
    if left >= right {
        return
    }
    // (*s)[left], (*s)[right] = (*s)[right], (*s)[left]
    reverse(s, left+1, right-1)
    (*s)[left], (*s)[right] = (*s)[right], (*s)[left]
}
func reverseString(s []byte)  {
    reverse(&s, 0, len(s)-1)
}
