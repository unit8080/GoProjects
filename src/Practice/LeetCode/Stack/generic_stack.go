package main

import (
	"fmt"
	"reflect"
)

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value
}
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	value := intStack.Pop()
	fmt.Println("Value:", value)

	stringStack := NewStack[string]()
	stringStack.Push("hello")
	stringStack.Push("there!")
	str := stringStack.Pop()
	fmt.Println("String:", str)
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) Peek() T {
	index := len(*s) - 1
	return (*s)[index]
}

func (s *Stack[T]) Clear() {
	(*s) = []T{}
}

func (s *Stack[T]) Contains(item T) bool {
	for _, value := range *s {
		if reflect.DeepEqual(item, value) {
			return true
		}
	}
	return false
}

func (s *Stack[T]) Values() []T {
	return *s
}
