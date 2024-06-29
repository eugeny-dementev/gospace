package main

import "fmt"

func channelsExp() {
	message := make(chan string)
	go func() {
		message <- "hello world"
	}()
  fmt.Println("Message:", <- message)
}
