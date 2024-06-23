// 142. Linked List Cycle II
// https://leetcode.com/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    
    if head == nil || head.Next == nil {
        return nil
    }
    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }

    if slow != fast {
        return nil
    }

    p := head
    for p != slow {
        p = p.Next
        slow = slow.Next
    }
    return p
}
