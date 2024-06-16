package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0)
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0)
	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func genericsExp() {
	m := map[string]int{
		"hello": 1,
		"world": 2,
    "blob": 3,
	}

	fmt.Println("keys", MapKeys(m))
	fmt.Println("values", MapValues(m))
}
