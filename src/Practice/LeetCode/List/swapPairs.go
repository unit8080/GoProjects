// 24. Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	head = curr.Next
	var prev *ListNode
	var next *ListNode
	var nextNext *ListNode

	for curr != nil {
		next = curr.Next
		if next != nil {
			if prev != nil {
				prev.Next = next
			}
            nextNext = next.Next
			next.Next = curr
            curr.Next = nil
			prev = curr
			curr = nextNext // even # of nodes termination

		} else {
			if prev != nil {
				prev.Next = curr
			}
			curr = nil // to terminate for loop for odd # of nodes
		}
	}

	return head
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    dummy := new(ListNode)
    // dummy.Next = head

    curr := head
    next := curr.Next
    // dummy.Next = next
    prev := dummy

    for next != nil {
        nextnext := next.Next

        curr.Next = nil
        next.Next = curr
        prev.Next = next

        prev = curr
        curr = nextnext
        if curr != nil {
            next = curr.Next
        } else {
            next = nil
        }
    }
    if curr != nil && next == nil {
        prev.Next = curr
    }
    return dummy.Next
}

// 24. Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func swapPairs(head *ListNode) *ListNode {

//     if head == nil || head.Next == nil {
//         return head
//     }

//     var prev *ListNode = nil
//     curr := head
//     head = head.Next

//     // A -> B -> C -> D ...
//     for curr != nil && curr.Next != nil {

//         nextNode := curr.Next.Next // Save Node C

//         if prev != nil {
//             prev.Next = curr.Next // A -> D in next iteration !!!!
//         }
//         prev = curr  // now prev points to node A

//         curr.Next.Next = curr  // B -> A
//         curr.Next = nextNode   // A -> C, for odd #nodes

//         curr = nextNode // move to Node C to swap next pair
//     }
//     return head
// }

// func swapPairs(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }

//     next := head.Next.Next // 3      // nil
//     tmp := head            // 1      // 3
//     head = head.Next       // 2      // 4
//     head.Next = tmp        // 2 -> 1 // 4 -> 3
//     tmp.Next = nil         // ???? - did not quite follow

//     // these lines are unnecessary !!!
//     // if next == nil { // 3
//     //     return head // 4
//     // }
//     nextHead := swapPairs(next) // 3
//     head.Next.Next = nextHead

//     return head
// }
