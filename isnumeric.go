package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	count := 0
	byte1 := []byte(s)
	for i := 0; i < len(byte1); i++ {
		if byte1[i] >= '0' && byte1[i] <= '9' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
