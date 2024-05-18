// 2595. Number of Even and Odd Bits
// https://leetcode.com/problems/number-of-even-and-odd-bits/
// AAAAAAA
// 555555


func evenOddBit(n int) []int {
    oddMask  := 0x2AAAAAAAAAAAAAAA
    evenMask := 0x5555555555555555

    oddNum := n & oddMask
    evenNum := n & evenMask

    eCount := 0
    for evenNum != 0 {
        eCount++
        evenNum = evenNum & (evenNum -1)
    }
    oCount := 0
    for oddNum != 0 {
        oCount++
        oddNum = oddNum & (oddNum -1)
    }
    return []int{eCount, oCount}
}

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
