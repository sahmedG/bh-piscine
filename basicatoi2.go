package piscine

func BasicAtoi2(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return res
		}
	}
	byte_str := []rune(s)
	for j := 0; j < len(byte_str); j++ {
		if byte_str[j] == ' ' {
			break
		} else {
			res = res*10 + int(byte_str[j]-'0')
		}
	}
	return res
}
