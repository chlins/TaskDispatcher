package main

import (
	"fmt"

	"github.com/chlins/TaskDispatcher"
)

type t1 struct{}

func (t *t1) Run() {
	fmt.Println("I am t1 func...")
}

type t2 struct{}

func (t *t2) Run() {
	fmt.Println("I am t2 func...")
}

type t3 struct{}

func (t *t3) Run() {
	fmt.Println("I am t3 func...")
}

type t4 struct{}

func (t *t4) Run() {
	fmt.Println("I am t4 func...")
}

type t5 struct{}

func (t *t5) Run() {
	fmt.Println("I am t5 func...")
}

func main() {
	f1, f2, f3, f4, f5 := new(t1), new(t2), new(t3), new(t4), new(t5)
	t := TaskDispatcher.NewTask(2)
	t.AddTasks(f1, f2, f3, f4, f5).Run()
}
