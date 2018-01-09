package TaskDispatcher

import (
	"fmt"
	"os"
	"time"
)

// Task sturct
type Task struct {
	maxGrNum  int
	tasks     []tasker
	pool      chan int
	taskQueue chan tasker
}

// NewTask constructor
func NewTask(gr int) *Task {
	if gr <= 0 {
		fmt.Println("Max goroutine num is invalid!")
		os.Exit(-1)
	}
	return &Task{
		maxGrNum:  gr,
		tasks:     make([]tasker, 0),
		pool:      make(chan int, gr),
		taskQueue: make(chan tasker, 1),
	}
}

// AddTasks adds tasks
func (t *Task) AddTasks(tasks ...tasker) *Task {
	for _, task := range tasks {
		t.tasks = append(t.tasks, task)
	}
	return t
}

// gin prepare goroutine index to pool
func (t *Task) gin() {
	for index := 1; index <= t.maxGrNum; index++ {
		t.pool <- index
	}
}

// accquire get a source index
func (t *Task) accquire() int {
	select {
	case i := <-t.pool:
		return i
	default:
		return 0
	}
}

// release back a source index
func (t *Task) release(i int) {
	t.pool <- i
}

func (t *Task) inQueue() {
	for _, task := range t.tasks {
		t.taskQueue <- task
	}
}

// Run start all task
func (t *Task) Run() {
	t.gin()
	go t.inQueue()
	t.taskLoop()
}

func (t *Task) taskLoop() {
	for {
		select {
		case task := <-t.taskQueue:
			go t.goTask(task)
		default:
			<-time.After(100 * time.Millisecond)
		}
	}
}

func (t *Task) goTask(task tasker) {
	index := t.accquire()
	for index == 0 {
		fmt.Println("Now have no enough goroutine resources, please wait...")
		<-time.After(100 * time.Millisecond)
		index = t.accquire()
	}
	go func(tasker, int) {
		defer func() {
			t.release(index)
		}()
		fmt.Printf("Goroutine %d do one task ... ", index)
		task.Run()
	}(task, index)
}
