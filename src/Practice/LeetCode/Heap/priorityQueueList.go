package main

import "fmt"

type Node struct {
	Data     int
	Priority int
	Next     *Node
}
type PQ struct {
	Head *Node
}

func newNode(data int, prior int) *Node {
	node := new(Node)
	node.Data = data
	node.Priority = prior
	return node
}
func (pq *PQ) IsEmpty() bool {
	return pq.Head.Next == nil
}
func (pq *PQ) Peek() int {
	if pq.Head.Next != nil {
		return pq.Head.Next.Data
	} else {
		return -1
	}
}
func (pq *PQ) Pop() {
	if pq.Head.Next != nil {
		pq.Head.Next = pq.Head.Next.Next
	}
}
func (pq *PQ) Push(data int, prior int) {
	newN := newNode(data, prior)

	curr := pq.Head.Next
	if curr == nil || curr.Priority < prior {
		newN.Next = pq.Head.Next
		pq.Head.Next = newN
		return
	}
	for curr.Next != nil && curr.Next.Priority > prior {
		curr = curr.Next
	}
	newN.Next = curr.Next
	curr.Next = newN
	return
}
func main() {
	pq := &PQ{}

	pq.Head = new(Node)
	pq.Push(10, 10)
	pq.Push(5, 5)
	pq.Push(2, 2)
	pq.Push(15, 15)
	pq.Pop()
	pq.Pop()
	pq.Pop()
	pq.Pop()
	fmt.Println("Peek:", pq.Peek())
}
