// 67. Add Binary
// https://leetcode.com/problems/add-binary/

func addBinary(a string, b string) string {
    
    ans := ""

    aByte := []byte(a)
    bByte := []byte(b)
    
    len1 := len(aByte)
    len2 := len(bByte)
    carry := 0

    i, j := len1-1, len2 -1
    for  i >= 0 && j >= 0 {
        x := aByte[i] 
        y := bByte[j]
        s := int(x - '0') + int(y - '0') + carry
        
        carry = s / 2
        ans = ans + strconv.Itoa(s % 2)
        i, j = i-1, j-1 
    }
    for i >= 0 {
        x := aByte[i]
       
        s := int(x-'0')  + carry
        carry = s / 2
        ans = ans + strconv.Itoa(s % 2)
        i -= 1
    }
    for  j >= 0 {
        
        y := bByte[j]
        s := int(y-'0') + carry
        carry = s / 2
        ans = ans + strconv.Itoa(s % 2)
        j -= 1
    }

    if carry != 0 {
        ans = ans + strconv.Itoa(carry)
    }
    left := 0
    right := len(ans) -1
    ans1 := []byte(ans)
    for left < right {
        ans1[left], ans1[right] = ans1[right], ans1[left]
        left++
        right--
    }
    ans = string(ans1)
    fmt.Println("ans:", ans)
    return ans
}
