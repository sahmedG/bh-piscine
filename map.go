package piscine

func Map(f func(int) bool, a []int) []bool {
	l := len(a)
	result := make([]bool, l)
	for i := range a {
		result[i] = f(a[i])
	}
	return result
}
