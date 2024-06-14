package main

func main() {
	slicesExperiments()
}

func clone[T comparable](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}
