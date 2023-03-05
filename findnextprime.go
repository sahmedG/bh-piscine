package piscine

func FindNextPrime(nb int) int {
	primecount := 0
	if nb == 0 {
		return 0
	} else if nb < 0 {
		return 2
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			primecount++
			break
		}
	}
	if primecount == 0 && nb != 1 {
		return nb
	}
	return FindNextPrime(nb + primecount)
}
