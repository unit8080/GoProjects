// 129. Sum Root to Leaf Numbers
// https://leetcode.com/problems/sum-root-to-leaf-numbers/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func solve(root *TreeNode, num int, total *int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        num = num *10 + root.Val
        *total = *total + num
    }
    solve(root.Left, num *10 + root.Val, total)
    solve(root.Right, num *10 + root.Val, total)
}
func sumNumbers(root *TreeNode) int {
    total := 0
    solve(root, 0, &total)
    return total
}
