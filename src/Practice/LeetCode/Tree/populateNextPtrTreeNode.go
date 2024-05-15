// 116. Populating Next Right Pointers in Each Node
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    
    node := root
    var queue []*Node
    
    queue = append(queue, node)
    for len(queue) > 0 {
        count := len(queue)
        for count > 0 {
            currNode := queue[0]
            queue = queue[1:]
            if count > 1 {
                currNode.Next = queue[0]
            }
            if currNode.Left != nil {
                queue = append(queue, currNode.Left)
            }
            if currNode.Right != nil {
                queue = append(queue, currNode.Right)
            }
            count--
        }
    }
    return root
}

// func connect(root *Node) *Node {
//     if root == nil {
//         return root
//     }

//     leftmost := root
//     for leftmost.Left != nil {

//         head := leftmost
//         for head != nil {
//             // connection 1
//             head.Left.Next = head.Right

//             // connection 2
//             if head.Next != nil {
//                 head.Right.Next = head.Next.Left
//             }
//             head = head.Next
//         }
//         leftmost = leftmost.Left
//     }
//     return root
// }
