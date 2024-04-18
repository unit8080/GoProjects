// 236. Lowest Common Ancestor of a Binary Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val == p.Val || root.Val == q.Val {
        return root
    }

    leftAns := lowestCommonAncestor(root.Left, p, q )
    rightAns := lowestCommonAncestor(root.Right, p, q )
    
    if leftAns != nil && rightAns != nil {
        return root
    } else if leftAns != nil && rightAns == nil {
        return leftAns
    } else if leftAns == nil && rightAns != nil {
        return rightAns
    } else {
        return nil
    }
}
