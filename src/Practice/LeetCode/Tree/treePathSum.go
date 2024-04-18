// 112. Path Sum
// https://leetcode.com/problems/path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
 func hasSum (root *TreeNode, sum, targetSum int ) bool {
    if root == nil { return false }
    fmt.Println("sum:", sum)
    
    sum += root.Val
    if root.Left == nil && root.Right == nil {
        return  sum == targetSum 
    }
    fmt.Println("sum:", sum)

    return hasSum(root.Left, sum, targetSum) || 
            hasSum(root.Right, sum, targetSum)
 }
func hasPathSum(root *TreeNode, targetSum int) bool {
    return hasSum(root, 0, targetSum)
}
*/

/* Another way:

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    targetSum -= root.Val
    if root.Left == nil && root.Right == nil {
            return targetSum == 0
    }
    return hasPathSum(root.Left, targetSum) || 
            hasPathSum(root.Right, targetSum)
}

*/

func hasPathSum(root *TreeNode, targetSum int) bool {
    
    if root == nil { return false }
    
    if root.Left == nil && root.Right == nil {
        targetSum -= root.Val
        if targetSum == 0 {
            return true
        } else {
            return false
        }
    }
    return hasPathSum(root.Left, targetSum - root.Val) || 
            hasPathSum(root.Right, targetSum - root.Val)
}
