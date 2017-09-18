package main

import (
	"log"
	"sync"
)

type Stack struct {
	stack map[int]interface{}
	sync.Mutex
}

func NewStack() *Stack {
	sk := make(map[int]interface{})
	return &Stack{
		stack: sk,
	}
}

func (s *Stack) Push(v interface{}) {
	defer s.Unlock()
	s.Lock()
	s.stack[len(s.stack)] = v
}

func (s *Stack) Pop() interface{} {
	defer func() {
		delete(s.stack, len(s.stack)-1)
		s.Unlock()
	}()
	s.Lock()
	return s.stack[len(s.stack)-1]
}

func main() {
	s1 := NewStack()
	s1.Push("1")
	s1.Push("2")
	s1.Push("3")
	s1.Push("4")
	s1.Push("5")
	log.Println(s1.Pop())
	log.Println(s1.Pop())
	log.Println(s1.Pop())
	log.Println(s1.Pop())
	log.Println(s1.Pop())
}
