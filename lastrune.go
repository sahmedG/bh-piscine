package piscine

func LastRune(s string) rune {
	firstrune := []rune(s)
	last := firstrune[len(s)-1]
	return last
}
