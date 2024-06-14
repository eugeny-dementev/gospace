package main

import (
	"fmt"
	"slices"

	"golang.org/x/exp/constraints"
)

func main() {
	slicesExperiments()
}

func slicesExperiments() {
	var a []int
	fmt.Println("uninit:", a, a == nil, len(a))

	a = []int{1, 2, 3, 4}
	fmt.Println("uninit:", a, a == nil, len(a))

	a = make([]int, 10)
	a[5] = 5
	fmt.Println("uninit:", a, a == nil, len(a))

	a = a[5:]
	b := clone(a)
	fmt.Println(a, b)
	a[2] = 3
	b[2] = 3
	fmt.Println(a, b, slices.Equal(a, b))
}

func clone[T constraints.Integer](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}
