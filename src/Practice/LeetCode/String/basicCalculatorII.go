// 227. Basic Calculator II
// https://leetcode.com/problems/basic-calculator-ii/

func calculate(s string) int {

    var ans, num int
    stack  := []int{}

    op := '+' // default first operator
    
    // append an arbitrary operator to simplify parsing
    s = s + "+" // to force the last operator to execute
    num = 0
    for _, ch := range s {
        if ch == ' ' {
            continue
        } else if ch >= '0' && ch <= '9' {
            num = num * 10 + int(ch - '0')
        } else {
            switch op { // act on previous operator
                case '+':
                    stack = append(stack, num)
                case '-':
                    stack = append(stack, -num) // negative !!!
                case '*':
                    stack[len(stack)-1] *= num // top * num !!!
                case '/':
                    stack[len(stack)-1] /= num // top / num !!!
            }
            op = ch // update operator here for next operation
            num = 0 // reset number now
        }
    }

    for _, v := range stack { // if still on stack
        ans += v
    }
    return ans
}
