// 71. Simplify Path
// https://leetcode.com/problems/simplify-path/

type Stack []string
func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}
func (s *Stack) Push(str string) {
    (*s) = append((*s), str)
}

func (s *Stack) Pop() {
    index := len(*s) -1
    (*s) = (*s)[:index]
}

func simplifyPath(path string) string {
    
    s := &Stack{}
    
    var strs []string
    strs = strings.Split(path, "/")

    // split into array of strings
    // then process "." and ".."
    // else push onto stack
    for _, str := range strs {
        if str == "" || str == "." {
            continue
        } else if str == ".." {
            if !s.IsEmpty() {
                s.Pop()
            }
        } else {
            s.Push(str)
        }
    }

    build := ""
    if len(*s) == 0 { return "/"} // Important
    for pos := 0; pos < len(*s); pos++ {
        build += "/"
        build +=  (*s)[pos]
    }
    
    return build
}
