// 938. Range Sum of BST
// https://leetcode.com/problems/range-sum-of-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func dfs(root *TreeNode, low int, high int, sum *int) {
     // preorder traversal
     if root == nil {
         return
     }
     // in range, add it
     if root.Val <= high && root.Val >= low {
         *sum = *sum + root.Val
     }
     if root.Val > low {
         dfs(root.Left, low, high, sum)
     }
     if root.Val < high {
         dfs(root.Right, low, high, sum)
     }
 }
func rangeSumBST(root *TreeNode, low int, high int) int {
    var sum int = 0

    dfs(root, low, high, &sum)
    return sum
}

