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

	anyChannelRead()
  timeouts()
}

func worker(id int, done chan<- bool) {
	fmt.Printf("worker %d is working\n", id)
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("worker %d is done\n", id)

	done <- true
}

func anyChannelRead() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(c chan<- string) {
		time.Sleep(time.Millisecond * 20)
		c <- "one"
	}(c1)
	go func(c chan<- string) {
		time.Sleep(time.Millisecond * 10)
		c <- "two"
	}(c2)

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("Message received from c1", msg1)
		case msg2 := <-c2:
			fmt.Println("Message received from c2", msg2)
		}
	}
}

func timeouts() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select { // Promise.race basically
	case res := <-c1:
		fmt.Println("Non blocking", res)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}
}
