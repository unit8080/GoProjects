// 1650. Lowest Common Ancestor of a Binary Tree III
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/

/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
    
    if p == nil || q == nil {
        return nil
    }
    nMap := make(map[*Node]struct{})

    n1 := p
    for n1 != nil {
        nMap[n1] = struct{}{}
        n1 = n1.Parent
    }

    n2 := q
    for n2 != nil {
        _, exists := nMap[n2]
        if exists { return n2 }
        n2 = n2.Parent
    }
    return nil
}
