package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""
	var str_arr []string
	if len(s) < 0 {
		return str_arr
	}
	for i := 0; i < len(s); i++ {
		if !(s[i] == ' ' || s[i] == '\n') {
			word = word + string(s[i])
		}
		if !(s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') {
			str_arr = append(str_arr, word)
			word = ""
		}
	}
	return str_arr
}
