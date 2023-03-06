package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	numStr := ""
	if n == 0 {
		numStr = string(rune(n%10)+'0') + numStr
		n /= 10
	}
	for n > 0 {
		numStr = string(rune(n%10)+'0') + numStr
		n /= 10
	}
	for i := '0'; i <= '9'; i++ {
		for _, char := range numStr {
			if char == i {
				z01.PrintRune(char)
			}
		}
	}
}
