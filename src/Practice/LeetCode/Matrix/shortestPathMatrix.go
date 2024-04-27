// 1091. Shortest Path in Binary Matrix
// https://leetcode.com/problems/shortest-path-in-binary-matrix/

func shortestPathBinaryMatrix(grid [][]int) int {
    // m x n grid
    m := len(grid)
    n := len(grid[0])
    if grid[0][0] != 0 || grid[m-1][n-1] != 0 {
        // invalid start or end cell
        return -1
    }

    // construct 2-D matrix
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    queue := [][]int{{0,0}} // enqueue start cell
    grid[0][0] = 1 // set distance from start

    for len(queue) > 0 {

        // dequeue
        cell := queue[0]
        queue = queue[1:]

        row := cell[0]
        col := cell[1]
        distance := grid[row][col]
        if row == m-1 && col == n-1 { // Did reach target cell
            return distance
        }
        neighbors := getNeighbors(row, col, grid)
        for i := 0; i < len(neighbors); i++ {
            nRow := neighbors[i][0]
            nCol := neighbors[i][1]
            queue = append(queue, []int{nRow, nCol})
            grid[nRow][nCol] = distance + 1
        }
    }
    return -1 // target was unreachable
}
var dir = [][]int {{-1, -1}, {-1, 0}, {-1,1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
func getNeighbors(row int, col int, grid [][]int) [][]int {

    list := [][]int{}
    for i :=0; i < len(dir); i++ {
        newRow := row + dir[i][0]
        newCol := col + dir[i][1]
        if newRow < 0 || newCol < 0 || newRow >= len(grid) ||
            newCol >= len(grid[0]) || grid[newRow][newCol] != 0 {
                continue
            } 
            list = append(list, []int{newRow, newCol})
    }
    return list
}
