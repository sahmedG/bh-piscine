package piscine

func IterativeFactorial(nb int) int {
	Result := 0
	if nb <= 12 {
		for nb >= 1 {
			Result *= nb
			nb -= 1
		}
		if Result < 0 {
			return 0
		}
		return Result
	}
	return 0
}
