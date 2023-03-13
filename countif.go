package piscine

func CountIf(f func(string) bool, a []string) int {
	l := len(a)
	count := 0
	result := make([]bool, l)
	for i := range a {
		result[i] = f(a[i])
	}
	for j := 0; j < len(result); j++ {
		if result[j] == true {
			count += 1
		}
	}
	return count
}
