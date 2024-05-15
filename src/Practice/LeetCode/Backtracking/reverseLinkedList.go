// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil { return head }
    
//     var prev *ListNode = nil
//     curr := head
    
//     for curr != nil {
//         next := curr.Next
//         curr.Next = prev
//         prev = curr
//         curr = next
//     }
//     return prev
// }

func solve(head, prev *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // 1 -> 2 -> nil
    next := head.Next           // next = 2, prev = nil
    head.Next = prev            // 1-> nil
    prev = head                 // 1
    head = solve(next, prev)    // (2, 1)
    next.Next = prev            // backtracking ...
    return head                 // 2
}
func reverseList(head *ListNode) *ListNode {
    return solve(head, nil)
}

