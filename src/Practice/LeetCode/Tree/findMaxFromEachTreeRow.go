// 515. Find Largest Value in Each Tree Row
// https://leetcode.com/problems/find-largest-value-in-each-tree-row/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    var ans []int
    var queue []*TreeNode

    ans = make([]int, 0)
    queue = make([]*TreeNode, 0)

    if root == nil { return ans }

    queue = append(queue, root)
    count := len(queue)

    for count > 0 {
        queueLen := len(queue)
        currMax := math.MinInt
        for queueLen > 0 {
            node := queue[0]
            queue = queue[1:]

            if node.Val > currMax {
                currMax = node.Val
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            queueLen -= 1
        }
        ans = append(ans, currMax)
        count = len(queue)
    }
    return ans
}
