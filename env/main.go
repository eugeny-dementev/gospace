package main

import (
	"fmt"
	"os"
)

func main() {
  fmt.Println("GOBIN:", os.Getenv("GOBIN"))
}
