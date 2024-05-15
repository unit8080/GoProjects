// 1047. Remove All Adjacent Duplicates In String
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

type Stack []string

func Top(stack Stack) string {
	Len := len(stack)
	if Len == 0 {
		return ""
	}
	return stack[Len-1]
}
func removeDuplicates(s string) string {
	var stack Stack
	if s == "" {
		return s
	}
	stack = append(stack, string(s[0]))
	for pos := 1; pos < len(s); pos++ {
		if Top(stack) == string(s[pos]) { // pop duplicate
			Len := len(stack) 
			stack = stack[:Len-1]
		} else {
			stack = append(stack, string(s[pos]))
		}
	}
	return strings.Join(stack, "")
}
