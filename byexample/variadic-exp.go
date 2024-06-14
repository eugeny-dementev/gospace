package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func sum[T constraints.Integer | constraints.Float](numbers ...T) T {
	var total T
	for _, v := range numbers {
		total += v
	}

	return total
}

func variadicExp() {
	fmt.Println("Total:", sum(1, 123, 3, 4, 3, 5))
}
