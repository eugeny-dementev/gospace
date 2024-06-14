package main

import "fmt"

func byvalue(value int) {
	value = 5
}

func bylink(pointer *int) {
	*pointer = 5
}

func pointerExp() {
	i := 1

	byvalue(i)
	fmt.Println("By value:", i)

  bylink(&i)
	fmt.Println("By value:", i)

  fmt.Println("Pointer:", &i)
}
