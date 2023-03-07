package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
    for _,char := range(a) {
        PrintStr(char)
    }
}

func PrintStr(s string) {
    str_rune := []rune(s)
	for _, word := range s {
		z01.PrintRune(word)
	}
    z01.PrintRune('\n')
}
