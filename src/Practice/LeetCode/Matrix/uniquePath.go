// 62. Unique Paths
// https://leetcode.com/problems/unique-paths/

// It needs to revisit with DP or/and Memoization 

func uniquePaths(m int, n int) int {
    
    maze := make([][]int, m)
    for i := 0; i < m; i++ {
        maze[i] = make([]int, n)
    }

    result := 0
    solve(maze, 0, 0, m , n, &result)
    return result
}

func solve(maze [][]int, i, j, m, n int, result *int) {
    if i < 0 || i >= m || j < 0 || j >= n || maze[i][j] == 1 {
        return
    }

    if i == m-1 && j == n-1 {
        (*result)++
        return
    }
    // i+ or j+
    maze[i][j] = 1
    solve(maze, i+1, j, m, n, result)

    solve(maze, i, j+1, m, n, result)
    maze[i][j] = 0
}
