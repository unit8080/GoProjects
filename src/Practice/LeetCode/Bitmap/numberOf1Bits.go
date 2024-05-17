// 191. Number of 1 Bits
// https://leetcode.com/problems/number-of-1-bits/

// func hammingWeight(num int) int {
//     count := 0
//     for num > 0 {
//         count += 1
//         num = num & (num - 1)
//     }
//     return count
// }

// func hammingWeight(num int) int {
//     n := strconv.IntSize

//     ans := 0
//     mask := 0x1
//     for i := 0; i < n; i++ {
//         if num & mask != 0 {
//             ans++
//         }
//         mask = mask << 1
//     }
//     return ans
// }
func hammingWeight(num int) int {
    n := strconv.IntSize

    ans := 0
    mask := 0x1
    for i := 0; i < n; i++ {
        if num == 0 {
            return ans
        }
        if num & mask != 0 {
            ans++
            num = num & (^mask)
        }
        mask = mask << 1
    }
    return ans
}

