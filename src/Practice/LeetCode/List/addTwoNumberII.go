// Implement using stacks (w/0 linked list reversal)


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse( l *ListNode) *ListNode {
    
    curr := l
    var prev *ListNode = nil
    
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1 = reverse(l1)
    l2 = reverse(l2)

    ans := new (ListNode)
    carry := 0
    sum := 0

    for l1 != nil || l2 != nil {

        sum = ans.Val
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        ans.Val = sum % 10
        carry = sum / 10

        node := new (ListNode)
        node.Val = carry
        node.Next = ans
        ans = node
    }
    if ans.Val == 0 {
        return ans.Next
    }

    return ans
}
