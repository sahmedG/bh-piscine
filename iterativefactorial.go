package piscine

func IterativeFactorial(nb int) int {
	resutl := 1
	if nb <= 0 {
		return 0
	}
	for i := 1; i < nb+1; i++ {
		resutl = resutl * i
	}
	return resutl
}
