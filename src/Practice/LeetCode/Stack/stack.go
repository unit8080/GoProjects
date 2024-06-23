package main

import (
	"errors"
	"fmt"
	"sync"
)

type Stack struct {
	stack []interface{}
	mutex sync.Mutex
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
func (s *Stack) Len() int {
	return len(s.stack)
}
func (s *Stack) Push(x interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	(*s).stack = append((*s).stack, x)
	return
}
func (s *Stack) Peek() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.IsEmpty() {
		length := s.Len()
		x := s.stack[length-1]
		return x
	}
	return nil
}
func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.IsEmpty() {
		length := s.Len()
		x := s.stack[length-1]
		(*s).stack = (*s).stack[:length-1]
		return x
	}
	errorMsg := "No entry exists"
	err := errors.New(errorMsg)
	fmt.Println(err.Error())
	return nil
}

func main() {
	s := &Stack{}
	fmt.Println("test")
	fmt.Println(s.IsEmpty())
	s.Push(2)
	s.Push(3)
	fmt.Println("top:", s.Peek())
	s.Pop()
	fmt.Println("top:", s.Peek())
	s.Pop()
	fmt.Println(s.IsEmpty())
	s.Pop()
}
