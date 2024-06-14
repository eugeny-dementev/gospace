package main

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {
	input := "\xff"
	rev, err := Reverse(input)
	if err != nil {
		log.Fatal(err)
	}
	doubleRev, err := Reverse(rev)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
	fmt.Println(rev)
	fmt.Println(doubleRev)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8 string")
	}

	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
