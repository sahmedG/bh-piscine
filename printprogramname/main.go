package main

import (
	"fmt"
	"os"
)

func printprogramname() {
	prg_name := os.Args[0]
	fmt.Printf(prg_name)
}

func main() {
	printprogramname()
}
