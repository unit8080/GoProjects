// 987. Vertical Order Traversal of a Binary Tree
// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/

/*
*
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }

func verticalTraversal(root *TreeNode) [][]int {

}
*/
package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	node *TreeNode
	col  int
}
type LevelVal struct {
	level int
	val   int
}

func verticalTraversal(root *TreeNode) [][]int {

	ans := [][]int{}
	if root == nil {
		return ans
	}

	// key : col, val : slice of struct{level, val}
	colmap := make(map[int][]LevelVal)

	queue := make([]Node, 0)
	queue = append(queue, Node{root, 0})
	row := 0
	for len(queue) > 0 {

		qsize := len(queue)
		for qsize > 0 {
			item := queue[0]
			queue = queue[1:]

			col := item.col
			node := item.node
			colmap[col] = append(colmap[col], LevelVal{row, node.Val})

			if node.Left != nil {
				lCol := col - 1
				queue = append(queue, Node{node.Left, lCol})
			}
			if node.Right != nil {
				rCol := col + 1
				queue = append(queue, Node{node.Right, rCol})
			}
			qsize -= 1
		}
		row++
	}

	// sort by column i.e. -2, -1, 0, 1, 2, ..
	cols := make([]int, 0, len(colmap))
	for col := range colmap {
		cols = append(cols, col)
	}

	sort.Ints(cols)

	for i := 0; i < len(cols); i++ {
		col := cols[i]

		// Now Sort by level then sort by Value for same level
		colSlice := colmap[col]
		sort.SliceStable(colSlice, func(i, j int) bool {
			if colSlice[i].level != colSlice[j].level {
				return colSlice[i].level < colSlice[j].level
			}

			return colSlice[i].val < colSlice[j].val
		})
		var temp []int
		for _, levelVal := range colSlice {
			temp = append(temp, levelVal.val)
		}
		ans = append(ans, temp)
	}
	return ans
}
