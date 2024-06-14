package main

import (
	"fmt"
	"unicode/utf8"
)

func stringsExp() {
	const s = "สวัสดี"

	r := []rune(s)
	fmt.Println(s, len(s), len(r), utf8.RuneCountInString(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for i, rune := range s {
		fmt.Printf("Rune: %#U %x %v\n", rune, rune, i)
	}
}
