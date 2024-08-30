// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

type Stack []byte

func isValid(s string) bool {
    if s == "" {return true}

    var stack  Stack
    for pos := 0; pos < len(s); pos++ {
        
        ch := (s[pos])
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)
        } else if ch == ')' || ch == '}' || ch == ']' {
            // if stack empty
            if len(stack) == 0 { return false}
            index := len(stack) -1
            top := stack[index]
            if top == '(' && ch != ')' || 
                top == '{' && ch != '}' ||
                top == '[' && ch != ']' {
                    return false
            }
            stack = (stack)[:index]
        }
    }
    if len(stack) != 0 { return false}
    return true
}
