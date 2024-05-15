i// 103. Binary Tree Zigzag Level Order Traversal
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {

	var ans [][]int
	var queue []*TreeNode

	if root == nil {
		return ans
	}

	queue = append(queue, root)
	altReverse := false // alternate reverse

	for len(queue) > 0 {

		count := len(queue)
		list := []int{}
		for count > 0 {

			node := queue[0]
			queue = queue[1:]

			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			count -= 1
		}
		if altReverse {
			left := 0
			right := len(list) - 1
			for left < right {
				list[left], list[right] = list[right], list[left]
				left++
				right--
			}
		}
		ans = append(ans, list)
		altReverse = !altReverse
	}
	return ans
}
