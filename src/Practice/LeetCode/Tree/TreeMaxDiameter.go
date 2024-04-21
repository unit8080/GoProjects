// 543. Diameter of Binary Tree
// https://leetcode.com/problems/diameter-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func longestPath(root *TreeNode, diameter *int) int {
     if root == nil { 
         return 0
    }
    left_path := longestPath(root.Left, diameter)
    right_path := longestPath(root.Right, diameter)

    *diameter = max(*diameter, left_path + right_path)
    return max(left_path, right_path) + 1
 }

func diameterOfBinaryTree(root *TreeNode) int {
    var dia int = 0

    longestPath(root, &dia)
    return dia

}
