// 200. Number of Islands
// https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
    count := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                count++
                bfs(grid, i, j)
            }
        }
    }
    return count
}

func bfs(grid [][]byte, r, c int) {

    queue := [][2]int{}
    queue = append(queue, [2]int{r, c})
    grid[r][c] = '2'

    offsets := [4][2]int{{-1, 0}, {1,0}, {0, 1}, {0, -1}}
    for len(queue) > 0 {
        cell := queue[0]
        queue = queue[1:]
        
        for _, offset := range offsets {
            row := cell[0] + offset[0]
            col := cell[1] + offset[1]
            if row >= 0 && row < len(grid) && col >= 0  && col < len(grid[row]) && grid[row][col] == '1' {
                queue = append(queue, [2]int{row, col})
                grid[row][col]= '2'
            } 
        }
    }
}
