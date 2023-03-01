package piscine

func BasicAtoi2(s string) int {
	var res int

	for i := 0; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return res
		}
		// Calculate the result for current iteration
		res = res*10 + int(s[i]-'0')
	}
	return res
}
