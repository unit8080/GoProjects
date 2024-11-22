// 547. Number of Provinces
// https://leetcode.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) int {
    
    adjList := make(map[int][]int, )
    n := len(isConnected)

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i != j && isConnected[i][j] == 1 {
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)
            }
        }
    }
    visited := make(map[int]bool, n)

    count := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            DFS(adjList, i, &visited)
            count++
        }
        
    }
    return count
}

func DFS(adjList map[int][]int, u int, visited * map[int]bool) {
    (*visited)[u] = true
    cities, exists := adjList[u]
    if !exists {
        return
    }
    for _, city := range cities {
        if !(*visited)[city] {
            DFS(adjList, city, visited)
        }
    }
}
