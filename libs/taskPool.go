package libs

import (
	"sync"
)

type Stack struct {
	stack map[int]*Task
	sync.Mutex
}

func NewStack() *Stack {
	sk := make(map[int]*Task)
	return &Stack{
		stack: sk,
	}
}

func (s *Stack) Push(v *Task) {
	defer s.Unlock()
	s.Lock()
	s.stack[len(s.stack)] = v
}

func (s *Stack) Pop() *Task {
	defer func() {
		delete(s.stack, len(s.stack)-1)
		s.Unlock()
	}()
	s.Lock()
	return s.stack[len(s.stack)-1]
}
