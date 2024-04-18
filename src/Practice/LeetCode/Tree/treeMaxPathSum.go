// 124. Binary Tree Maximum Path Sum
// https://leetcode.com/problems/binary-tree-maximum-path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans := math.MinInt
    
    findSum(root, &ans)
    return ans
}

func findSum(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }
    
    leftSum := max(findSum(root.Left, ans), 0)  // why 0?
    rightSum := max(findSum(root.Right, ans), 0) // why 0?
    *ans = max(*ans, leftSum + root.Val + rightSum )
    return root.Val + max(leftSum, rightSum)
}
