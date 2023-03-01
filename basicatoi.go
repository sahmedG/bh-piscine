package piscine

func BasicAtoi(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		// Calculate the result for current iteration
		res = res*10 + int(s[i]-'0')
	}
	return res
}
