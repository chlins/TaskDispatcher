package libs

import (
	"sync"
)

//Stack is a task pool
type Stack struct {
	stack map[int]*Task
	sync.Mutex
}

//NewStack is a construct of task pool
func NewStack() *Stack {
	sk := make(map[int]*Task)
	return &Stack{
		stack: sk,
	}
}

//Push a task
func (s *Stack) Push(v *Task) {
	defer s.Unlock()
	s.Lock()
	s.stack[len(s.stack)] = v
}

//Pop a task
func (s *Stack) Pop() *Task {
	defer func() {
		delete(s.stack, len(s.stack)-1)
		s.Unlock()
	}()
	s.Lock()
	return s.stack[len(s.stack)-1]
}
