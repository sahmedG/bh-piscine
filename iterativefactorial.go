package piscine

func IterativeFactorial(nb int) int {
	resutl := 1
	if nb <= 0 {
		return 0
	} else {
		for nb >= 1 {
			resutl *= nb
			nb -= 1
		}
	}
	if resutl < 0 {
		return 0
	}
	return resutl
}
