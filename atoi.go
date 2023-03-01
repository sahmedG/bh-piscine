package piscine

func Atoi(s string) int {
	var start int
	var res int

	var signMultiplier int = 1
	if s[0] == '-' {
		signMultiplier = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	byte_str := []rune(s)
	for j := 0; j < len(byte_str); j++ {
		if byte_str[j] == ' ' {
			res = 0
		}
	}

	for i := start; i < len(byte_str); i++ {
		if !(byte_str[i] >= '0' && byte_str[i] <= '9') {
			return 0
		}
		res = res*10 + int(s[i]-'0')
	}

	return res * signMultiplier
}
