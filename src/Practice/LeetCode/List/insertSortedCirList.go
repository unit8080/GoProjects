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
    head := aNode

    if head == nil {
        head = newNode
        head.Next = head
        return head
    }

    prev := head
    curr := head.Next

    for curr != head {
        if prev.Val <= x && x <= curr.Val {
            newNode.Next = curr
            prev.Next = newNode
            return head
        } else if prev.Val > curr.Val && x < curr.Val {
            // 2, 3, 5 , 1 insert 0
            newNode.Next = curr
            prev.Next = newNode
            return head
        } else if prev.Val > curr.Val && x >= prev.Val {
            // []2, 3, 5, 1], insert 5 or 6
            newNode.Next = curr
            prev.Next = newNode
            return head
        }
        prev = curr
        curr = curr.Next
    }
    newNode.Next = curr
    prev.Next = newNode
    return head
}

// case 1: [3,4,1], 2 >> [3,4,1,2]
// case 2: [], 1 >> [1] 
// case 3: [1], 0 >> [1,0]
// case 4: [3,5,1], 0 >> [3,5,0,1]
// case 5: [3,5,1], 5 >> [3,5,5,1]
// case 6: [3,5,1], 6 >> [3,5,6,1]
// case 7: [1,3,3], 4 >> [1,3,3,4]
