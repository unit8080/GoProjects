// 101. Symmetric Tree
// https://leetcode.com/problems/symmetric-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isMirror(r1 *TreeNode, r2 *TreeNode) bool {
     if r1 == nil && r2 == nil {
         return true
     }
     if r1 == nil || r2 == nil {
         return false
     }
     if r1.Val != r2.Val {
         return false
     }
     return isMirror(r1.Right, r2.Left) && isMirror(r1.Left, r2.Right)
 }
func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root)
}
