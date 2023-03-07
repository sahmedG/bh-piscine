package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printprogramname() {
	prg_name := os.Args[0]
	z01.PrintRune(prg_name)
}

func main() {
	printprogramname()
}
