// 1404. Number of Steps to Reduce a Number in Binary Representation to One
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/

func numSteps(s string) int {
    
    operations := 0
    buf := []byte(s)

    for len(buf) != 1 {
        n := len(buf)
        if buf[n-1] == '0' {
            buf = buf[:n-1]
        } else {
            // add 1
            overflow := true
            for i := len(buf) -1; i >= 0; i-- {
                if buf[i] == '1' {
                    buf[i] = '0'
                } else {
                    buf[i] = '1'
                    overflow = false
                    break
                }
            }
            if overflow {
                buf = append([]byte{'1'}, buf...)
            }
        }
        operations++
    }
    return operations
}
