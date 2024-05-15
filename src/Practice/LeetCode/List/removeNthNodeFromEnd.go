// 19. Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil { return head }
    i := 0
    fwd := head
    back := head
    for i < n && fwd != nil {
        // fwd node n-times
        fwd = fwd.Next
        i += 1
    }
    var prev *ListNode = nil
    for fwd != nil {
        prev = back
        back = back.Next
        fwd = fwd.Next
    }
    if prev != nil {
        prev.Next = back.Next
        return head
    } else {
        return back.Next
    }
}
