package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func goroutinesExp() {
	fmt.Println("Goroutines experiments")

	var wg sync.WaitGroup

	gof("direct")
	func() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			gof("goroutine")
		}()
	}()
	func() {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			SleepRand(1000)
			fmt.Println(msg)
		}("going")
	}()

	wg.Wait()
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