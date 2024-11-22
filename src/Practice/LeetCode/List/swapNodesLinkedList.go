// 1721. Swapping Nodes in a Linked List
// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func swapNodes(head *ListNode, k int) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }

//     frontNode := head
//     endNode := head
//     currNode := head
//     i := 1

//     for i < k && currNode != nil {
//         currNode = currNode.Next
//         i++
//     }

//     if currNode != nil {
//         frontNode = currNode
//     }

//     for currNode.Next != nil {
//         endNode = endNode.Next
//         currNode = currNode.Next
//     }

//     if endNode != nil {
//         frontNode.Val, endNode.Val = endNode.Val, frontNode.Val
//     }
    
//     return head
// }

func swapNodes(head *ListNode, k int) *ListNode {

    var p1  *ListNode
    var p2  *ListNode

    temp := head

    for  temp != nil {
        if p2 != nil {
            p2 = p2.Next
        }
        k--
        if k == 0 {
            p1 = temp
            p2 = head
            temp = p1
        }
        temp = temp.Next
    }
    p1.Val, p2.Val = p2.Val, p1.Val
    return head

}
