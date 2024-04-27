package main

import (
	"fmt"
	"sync/atomic"
)

// Node Struct: This is a basic element of the queue. Each Node
// holds two pieces of information:

// value: Stored using the type interface{}, which means it can
// hold any value, making the queue generic. This allows the
// queue to store any data type, providing flexibility.

// next: This is a pointer to the next node in the queue,
// and it uses Go’s atomic.Pointer[Node] type introduced
// in Go 1.19. The use of atomic.Pointer[Node] ensures that
// updates to the pointer (like changing the next node it points
// to) can be made atomically, avoiding race conditions without
// the need for mutexes.

type Node struct {
	value interface{}
	next  atomic.Pointer[Node]
}

// Queue Struct: This represents the queue itself,
// which is a series of nodes linked from head to tail.

// head: An atomic pointer to a Node that represents the
// front of the queue. The head is where items are removed
// from the queue (Pop operation). Using atomic.Pointer[Node]
// ensures that concurrent modifications to the head pointer
// (like removing the head node) are thread-safe.

// tail: Similar to the head, this is an atomic pointer to a
// Node but represents the end of the queue. This is where new
// items are added (Push operation). Again, using atomic.Pointer[Node]
// for the tail pointer ensures that appending to the queue can be
// done safely in a concurrent environment.

type Queue struct {
	head atomic.Pointer[Node] // Head pointer to Node
	tail atomic.Pointer[Node] // Tail pointer to Node
}

// func NewQueue(): Declares a function named NewQueue that takes no parameters.
// *Queue: The function returns a pointer to a Queue instance. This indicates
// that the function creates a Queue object and returns its memory address.

// n := &Node{}: This line initializes a new Node struct and assigns its
// memory address to the variable n. This Node is a "dummy" node, meaning
// it initially contains no data (value field is nil) and is not linked to
// any other nodes (next is nil by default because it’s zero-valued).

// q := &Queue{}: A new Queue struct is created, and its address is
// stored in the variable q. Initially, the head and tail atomic pointers
// within this struct are zero-valued (i.e., they are nil).

// q.head.Store(n) and q.tail.Store(n): These lines set both the head and
// tail pointers of the queue to point to the dummy node n. The use of the
// Store method ensures that these pointers are assigned atomically, meaning
// that the assignment operations are completed fully before any other operations
// can interact with these pointers. This atomicity is crucial for concurrent
// operations on the queue to ensure thread safety.

// head.Store(n) sets the head of the queue to the dummy node. In a queue,
// the head typically points to where items will be removed (dequeued).
// tail.Store(n) sets the tail of the queue to the same dummy node. In a queue,
// the tail points to where new items will be added (enqueued).

// return q: The function returns the pointer to the newly created Queue
// object. Now, the queue is initialized with both its head and tail pointers
// pointing to a dummy node, making it ready for operations like Push, Pop,
// Peek, etc.

// Summary
// The NewQueue function initializes a new queue with a single dummy
// node to simplify the enqueue and dequeue operations. By starting with a
// dummy node, you ensure that the head and tail are never nil, simplifying
// the logic in the other queue methods (particularly in concurrent contexts
// where checking for nil could complicate the implementation). This approach
// helps in managing edge cases, like popping or peeking in an empty queue,
// without special additional checks.

func NewQueue() *Queue {
	n := &Node{} // Initialize dummy node
	q := &Queue{}
	q.head.Store(n) // Set head to dummy node
	q.tail.Store(n) // Set tail to dummy node
	return q
}

// func (q *Queue) IsEmpty(): This declares a method named IsEmpty that
// is associated with the Queue type. The method does not take any parameters
// other than the implicit q, which is a pointer to the Queue instance (*Queue)
// on which the method is called.
// bool: The return type of the method is bool, indicating that this function
// returns a boolean value, true if the queue is empty and false otherwise.

func (q *Queue) IsEmpty() bool {
	return q.head.Load() == q.tail.Load()
}

func (q *Queue) Len() int64 {
	count := int64(0)
	for n := q.head.Load(); n != nil; n = n.next.Load() {
		if n.value != nil {
			count++
		}
	}
	return count
}

func (q *Queue) Push(x interface{}) {
	n := &Node{value: x}
	for {
		tail := q.tail.Load()
		if tail.next.CompareAndSwap(nil, n) {
			q.tail.CompareAndSwap(tail, n)
			return
		}
	}
}

func (q *Queue) Peek() interface{} {
	head := q.head.Load()
	next := head.next.Load()
	if next != nil {
		return next.value
	}
	return nil
}

func (q *Queue) Pop() interface{} {
	for {
		head := q.head.Load()
		tail := q.tail.Load()
		next := head.next.Load()
		if head == tail {
			if next == nil {
				return nil // Queue is empty
			}
			// Tail falling behind, help advance it
			q.tail.CompareAndSwap(tail, next)
		} else {
			if next != nil && q.head.CompareAndSwap(head, next) {
				return next.value
			}
		}
	}
}

func main() {
	//q := &Queue{}
	q := NewQueue()

	fmt.Println("this is queue library !!!")
	fmt.Println(q.IsEmpty())
	q.Push(1)
	q.Push(2)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Peek())
	q.Pop()
	fmt.Println(q.Peek())
	q.Pop()
	fmt.Println(q.Peek())
	fmt.Println(q.IsEmpty())
}
