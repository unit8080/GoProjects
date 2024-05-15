// 670. Maximum Swap
// https://leetcode.com/problems/maximum-swap/

func maximumSwap(num int) int {
	// convert the number to byte array for easier manipulation
    s := []byte(strconv.Itoa(num))
    max, maxIdx := -1, len(s) - 1
	// i and j are used as indexes for swapping
    i, j := 0, 0
    
    for k := len(s) - 1; k >= 0; k-- {
        digit := int(rune(s[k]) - '0')
        if max < digit {
            max = digit
            maxIdx = k
        }
        if digit < max {
            // the max digit can be swapped with the current digit
			// to get a greater number
            i, j = maxIdx, k
        }
    } 
	
	// swap
    if i != j {
        s[i], s[j] = s[j], s[i]
    }
   
    // convert back to integer and return
    res, _ := strconv.Atoi(string(s))
    return res
}
