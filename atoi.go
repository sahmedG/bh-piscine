package piscine

func Atoi(s string) int {
	var start int = 0
	var res int
	var signMultiplier int = 1

	if len(s) == 0 {
		return 0
	}
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

	for ; start < len(byte_str); start++ {
		if !(byte_str[start] >= '0' && byte_str[start] <= '9') {
			return 0
		}
		res = res*10 + int(s[start]-'0')
	}

	return res * signMultiplier
}
