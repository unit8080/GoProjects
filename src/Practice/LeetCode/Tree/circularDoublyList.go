// 426. Convert Binary Search Tree to Sorted Doubly Linked List
// https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */
var first *Node = nil
var last  *Node = nil

func treeToDoublyList(root *Node) *Node {
    first = nil
    last  = nil
    if root == nil {
        return root
    }
    
    convert(root)
    last.Right = first // Circular doubly list
    first.Left = last  // Circular doubly list
    return first
}
func convert(node *Node) {

    if node == nil { return }

    convert(node.Left)
    
    // do with node
    if last != nil {
        last.Right = node
        node.Left = last
    } else {
        first = node
    }
    last = node

    convert(node.Right)
}
