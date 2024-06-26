package main

func main() {
	slicesExp()
	mapsExp()
  rangeExp()
  variadicExp()
  pointerExp()
  stringsExp()
  enumsExp()
  structEmbExp()
  genericsExp()
  errorsExp()
  goroutinesExp()
}

func clone[T comparable](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}
