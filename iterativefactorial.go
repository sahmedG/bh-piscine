package piscine

func IterativeFactorial(nb int) int {
	resutl := 1

	for i := 1; i < nb+1; i++ {
		resutl = resutl * i
	}
	return resutl
}