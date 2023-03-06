package piscine

func TrimAtoi(s string) int {
	// var start int = 0
	var res int
	SignMultiplier := 1
	if len(s) == 0 {
		return 0
	}
	byte_str := []rune(s)
	for i := 0; i < len(byte_str); i++ {
		if byte_str[i] == '-' && res == 0 {
			SignMultiplier = -1
		}
		if !(byte_str[i] >= '0' && byte_str[i] <= '9') {
			continue
		}

		res = res*10 + int(s[i]-'0')
	}
	return res * SignMultiplier
}
