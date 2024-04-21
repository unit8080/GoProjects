// 708. Insert into a Sorted Circular Linked List
// https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
    newNode := new(Node)
    newNode.Val = x
    if head == nil {
        head = newNode
        head.Next = head
        return head
    }
    head := aNode
    node := head
    for node.Next != head {
        next := node.Next
        if node.Val <= x && x <= next.Val {
            node.Next = newNode
            newNode.Next = next
            return head
        } else if node.Val > x && next.Val < node.Val && x < next.Val {

            node.Next = newNode
            newNode.Next = next
            return head
        } else if node.Val <= x && node.Val > next.Val && x >= next.Val {
            node.Next = newNode
            newNode.Next = next
            return head
        }
        node = node.Next
    }

    newNode.Next = node.Next
    node.Next = newNode
    return head
}

// case 1: [3,4,1], 2 >> [3,4,1,2]
// case 2: [], 1 >> [1] 
// case 3: [1], 0 >> [1,0]
// case 4: [3,5,1], 0 >> [3,5,0,1]
// case 5: [3,5,1], 5 >> [3,5,5,1]
// case 6: [3,5,1], 6 >> [3,5,6,1]
// case 7: [1,3,3], 4 >> [1,3,3,4]
