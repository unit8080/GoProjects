// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/

// Algorithm:
// 1. Notice that input is always lower case 26 alphabet characters
// 2. build the unique array of charaters
// 3. Define Key array of 26 bytes for alphabets
// 4. count the each alphabets characters
// 5. use map to store the str with above key
// return the values in map as result

// func groupAnagrams(strs []string) [][]string {
    
//     var ans [][]string
//     group := make(map[string][]string)

//     for _, str := range strs {
//         hash := ""
//         alpha := make([]int, 26)
//         for _, val := range str {
//             alpha[val - 'a'] += 1
//         }
//         for i := 0; i < 26; i++ {
//             for alpha[i] > 0 {
//                 hash = hash + string(i + 'a')
//                 alpha[i]--
//             }
//         }
//         group[hash] = append(group[hash], str)
        
//         // chars := strings.Split(str, "")
//         // sort.Strings(chars)
//         // strSorted := strings.Join(chars, "")
//         // group[strSorted] = append(group[strSorted], str)
//     }
//     for _, v := range group {
//         ans = append(ans, v)
//     }
//     return ans
// }

// func groupAnagrams(strs []string) [][]string {

//     ans := [][]string{}

//     strMap := make(map[string][]string)
//     for _, str := range strs {
//         // compute hash key
//          a := make([]int, 26)
//         for _, s := range str {
//             ch := int(s - 'a')
//             a[ch] += 1
//         }
//         hash := ""
//         for _, v := range a {
//             hash += string(v)
//         }
//         strMap[hash] = append(strMap[hash] , str)
//     }

//     for _, v := range strMap {
//         ans = append(ans, v)
//     }
//     return ans
// }

// type Key [26]byte

// func strKey(str string) (key Key) {
//     for _, v := range str {
//         key[v - 'a'] += 1
//     }
//     return key
// }
// func groupAnagrams(strs []string) [][]string {

//     ans := [][]string{}
//     group := map[Key][]string{}

//     for _, str := range strs {
//         key := strKey(str)
//         group[key] = append(group[key], str)
//     }
//     for _, val := range group {
//         ans = append(ans, val)
//     }
//     return ans
// }

type Key [26]byte

func getKey(str string) (key Key) {
    for _, v := range str {
        key[v - 'a'] += 1
    }
    return key
}
func groupAnagrams(strs []string) [][]string {
    
    ans := [][]string{}

    group := map[Key][]string{}

    for _, str := range strs {
        key := getKey(str)
        group[key] = append(group[key], str)
    }
    for _, v := range group {
        ans = append(ans, v)
    }
    return ans
}
