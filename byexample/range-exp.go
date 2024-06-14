package main

import "fmt"

func rangeExp() {
  for i, c := range "hello" {
    fmt.Printf("Index: %q, %v\n", i, i)
    fmt.Printf("Value: %q, %v\n", c, c)
  }
}
