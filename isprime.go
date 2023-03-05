package piscine

func IsPrime(nb int) bool {
	primecount := 0
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			primecount++
			break
		}
	}
	if primecount == 0 && nb != 1 {
		return true
	}
	return false
}
