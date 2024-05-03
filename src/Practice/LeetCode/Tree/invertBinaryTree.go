// 226. Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/
// Logic: for each node, sawp the left & right subtree ptr
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    
    if root == nil { return root }

    queue := []*TreeNode{} 
    queue = append(queue, root)
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        node.Left, node.Right = node.Right, node.Left
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    return root
}
