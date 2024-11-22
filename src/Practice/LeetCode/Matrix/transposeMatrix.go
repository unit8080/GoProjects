// 867. Transpose Matrix
// https://leetcode.com/problems/transpose-matrix/

func transpose(matrix [][]int) [][]int {
    m := len(matrix)
    n := len(matrix[0])

    mat := make([][]int, n)
    for i := 0; i < n; i++ {
        mat[i] = make([]int, m)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
            mat[j][i] = matrix[i][j]
        }
    }
    return mat
}
