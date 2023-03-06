package piscine

func IsPrintable(s string) bool {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if str[i] == 9 || str[i] == 10 || str[i] == 11 || str[i] == 12 || str[i] == 13 {
			return false
		}
	}
	return true
}
