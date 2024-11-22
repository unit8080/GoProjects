// 160. Intersection of Two Linked Lists
// https://leetcode.com/problems/intersection-of-two-linked-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
//     nodeMap := make(map[*ListNode]struct{})
    
//     for node := headA; node != nil; node = node.Next {
//         nodeMap[node] = struct{}{}
//     }
//     for node := headB; node != nil; node = node.Next {
//         if _, exists := nodeMap[node]; exists {
//             return node
//         }
//     }
//     return nil
// }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    if headA == headB {
        return headA
    }
    len1 := 0
    len2 := 0
    l1 := headA
    l2 := headB
    for l1 != nil {
        len1++
        l1 = l1.Next
    }
    for l2 != nil {
        len2++
        l2 = l2.Next
    }
    l1 = headA
    l2 = headB
    if len1 > len2 {
        diff := len1 - len2
        for diff > 0 {
            l1 = l1.Next
            diff--
        }
    } else {
        diff := len2 - len1
        for diff > 0 {
            l2 = l2.Next
            diff--
        }
    }
    for l1 != nil {
        if l1 == l2 {
            return l1
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    return nil
}
