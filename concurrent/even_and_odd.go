package main

import (
	"fmt"
	"sync"
)

//交替打印奇偶数

var NUMBER = 100

func oddGoRoutine(p chan int, group *sync.WaitGroup) {
	defer group.Done()
	for i := 1; i <= NUMBER; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("奇数打印:", i)
		}
	}
}

func evenGoRoutine(p chan int, group *sync.WaitGroup) {
	defer group.Done()
	for i := 1; i <= NUMBER; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("偶数打印:", i)
		}
	}
}

func main() {
	msg := make(chan int)
	var s sync.WaitGroup
	s.Add(2)
	go oddGoRoutine(msg, &s)
	go evenGoRoutine(msg, &s)
	s.Wait()
}
