package main

import (
	"./libs"
	"log"
	"time"
)

func main() {
	s1 := libs.NewStack()
	f1 := func() {
		log.Println("I am func1")
	}
	f2 := func() {
		log.Println("I am func2")
	}
	f3 := func() {
		log.Println("I am func3")
	}
	f4 := func() {
		log.Println("I am func4")
	}
	f5 := func() {
		log.Println("I am func5")
	}
	f6 := func() {
		log.Println("I am func6")
	}
	f7 := func() {
		log.Println("I am func7")
	}
	f8 := func() {
		log.Println("I am func8")
	}
	t1 := libs.NewTask(f1, f2)
	t2 := libs.NewTask(f3, f4)
	t3 := libs.NewTask(f5)
	t4 := libs.NewTask(f6)
	t5 := libs.NewTask(f7)
	t6 := libs.NewTask(f8)
	s1.Push(t1)
	s1.Push(t2)
	s1.Push(t3)
	s1.Push(t4)
	s1.Push(t5)
	s1.Push(t6)
	dispatcher := libs.NewDispatcher(3, s1)
	dispatcher.Run()
	<-time.After(3 * time.Second)
}
