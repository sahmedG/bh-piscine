package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '8'; j++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			z01.PrintRune(' ')
			for k := '0'; k <= '9'; k++ {
				for y := '1'; y <= '9'; y++ {
					z01.PrintRune(k)
					z01.PrintRune(y)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
