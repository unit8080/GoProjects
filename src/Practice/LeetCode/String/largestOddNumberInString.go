// 1903. Largest Odd Number in String
// https://leetcode.com/problems/largest-odd-number-in-string/

func largestOddNumber(num string) string {
    n := len(num)

    for i := n-1; i >=0; i-- {
        digit := int (num[i] - byte('0'))
        if digit % 2 != 0 { // odd number 
            return num[:i+1]
        }
    }
    return ""
}
