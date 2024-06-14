package main

import (
	"fmt"
	"maps"
)

func mapsExp() {
	m := make(map[string]int)

	m["some"] = 80085
	m["yak"] = 6006
	fmt.Println(m, m["yak"], m["key"])

	fmt.Println("len", len(m))

	value, exist := m["yak"]
	fmt.Println(value, exist)
	value, exist = m["key"]
	fmt.Println(value, exist)

  s := map[string]int{"some":80085, "yak": 6006}
  fmt.Println("Equal:", maps.Equal(m, s))
}
