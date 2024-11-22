// 1171. Remove Zero Sum Consecutive Nodes from Linked List
// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    
    dummy := new(ListNode)
    dummy.Next = head
    prefixMap := map[int] *ListNode{0:dummy}

    prefixSum := 0
    for head != nil {

        prefixSum += head.Val
        if node, exists := prefixMap[prefixSum]; !exists {
            prefixMap[prefixSum] = head

        } else {
            start := node
            temp := start.Next
            sum := prefixSum
            for temp != head {
                sum += temp.Val
                delete(prefixMap, sum)
                temp = temp.Next
            }
            start.Next = head.Next
        }
        head = head.Next
    } 
    return dummy.Next
}
