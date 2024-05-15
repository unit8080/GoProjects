// 66. Plus One
// https://leetcode.com/problems/plus-one/

func plusOne(digits []int) []int {

    carry := 1
    n := len(digits)
    if n == 0 {
        return []int{0}
    }
    
    temp := make([]int, n)
    for i := n-1; i >= 0; i-- {
        curr := digits[i]
        sum := curr + carry
        num := sum % 10
        carry = sum / 10
        temp[i] = num
    }

    if carry > 0 {
        ans := make([]int, n +1)
        ans[0] = carry
        for i := 0; i < n; i++ {
            ans[i+1] = temp[i]
        }
        return ans
    } else {
        return temp
    }
    
}
