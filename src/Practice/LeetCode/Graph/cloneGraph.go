// 133. Clone Graph
// https://leetcode.com/problems/clone-graph/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneDFS(node *Node, cloned map[*Node]*Node) {

    if _, exists := cloned[node]; exists {
        return
    }
    newNode := new(Node)
    newNode.Val = node.Val
    cloned[node] = newNode
    for _, neighbor := range node.Neighbors {
        
        cloneDFS(neighbor, cloned)
        
        newNode.Neighbors = append(newNode.Neighbors, cloned[neighbor])
    }
}
func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }

    cloned := make(map[*Node]*Node)
	
    cloneDFS(node, cloned)
    return cloned[node]
}
