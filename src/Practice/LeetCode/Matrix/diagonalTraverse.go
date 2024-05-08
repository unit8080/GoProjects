// 498. Diagonal Traverse
// https://leetcode.com/problems/diagonal-traverse/description/

func findDiagonalOrder(mat [][]int) []int {

    diagonals := make([][]int, len(mat) + len(mat[0]))
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            diagonals[i+j] = append(diagonals[i+j], mat[i][j])
        }
    }

    ans := make([]int, 0, len(mat)*len(mat[0]))
    for d, diagonal := range diagonals {
        if d % 2 == 0 {
            // reverse the diagonal
            for i, j := 0, len(diagonal) -1; i < j; i, j = i +1, j -1 {
                diagonal[i], diagonal[j] = diagonal[j], diagonal[i]
            }
        }
        ans = append(ans, diagonal...)
    }
    return ans
}
