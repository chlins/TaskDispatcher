package libs

import (
	"log"
)

//Dispatcher is task dispatcher
type Dispatcher struct {
	stack     *Stack
	workerNum int
	workers   []*worker
}

type worker struct {
	workerIndex int
	status      chan string
	info        map[int]int //map[workerIndex]taskIndex
}

//NewDispatcher is construct func of dispatcher
func NewDispatcher(num int, s *Stack) *Dispatcher {
	ws := make([]*worker, num)
	for i := 0; i < num; i++ {
		st := make(chan string, 1)
		st <- "free"
		inf := make(map[int]int)
		ws[i] = &worker{workerIndex: i, status: st, info: inf}
	}
	return &Dispatcher{
		stack:     s,
		workerNum: num,
		workers:   ws,
	}
}

//Run is func of dispatcher starting
func (d *Dispatcher) Run() {
	for {
		for index, w1 := range d.workers {
			select {
			case ss := <-w1.status:
				if ss == "free" {
					w1.info[index] = len(d.stack.stack) - 1
					w1.status <- "busy"
					d.stack.Pop().Run()
					log.Printf("Worker %d has done task %d \n", index, w1.info[index])
					<-w1.status
					w1.status <- "free"
				}
				if len(d.stack.stack) == 0 {
					return
				}

			}

		}
	}
}
