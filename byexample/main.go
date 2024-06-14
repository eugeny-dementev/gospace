package main

func main() {
	slicesExp()
	mapsExp()
  rangeExp()
  variadicExp()
}

func clone[T comparable](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}
