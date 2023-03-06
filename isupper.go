package piscine

func IsUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	count := 0
	byte1 := []rune(s)
	for i := 0; i < len(s); i++ {
		if byte1[i] >= 'A' && byte1[i] <= 'Z' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
