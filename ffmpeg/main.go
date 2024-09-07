package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "./ffmpeg/main.go")

	fmt.Println("cmd", cmd)

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

  fmt.Println("Output:", string(out))
}
