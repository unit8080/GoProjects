// 662. Maximum Width of Binary Tree
// https://leetcode.com/problems/maximum-width-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type QueueNode struct {
    node *TreeNode
    num int
    level int
}
func max (a int, b int) int {
    if a > b { return a };
    return b;
} 
func widthOfBinaryTree(root *TreeNode) int {
    
    queue := []QueueNode {}
    queue = append(queue, QueueNode{
            node: root,
            num: 1,
            level: 0,
    })

    lastLevel, lastNum := 0, 1
    ans := 0
    for len(queue) > 0 {
        q := queue[0]
        queue = queue[1:]
        node := q.node
        level := q.level
        num := q.num

        if level > lastLevel {
            lastLevel = level
            lastNum = num
        }
        ans = max(ans, num - lastNum +1)
        if node.Left != nil {
            queue = append(queue, QueueNode{
                    node: node.Left,
                    level: level + 1,
                    num: num * 2,
            })
        }
        if node.Right != nil {
            queue = append(queue, QueueNode{
                    node: node.Right,
                    level: level + 1,
                    num: num * 2 + 1,
            })
        }
    }

    return ans
}
