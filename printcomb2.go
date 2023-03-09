package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			y := j + 1
			for k := i; k <= '9'; k++ {
				for ; y <= '9'; y++ {
					z01.PrintRune((i))
					z01.PrintRune((j))
					z01.PrintRune(' ')
					z01.PrintRune((k))
					z01.PrintRune((y))
					if i < '9' || j < '8' || k < '9' || y < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				y = '0'
			}
		}
	}
	z01.PrintRune('\n')

}