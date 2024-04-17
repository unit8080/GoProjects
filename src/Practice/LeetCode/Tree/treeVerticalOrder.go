/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Node struct {
    node *TreeNode
    col int
 }
func verticalOrder(root *TreeNode) [][]int {
    
    ans := [][]int{}
    if root == nil { return ans }

    // key : col, val : slice of node Val
    colmap := make(map[int][]int)
    queue := make([]Node, 0)
    queue = append(queue, Node{root, 0})

    for len(queue) > 0 {

        qsize := len(queue)
        for qsize > 0 {
            item := queue[0]
            queue = queue[1:]

            col := item.col
            node := item.node
            val, exists := colmap[col]
            if exists {
                val = append(val, node.Val)
            } else {
                val = []int{node.Val}
            }
            colmap[col] = val
            if node.Left != nil {
                lCol := col - 1
                queue = append(queue, Node{node.Left, lCol})
            }
            if node.Right != nil {
                rCol := col + 1
                queue = append(queue, Node{node.Right, rCol})
            }
            qsize -= 1
        }
    }

    keys := make([]int, 0, len(colmap))

    for key := range colmap {
        keys = append(keys, key)
    }
    sort.Ints(keys)

    for _, key := range keys {
        ans = append(ans, colmap[key])
    }
    return ans
}
