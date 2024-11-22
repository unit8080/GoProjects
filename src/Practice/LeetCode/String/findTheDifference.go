// 389. Find the Difference
// https://leetcode.com/problems/find-the-difference/

func findTheDifference(s string, t string) byte {
    var result byte
    for i := 0; i < len(s); i++ {
        result ^= s[i]
    }
    for i := 0; i < len(t); i++ {
        result ^= t[i]
    }
    return result
}

// using ASCII difference
// func findTheDifference(s string, t string) byte {

//     var snum int = 0

//     for i := 0; i < len(s); i++ {
//         snum +=  int (byte(s[i]) - byte('a'))
//     }
//     var tnum int = 0
//     for i := 0; i < len(t); i++ {
//         tnum += int (byte(t[i]) - byte('a'))
//     }
//     return byte (tnum -snum + int (byte('a')))
// }

// using map
func findTheDifference(s string, t string) byte {
    strMap := map[byte] int {}

    for i := 0; i < len(s); i++ {
        strMap[s[i]] += 1
    }
    for i := 0; i < len(t); i++ {
        if _, exists := strMap[t[i]]; !exists {
            return t[i]
        } else {
            strMap[t[i]] -= 1
            if strMap[t[i]] == 0 {
                delete(strMap, t[i])
            }

        }
    }
    return byte('0')
}
