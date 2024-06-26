package main

import (
	"fmt"
	"time"
)

func goroutinesExp() {
	fmt.Println("Goroutines experiments")

  gof("direct")
  go gof("goroutine")
  go func(msg string) {
    time.Sleep(time.Millisecond)
    fmt.Println(msg)
  }("going")

  time.Sleep(time.Second)
  fmt.Println("done")
}

func gof(from string) {
  for i := 0; i < 3; i++ {
    time.Sleep(time.Millisecond)
    fmt.Println(from, ":", i)
  }
}
