package piscine

func IsPrime(nb int) bool {
	if nb%2 == 0 || nb < 2{
		return false
	}
	return true

}
