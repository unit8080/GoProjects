// 118. Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/

// func generate(numRows int) [][]int {
//     ans := make([][]int, numRows)

//     ans[0] = append(ans[0], 1)
//     for rowNum := 1; rowNum < numRows; rowNum++ {
//         prevRow := ans[rowNum-1]

//         currRow := []int{1} // first row element - 1
//         for j := 1; j < rowNum; j++ {
//             currRow = append(currRow, prevRow[j-1] + prevRow[j])
//         }
//         currRow = append(currRow, 1) // last row element - 1
//         ans[rowNum] = currRow
//     }

//     return ans
// }

// Recursion Method
func generate(numRows int) [][]int {
    ans := make([][]int, numRows)
    ans[0] = append(ans[0], 1)
    
    // solve(numsRows, &ans)
    var solve func(rowNum int)
    solve = func(rowNum int) {
        if rowNum >= numRows {
            return
        }
        currRow := []int{1}
        prevRow := ans[rowNum -1]
        for j := 1; j < rowNum; j++ {
            currRow = append(currRow, prevRow[j-1]+prevRow[j])
        }
        
        currRow = append(currRow, 1)
        ans[rowNum] = currRow
        solve(rowNum+1)
    }
    solve(1)
    return ans
}
