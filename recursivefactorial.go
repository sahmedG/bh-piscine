package piscine

func RecursiveFactorial(nb int) int {
	Result := 1
	if nb < 0 || nb > 26 {
		Result = 0
	}
	if nb <= 26 {
		if nb == 1 || nb == nb {
			Result = nb
		}
		Result = nb * RecursiveFactorial(nb-1)
	}
	if Result < 0 {
		Result = 0
	} 
	return Result
}
