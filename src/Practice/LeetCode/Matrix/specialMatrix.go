// 1582. Special Positions in a Binary Matrix
// https://leetcode.com/problems/special-positions-in-a-binary-matrix/

// (m * n)(m + n)
// func numSpecial(mat [][]int) int {
// 	m := len(mat)
// 	n := len(mat[0])

// 	result := 0
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if mat[i][j] == 1 {
// 				special := true
// 				for x := 0; x < m; x++ {
// 					if x != i && mat[x][j] == 1 {
// 						special = false
// 						break
// 					}
// 				}
// 				if special == false {
// 					continue
// 				}
// 				for y := 0; y < n; y++ {
// 					if y != j && mat[i][y] == 1 {
// 						special = false
// 						break
// 					}
// 				}
// 				if special == true {
// 					result++
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

func numSpecial(mat [][]int) int {
    m := len(mat) // number of rows
    n := len(mat[0]) // number of columns

    rowArr := make([]int, m)
    colArr := make([]int, n)

    // fill row & col Arr
    for row := 0; row < m; row++ {
        for col := 0; col < n; col++ {
            if mat[row][col] == 1 {
                rowArr[row]++
                colArr[col]++
            }
        }
    }
    result := 0
    for row := 0; row < m; row++ {
        for col := 0; col < n; col++ {
            if mat[row][col] == 1 && rowArr[row] == 1 && colArr[col] == 1 {
                result++
            }
        }
    }
    return result
}
