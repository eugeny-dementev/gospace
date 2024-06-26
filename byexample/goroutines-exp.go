package main

import (
	"fmt"
	"math/rand"
	"time"
)

func goroutinesExp() {
	fmt.Println("Goroutines experiments")

	gof("direct")
	go gof("goroutine")
	go func(msg string) {
		SleepRand(100)
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func gof(from string) {
	for i := 0; i < 3; i++ {
		SleepRand(50)
		fmt.Println(from, ":", i)
	}
}

func SleepRand(n int) {
	time.Sleep(time.Millisecond * time.Duration((rand.Intn(n))))
}
