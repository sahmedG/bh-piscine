package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr_1 := os.Args[1:]
	arr_2 := os.Args[1:]
	ln := 0
	for i := range arr_1 {
		ln = i
	}
	for i := 0; i <= ln; i++ {
		for j := 0; j <= ln; j++ {
			if arr_1[i] < arr_2[j] {
				arr_2[j], arr_1[i] = arr_1[i], arr_2[j]
			}
		}
	}
	for i := 0; i <= ln; i++ {
		for _, arg := range arr_2[i] {
			z01.PrintRune(arg)
		}
		z01.PrintRune(10)
	}
}
