package main

import (
	"fmt"
	"time"
)

func channelsExp() {
	message := make(chan string)
	go func() {
		message <- "hello world"
	}()
	fmt.Println("Message:", <-message)

	done := make(chan bool)

	list := []int{1, 2, 3, 4, 5}
	doneList := make([]int, 5)
	for _, v := range list {
		go worker(v, done)
	}

	for v := range len(list) {
		<-done
		doneList[v] = 1
	}

	for _, v := range doneList {
		if v != 1 {
			panic("not all workers finished")
		}
	}
	fmt.Println("All workers finished", len(doneList))
}

func worker(id int, done chan bool) {
	fmt.Printf("worker %d is working\n", id)
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("worker %d is done\n", id)

	done <- true
}
