package piscine

func BasicAtoi2(s string) int {
	var start int
	var res int

	// Handling numbers with +/- sign
	signMultiplier := 1
	if s[0] == '-' {
		// Handling for negative numbers
		signMultiplier = -1
		start = 1
	} else if s[0] == '+' {
		// Handling for signed positive number
		start = 1
	}

	//return res * signMultiplier

	byte_str := []rune(s)
	for j := start; j < len(byte_str); j++ {
		if byte_str[j] == ' ' {
			break
		}
		if !(byte_str[j] >= '0' && byte_str[j] <= '9') {
			return res * signMultiplier
		} else {
			res = res*10 + int(byte_str[j]-'0')
		}
	}
	return res * signMultiplier
}
