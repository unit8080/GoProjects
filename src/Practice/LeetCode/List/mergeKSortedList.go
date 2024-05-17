//  Merge k Sorted Linked Lists in O(N log k) Time with Priority Queue 
// https://leetcode.com/problems/merge-k-sorted-lists/

type MinHeap []*ListNode

func (h MinHeap) Len() int {
    return len(h)
}
func (h MinHeap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}
func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(node interface{}) {
    *h = append(*h, node.(*ListNode))
}
func (h *MinHeap) Pop() interface{}  {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    minHeap := &MinHeap{}

    // build the heap
    for _, list := range lists {
        if list != nil {
            heap.Push(minHeap, list)
        }
    }

    head := new(ListNode)
    prev := head

    for len(*minHeap) > 0 {
        node := heap.Pop(minHeap)
        if node.(*ListNode) == nil {
            continue
        }

        prev.Next = node.(*ListNode)
        prev = node.(*ListNode)

        if node.(*ListNode).Next != nil {
            heap.Push(minHeap, node.(*ListNode).Next)
        }
    }
    return head.Next
}
