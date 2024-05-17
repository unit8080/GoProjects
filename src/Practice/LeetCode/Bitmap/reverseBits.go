// 190. Reverse Bits
// https://leetcode.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {

    lmask := uint32 (0x00000001)
    rmask := uint32 (0x80000000)
    ans := uint32 (0)
    for i := 0; i < 32; i++ {
        if num & lmask != 0 {
            ans = ans | rmask
        }
        lmask = lmask << 1
        rmask = rmask >> 1
    }
    return ans
    // var output uint32 = 0
    // for i := 0; i < 32; i++ {
    //     bit := (num >> i ) & 0x1
    //     output |= bit << (32 - i - 1)
    // }
    // return output
}
