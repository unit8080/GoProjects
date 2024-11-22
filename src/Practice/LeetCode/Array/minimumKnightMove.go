// 1197. Minimum Knight Moves
// https://leetcode.com/problems/minimum-knight-moves/?envType=company&envId=nvidia&favoriteSlug=nvidia-thirty-days
// https://www.youtube.com/watch?v=NtwSlfOpEMg

// func minKnightMoves(x int, y int) int {
    
// }

func minKnightMoves(x int, y int) int {

   direction := [][]int{{2,1}, {1,2}, {-1, 2}, {-2, 1},{-2,-1}, {-1,-2}, {1, -2}, {2, -1}} 
   queue := [][]int{{0,0,0}}

   visited := make(map[[2]int]bool)
   visited[[2]int{0,0}] = true
   for len(queue) > 0 {
    q := queue[0]
    queue = queue[1:]
    for _, dir := range direction {
        if x == q[0] && y == q[1] {
            return q[2]
        }
        row := q[0] + dir[0]
        col := q[1] + dir[1]
        step := q[2]
        point := [2]int{row, col}
        if !visited[point] {
            visited[point] = true
            queue = append(queue, []int{row, col, step+1})
        }
    }
   }
   return -1 
}
