package piscine

func IterativeFactorial(nb int) int {
	Result := 1
	if nb < 0 || nb > 12 {
		Result = 0
	}
	if nb <= 12 {
		for nb >= 1 {
			Result *= nb
			nb -= 1
		}
		if Result < 0 {
			Result = 0
		}
		return Result
	}
	return Result
}
