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
	nonBlocking()
	blocking()
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
		time.Sleep(time.Millisecond * 20)
		c1 <- "result 1"
	}()

	select { // Promise.race basically
	case res := <-c1:
		fmt.Println("Non blocking", res)
	case <-time.After(time.Millisecond * 10):
		fmt.Println("timeout 1")
	}
}

func nonBlocking() {
	messages := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Message received", msg)
	default:
		fmt.Println("No messages received")
	}

	msg := "hi"
	select {
	case messages <- msg: // message cannot be sent if channel is non-buffered
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No messages sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("No activity")
	}
}

func blocking() {
	messages := make(chan string, 1) // 2. But if channel is buffered the panic will not happen and sent to channel would be immediately unblocked

	msg := "Hello"
	messages <- msg // 1. This will panic because channel is non-buffered and no receive setup before sent to the channel

	receivedMsg := <-messages
	fmt.Println("Received message:", receivedMsg)
}
