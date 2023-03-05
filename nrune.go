package piscine

func NRune(s string, n int) rune {
	firstrune := []rune(s)
	if n > len(firstrune) -1 {
		return 0
	} else if n < 0 {
		n = n * -1
		last := firstrune[len(firstrune)-1]
		return last
	} else {
		last := firstrune[n-1]
		return last
	}
}
