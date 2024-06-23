// 461. Hamming Distance
// https://leetcode.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {
    xor := x ^ y
    distance := 0
    for xor > 0 {
        distance++
        xor = xor & (xor -1)
    }
    return distance
}
