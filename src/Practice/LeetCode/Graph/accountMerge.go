// 721. Accounts Merge
// https://leetcode.com/problems/accounts-merge/
// https://leetcode.com/problems/accounts-merge/solutions/4891520/golang-simple-solution-using-dfs-and-map-in-o-n2/
// Video: https://www.youtube.com/watch?v=f17PKE8W2p8

func accountsMerge(accounts [][]string) [][]string {

    // build Adj graph
    relationMap := make(map[string][]string)
    visited := make(map[string]bool)

    for _, account := range accounts {
        firstEmail := account[1]
        for j := 2; j < len(account); j++ {
            relationMap[firstEmail] = append(relationMap[firstEmail], account[j])
            relationMap[account[j]] = append(relationMap[account[j]], firstEmail)
        }
    }

    res := make([][]string, 0)
    for _, emails := range accounts {
        name := emails[0]
        email := emails[1]
        if !visited[email] {
            ans := []string{name}
            DFS(&ans, email, relationMap, visited)
            sort.Strings(ans[1:])
            res = append(res, ans)
        }
    }
    return res
}

func DFS(ans *[]string, email string, relationMap map[string][]string, visited map[string]bool) {
    if visited[email] {
        return
    }

    visited[email] = true
    *ans = append(*ans, email)
    for _, adEmail := range relationMap[email] {
        if !visited[adEmail] {
            DFS(ans, adEmail, relationMap, visited)
        }
    }
}
