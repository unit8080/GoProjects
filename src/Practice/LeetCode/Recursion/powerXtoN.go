// 50. Pow(x, n)
// https://leetcode.com/problems/powx-n/

func power(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    } // devide and conquer
    ans := power(x, n/2)
    if n % 2 == 0 {
        return  ans * ans
    } else {
        return x * ans * ans
    }
}
func myPow(x float64, n int) float64 {
    numNeg := false
    powerNeg := false
    if x < 0 {
        numNeg = true
        x = x * -1
    }
    if n < 0 {
        powerNeg = true
        n = n * -1
    }

    final := power(x, n)
    if powerNeg {
        final = 1/final
    }
    if numNeg && n % 2 != 0 {
        final = final * -1
    }
    return final
}
