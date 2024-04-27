package main

import "fmt"

type Items struct {
	items []interface{}
}

func (q *Items) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Items) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	q := new(Items)

	q.items = make([]interface{}, 0)

	q.Enqueue(10)
	q.Enqueue("Hello")
	q.Enqueue(20)

	fmt.Println("Dequeued item:", q.Dequeue())
	fmt.Println("Dequeued item:", q.Dequeue())
	fmt.Println("Dequeued item:", q.Dequeue())
	fmt.Println("Dequeued item:", q.Dequeue())
	fmt.Println("Dequeued item:", q.Dequeue())
}
