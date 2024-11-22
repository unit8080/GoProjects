// 86. Partition List
// https://leetcode.com/problems/partition-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    
    small := new (ListNode)
    large := new(ListNode)

    smallP := small
    largeP := large

    for head != nil {
        if head.Val < x {
            smallP.Next = head
            smallP = smallP.Next
        } else {
            largeP.Next = head
            largeP = largeP.Next
        }
        head = head.Next
    }
    largeP.Next = nil
    smallP.Next = large.Next
    return small.Next
}
