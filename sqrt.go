package piscine

func Sqrt(nb int) int {
	i := 0
	for i <= nb {
		if i*i == nb {
			return i
		}
		i++
	}
	return 0
}
