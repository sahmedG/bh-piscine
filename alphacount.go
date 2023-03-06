package piscine

func AlphaCount(s string) int {
	count := 0
	byte1 := []rune(s)
	for i := 0; i < len(s); i++ {
		if byte1[i] >= 'a' && byte1[i] <= 'z' || byte1[i] >= 'A' && byte1[i] <= 'Z' {
			count++
		}
	}
	return count
}
