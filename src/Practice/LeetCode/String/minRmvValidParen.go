type stack []int

func IsEmpty(s []int) bool {
    return len(s) == 0
}

func minRemoveToMakeValid(s string) string {
    
    st := make([]int, 0)
    length := len(s)
    rbkt := make([]int, 0)
    for pos := 0; pos < length; pos++ {
        ch := s[pos]
        if ch == '(' {
            st = append(st, pos)
        }
        if ch == ')' {
            // pop
            if IsEmpty(st) {
                rbkt = append(rbkt, pos)
            } else {
                size := len(st)
                st = st[:size -1]
            }
        }
    }
    str := []byte (s) // convert to byte array
    for i := 0; i < len(st); i++ {
        index := st[i]
        str[index] = ' '
    }
    for i := 0; i < len(rbkt); i++ {
        index := rbkt[i]
        str[index] = ' '
    }
    myStr := string(str[:]) // convert back to string
    mystr := strings.Split(myStr, " ") // split string by space
    myStr = strings.Join(mystr, "") // join back without space

    return myStr
}
