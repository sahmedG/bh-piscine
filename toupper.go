package piscine

func ToUpper(str string) string {
	ret_str := ""
	for _, chr := range str {
		if chr >= 97 && chr <= 122 {
			chr -= 32
		}
		ret_str += string(chr)
	}
	return ret_str
}
