// 98. Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func validateBST(node *TreeNode, min, max int) bool {
     // base case
     if node == nil {
         return true
     }
     if node.Val <= min || node.Val >= max {
         return false
     }

     return validateBST(node.Left, min, node.Val) &&
            validateBST(node.Right, node.Val, max)
 }
func isValidBST(root *TreeNode) bool {
    return validateBST(root, math.MinInt, math.MaxInt)
}
