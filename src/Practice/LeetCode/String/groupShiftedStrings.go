// 249. Group Shifted Strings
// https://leetcode.com/problems/group-shifted-strings/

func groupStrings(strs []string) [][]string {
    
    ans := [][]string{}
    groupMap := make(map[string][]string)

    for _, str := range strs {
        
        sb := strings.Builder{}
        
        // build hash
        for i := 0; i < len(str) -1; i++ {
            diff := (len(str) + 26 + int(str[i]) - int(str[i+1]) ) % 26
            sb.WriteString(fmt.Sprintf("%d", diff))
        }
        hash := sb.String()
        groupMap[hash] = append(groupMap[hash], str)
    }

    for _, v := range groupMap {
        ans = append(ans, v)
    }
    return ans
}
