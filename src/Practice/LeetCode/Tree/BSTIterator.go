// 173. Binary Search Tree Iterator
// https://leetcode.com/problems/binary-search-tree-iterator/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    stack := make([]*TreeNode, 0)
    for root != nil {
        stack = append(stack, root)
        root = root.Left
    }
    return BSTIterator { stack: stack }
}


func (this *BSTIterator) Next() int {
    n := len(this.stack)
    top := this.stack[n-1]
    this.stack = this.stack[:n-1]
    node := top.Right
    for node != nil {
        this.stack = append(this.stack, node)
        node = node.Left
    }
    return top.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.stack) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
