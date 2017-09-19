package main

import (
	"./libs"
	"log"
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
	t1 := libs.NewTask(f1, f2)
	t2 := libs.NewTask(f3, f4)
	s1.Push(t1)
	s1.Push(t2)
	s1.Pop().Run()
	s1.Pop().Run()
}
