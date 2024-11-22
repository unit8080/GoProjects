
//  Merge k Sorted Linked Lists in O(N log k) Time with Priority Queue 
// https://leetcode.com/problems/merge-k-sorted-lists/

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */

func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val <= l2.Val {
        l1.Next = mergeTwoSortedLists(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeTwoSortedLists(l1, l2.Next)
        return l2
    }
}

func mergeKLists(lists []*ListNode) *ListNode {

    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }

    return mergeRecursive(lists, 0, len(lists)-1)
}

func mergeRecursive(lists []*ListNode, star, end int) *ListNode {
    
    if start == end {
        return lists[start]
    } 

    mid := start + (end -start)/2
    l1 := mergeRecursive(lists, start, mid)
    l2 := mergeRecursive(lists, mid+1, end)

    return mergeTwoSortedLists(l1, l2)
}
