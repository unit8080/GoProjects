// 2595. Number of Even and Odd Bits
// https://leetcode.com/problems/number-of-even-and-odd-bits/

func evenOddBit(n int) []int {
    fmt.Println(strconv.IntSize)
    size := strconv.IntSize
    odd := 0
    even := 0
    mask := 0x1
    for i := 0; i < size; i++ {
        if i % 2 == 0 {
            if n & mask != 0 {
                even++
            }
        } else {
            if n & mask != 0 {
                odd++
            }
        }
        mask = mask << 1
    }
    return []int{even, odd}
}
