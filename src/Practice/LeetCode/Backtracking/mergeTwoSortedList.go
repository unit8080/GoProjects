// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
//     var head *ListNode = nil

//     curr1 := list1
//     curr2 := list2

//     var prev *ListNode = nil

//     for curr1 != nil && curr2 != nil {
//         if curr1.Val <= curr2.Val {
//             // delink node from list1
//             if prev != nil {
//                 prev.Next = curr1
//             }
//             prev = curr1
//             curr1 = curr1.Next
//         } else {
//             // delink node from list2
//             if prev != nil {
//                 prev.Next = curr2
//             }
//             prev = curr2
//             curr2 = curr2.Next
//         }
//         if head == nil {
//             head = prev
//         }
//     }

//     if curr1 != nil {
//         if prev != nil {
//             prev.Next = curr1
//         } else {
//             head = curr1
//         }
//     }
//     if curr2 != nil {
//         if prev != nil {
//             prev.Next = curr2
//         } else {
//             head = curr2
//         }
//     }
    
//     return head
// }

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

//     dummyHead := &ListNode{-1, nil}
//     prev := dummyHead

//     for l1 != nil && l2 != nil {
//         if l1.Val < l2.Val {
//             prev.Next = l1
//             prev = l1
//             l1 = l1.Next
//         } else {
//             prev.Next = l2
//             prev = l2
//             l2 = l2.Next
//         }
//     }
//     if l1 != nil {
//         prev.Next = l1
//     }
//     if l2 != nil {
//         prev.Next = l2
//     }
//     return dummyHead.Next
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyNode := &ListNode{-1, nil}
    prev := dummyNode

    var solve func(l1 *ListNode, l2 *ListNode)
    solve = func(l1 *ListNode, l2 *ListNode) {
        if l1 == nil {
            prev.Next = l2
            return
        } else if l2 == nil {
            prev.Next = l1
            return
        } else if l1.Val < l2.Val {
            prev.Next = l1
            prev = l1
            l1 = l1.Next
        } else {
            prev.Next = l2
            prev = l2
            l2 = l2.Next
        }
        solve(l1, l2)
    }
    solve(l1, l2)
    return dummyNode.Next
}
