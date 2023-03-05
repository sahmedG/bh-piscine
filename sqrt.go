package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 0
	result = nb / 2
	if result%2 == 0 {
		return result
	}
	return 0
}
