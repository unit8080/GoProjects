// 270. Closest Binary Search Tree Value
// https://leetcode.com/problems/closest-binary-search-tree-value/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func solve(root *TreeNode, target float64, diff *float64, ans *int)  {
    if root == nil {
        return
    }

    tmp := math.Abs(target - float64(root.Val))
    fmt.Println("Tmp:", tmp)
    if tmp < *diff {
        *diff = tmp
        *ans = root.Val
        solve(root.Left, target, diff, ans)
    } else if tmp == *diff && *ans > root.Val  {
        *ans = root.Val
        solve(root.Right, target, diff, ans)
    }

    // solve(root.Left, target, diff, ans)
    // solve(root.Right, target, diff, ans)
}
func closestValue(root *TreeNode, target float64) int {
    
    ans := math.MaxInt
    var diff float64 = math.MaxFloat64
    
    solve(root, target, &diff, &ans)
    return ans
}
