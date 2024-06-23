// 2487. Remove Nodes From Linked List
// https://leetcode.com/problems/remove-nodes-from-linked-list/

// You are given the head of a linked list.

// Remove every node which has a node with a greater value anywhere to the right side of it.

// Return the head of the modified linked list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNodes(head *ListNode) *ListNode {


    var reverse func(head *ListNode) *ListNode
    reverse = func(head *ListNode) *ListNode {
        var prev *ListNode = nil
        curr := head
        for curr != nil {
            next := curr.Next
            curr.Next = prev
            prev = curr
            curr = next
        }
        return prev
    }
    head = reverse(head)
    dummy := new(ListNode)
    dummy.Next = head
    curr := head
    curr_max := curr.Val
    for curr.Next != nil {
        if curr.Next.Val < curr_max {
            curr.Next = curr.Next.Next
        } else {
            curr_max = curr.Next.Val
            curr = curr.Next
        }
    }
    head = dummy.Next
    return reverse(head)
}
