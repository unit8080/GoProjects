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

//     var prev *ListNode =  nil
//     dummy := head.Next

//     curr := head

//     for curr != nil && curr.Next != nil {
//         nextNode := curr.Next.Next // Node C

//         if prev != nil {
//             prev.Next = curr.Next // A -> D in next iteration !!!!
//         }
//         prev = curr  // node A

//         curr.Next.Next = curr  // B -> A
//         curr.Next = nextNode   // A -> C, for odd #nodes

//         curr = nextNode // move to next pair
//     }
//     return dummy
// }

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    next := head.Next.Next // 3      // nil
    tmp := head            // 1      // 3
    head = head.Next       // 2      // 4
    head.Next = tmp        // 2 -> 1 // 4 -> 3
    tmp.Next = nil         // ???? - did not quite follow

    // these three lines not necessary.
    // if next == nil { // 3     
    //    return head // 4
    // }
    nextHead := swapPairs(next) // 3
    head.Next.Next = nextHead
    
    return head
}
