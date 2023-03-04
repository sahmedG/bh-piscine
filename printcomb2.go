package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for y := '0'; y <= '9'; y++ {
					if i == '0' && j == '0' && k == '0' && y == '0' {
						continue
					} else if i == k && y <= j {
						continue
					} else {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(y)
						if i == '9' && j == '8' && k == '9' && y == '9' {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
