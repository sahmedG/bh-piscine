package piscine

func NRune(s string, n int) rune {
	firstrune := []rune(s)
	if n > len(firstrune) || n == 0 {
		return 0
	} else if n < 0 {
		return 0
	} else {
		last := firstrune[n-1]
		return last
	}
}
