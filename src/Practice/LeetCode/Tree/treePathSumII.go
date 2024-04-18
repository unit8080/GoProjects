// 113. Path Sum II
// https://leetcode.com/problems/path-sum-ii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func solve(root *TreeNode, targetSum int, slice *[]int, ans *[][]int) {

    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil {
        tmp := targetSum - root.Val
        if tmp == 0 {
            *slice = append(*slice, root.Val)
            tmpSlice := make([]int, len(*slice)) // ?? tmpSlice
            copy(tmpSlice, *slice)
            *ans = append(*ans, tmpSlice)
            // fmt.Println("Slice:", *slice)
            // fmt.Println("Ans:", *ans)
            index := len(*slice) - 1
            *slice = (*slice)[:index]
        }
        
    } else {
        targetSum = targetSum - root.Val
        *slice = append(*slice, root.Val)
    
        solve(root.Left, targetSum, slice, ans)
        solve(root.Right, targetSum, slice, ans)

        index := len(*slice) - 1
        *slice = (*slice)[:index]
        targetSum = targetSum + root.Val
    }
}

func pathSum(root *TreeNode, targetSum int) [][]int {
    var ans =  [][]int{}
    var slice =  []int{}

    solve(root, targetSum, &slice, &ans)
    return ans
}
