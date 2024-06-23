// 1328. Break a Palindrome
// https://leetcode.com/problems/break-a-palindrome/

// Given a palindromic string of lowercase English letters palindrome, replace
// exactly one character with any lowercase English letter so that the resulting 
// string is not a palindrome and that it is the lexicographically smallest one possible.

// Return the resulting string. If there is no way to replace a character to make 
// it not a palindrome, return an empty string.

// A string a is lexicographically smaller than a string b (of the same length) 
// if in the first position where a and b differ, a has a character strictly 
// smaller than the corresponding character in b. For example, "abcc" is 
// lexicographically smaller than "abcd" because the first position they 
// differ is at the fourth character, and 'c' is smaller than 'd'.

// Algorithm:
// This is quite a simple problem:
// (1) You cannot do anything about a single length string.
// Return "".
// (2) For the first half of the string, check if there's a character
// greater than 'a' & replace it by 'a'
// (3) ALL the first half characters are 'a's. Hence, replace the
// LAST element of the string by 'b'.

func breakPalindrome(p string) string {
    if len(p) < 2 {
        return ""
    }

    i := 0
    for i < len(p)/2 {
        if p[i] != 'a' {
            pbyte := []byte(p)
            pbyte[i] = 'a'
            str := string(pbyte)
            return str
        }
        i++
    }

    str := p[:len(p) -1]
    str = str + "b"
    return str

}
