package piscine

func ToLower(str string) string {
	var ret_str = ""
	for _, chr := range str {
		if chr >= 64 && chr <= 90 {
			chr += 32
		}
		ret_str += string(chr)
	}
	return ret_str
}
