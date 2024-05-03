// h863. All Nodes Distance K in Binary Tree
// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {

    Adj := make(map[int][]int)
    var buildGraph func(curr *TreeNode, parent *TreeNode)
    buildGraph = func(curr *TreeNode, parent *TreeNode) {
        if curr == nil {
            return
        }
        if curr != nil && parent != nil {
            Adj[curr.Val] = append(Adj[curr.Val], parent.Val)
            Adj[parent.Val] = append(Adj[parent.Val], curr.Val)
        }
        if curr.Left != nil {
            buildGraph(curr.Left, curr)
        }
        if root.Right != nil {
            buildGraph(curr.Right, curr)
        }
    }
    buildGraph(root, nil)
    ans := []int{}

    visited := make(map[int]struct{})
    visited[target.Val] = struct{}{}
    var dfs func(val int, distance int)
    dfs = func(val int, distance int) {
        if distance == k {
            ans = append(ans, val)
        }
        for _, neighbor := range Adj[val] {
            _, ok := visited[neighbor]
            if !ok {
                visited[neighbor] = struct{}{}
                dfs(neighbor, distance+1)
            }
        }
    }
    dfs(target.Val, 0)
    return ans
}
