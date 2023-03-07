package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argumet := os.Args
	for i := len(argumet) - 1; i >= 1; i-- {
		name := []rune(argumet[i])
		for i := 0; i <= len(name)-1; i++ {
			z01.PrintRune(name[i])
		}
		z01.PrintRune('\n')
	}
}
