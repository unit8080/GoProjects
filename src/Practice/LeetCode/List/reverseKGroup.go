// 25. Reverse Nodes in k-Group
// https://leetcode.com/problems/reverse-nodes-in-k-group/description/?envType=company&envId=microsoft&favoriteSlug=microsoft-thirty-days


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode, k int) *ListNode {

	var rev_head *ListNode = nil
	ptr := head
	for k > 0 {
		nextNode := ptr.Next
		ptr.Next = rev_head
		rev_head = ptr
		ptr = nextNode
		k--
	}
	return rev_head
}
func reverseKGroup(head *ListNode, k int) *ListNode {

	ptr := head
	var kTail *ListNode = nil
	var newHead *ListNode = nil
	var rev_head *ListNode = nil
	for ptr != nil {
		count := 0
		head = ptr
		for count < k && ptr != nil {
			ptr = ptr.Next
			count++
		}
		if count == k {
			rev_head = reverseList(head, k)
		} else {
			rev_head = head
		}
		if newHead == nil {
			newHead = rev_head
		}
		if kTail != nil {
			kTail.Next = rev_head
		}
		kTail = head
	}

	if newHead == nil {
		return head
	} else {
		return newHead
	}
}
