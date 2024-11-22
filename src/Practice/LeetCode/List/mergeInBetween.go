// 1669. Merge In Between Linked Lists
// https://leetcode.com/problems/merge-in-between-linked-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

    var prev *ListNode = nil
    var last *ListNode = nil

    i := 0
    prev = list1
    for i < a-1 {
        prev = prev.Next
        i++
    }

    j := 0
    last = list1
    for j < b+1 {
        last = last.Next
        j++
    }

    prev.Next = list2
    next := list2
    for next.Next != nil {
        next = next.Next
    }
    next.Next = last
    return list1
}
