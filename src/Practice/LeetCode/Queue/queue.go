package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Queue struct {
	queue   []interface{}
	size    atomic.Int64
	rwMutex sync.RWMutex
}

func (q *Queue) IsEmpty() bool {

	size := q.size.Load()
	return size == 0
}
func (q *Queue) Len() int64 {

	size := q.size.Load()
	return size
}
func (q *Queue) Push(x interface{}) {

	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()

	(*q).queue = append((*q).queue, x)
	q.size.Add(1)
	return
}
func (q *Queue) Peek() interface{} {

	q.rwMutex.RLock()
	defer q.rwMutex.RUnlock()

	if !q.IsEmpty() {
		x := q.queue[0]
		return x
	}
	return nil
}
func (q *Queue) Pop() interface{} {

	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()

	if q.IsEmpty() {
		return nil
	}
	popped := q.queue[0]
	(*q).queue = (*q).queue[1:]

	q.size.Add(-1) // reason why size shoudld be signed int
	return popped
}
func main() {
	q := &Queue{}

	fmt.Println("this is queue library !!!")
	fmt.Println(q.IsEmpty())
	q.Push(1)
	q.Push(2)
	q.Push("Hello")
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Peek())
	q.Pop()
	fmt.Println(q.Peek())
	q.Pop()
	fmt.Println(q.Peek())
	fmt.Println(q.IsEmpty())
}
