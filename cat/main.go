package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('H')
		z01.PrintRune('e')
		z01.PrintRune('l')
		z01.PrintRune('l')
		z01.PrintRune('o')
		z01.PrintRune('\n')
		z01.PrintRune('H')
		z01.PrintRune('e')
		z01.PrintRune('l')
		z01.PrintRune('l')
		z01.PrintRune('o')
		z01.PrintRune('\n')
	}
	for i := 0; i < len(args); i++ {
		data, err := ioutil.ReadFile(args[i])
		if err != nil {
			for _, err := range err.Error() {
				z01.PrintRune(err)
			}
			z01.PrintRune('\n')
		}
		for _, char := range data {
			z01.PrintRune(rune(char))
		}
	}
}
