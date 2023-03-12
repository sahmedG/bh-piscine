package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		input, err := ioutil.ReadAll(os.Stdin)
		for _, char := range string(input) {
			z01.PrintRune((char))
		}

		if err != nil {
			for _, err := range err.Error() {
				z01.PrintRune(err)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
	}

	for i := 0; i < len(args); i++ {
		data, err := ioutil.ReadFile(args[i])
		if err != nil {
			z01.PrintRune('E')
			z01.PrintRune('R')
			z01.PrintRune('R')
			z01.PrintRune('O')
			z01.PrintRune('R')
			z01.PrintRune(':')
			z01.PrintRune(' ')
			for _, err := range err.Error() {
				z01.PrintRune(err)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		for _, char := range data {
			z01.PrintRune(rune(char))
		}
	}
}
