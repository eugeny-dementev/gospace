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
  fmt.Println("Message:", <- message)

  done := make(chan bool)
  go worker(5, done)

  <- done
  fmt.Println("All workers finished")
}

func worker(id int, done chan bool) {
  fmt.Printf("worker %d is working\n", id)
  time.Sleep(time.Millisecond * 10)
  fmt.Printf("worker %d is done\n", id)

  done <- true
}
