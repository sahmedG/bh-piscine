package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argumet := os.Args[0]
	d := 0
	for i, letter := range argumet {
		if letter == '/' {
			d = i
		}
	}
	name := []rune(argumet)[d+1:]
	for _, i := range name {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
